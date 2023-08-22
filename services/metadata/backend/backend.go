package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/snabb/sitemap"
)

type Project struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	Metadata    interface{} `json:"metadata"`
}

type Projects struct {
	dataJSON   []Project
	dataJSONLD []Project
	dataTTL    []Project
	dataXML    []Project
}

// All projects that are being served
var projects Projects

type spaHandler struct {
	staticPath string
	indexPath  string
}

// Full text search of a project.
// Returns a slice of Projects where each project matches the search query.
// Note: The query is a regex pattern and is matched against the JSON representation of the project.
func searchProjects(query string) []Project {
	var res []Project
	for _, project := range projects.dataJSON {
		content, _ := json.Marshal(project.Metadata)
		match, _ := regexp.Match("(?i)"+query, content)
		if match {
			res = append(res, project)
		}
	}
	return res
}

func filterProjectsByStatus(projects []Project, filter string) []Project {
	if filter == "" {
		return projects
	}
	showOngoing := !strings.Contains(filter, "o")
	showFinished := !strings.Contains(filter, "f")
	var res []Project
	for _, project := range projects {
		if project.Status == "ongoing" && showOngoing {
			res = append(res, project)
		} else if project.Status == "finished" && showFinished {
			res = append(res, project)
		}
	}

	return res
}

func getStatus(shortcode string) string {
	projectOngoingTable := map[string]bool{
		"0118": true,
		"0806": true,
		"0813": true,
		"081B": true,
		"0805": true,
		"082C": true,
		"0107": true,
		"0836": true,
		"080E": true,
		"0105": true,
		"0828": true,
		"0112": true,
		"0820": true,
		"0114": true,
		"0116": true,
		"0801": true,
		"082B": true,
		"0103": true,
		"083A": true,
		"083B": true,
		"0812": true,
		"0827": true,
		"0807": true,
		"0816": true,
		"0804": true,
		"080C": true,
		"083C": true,
		"0119": true,
		"0102": true,
		"0843": true,
		"0121": true,
		"0844": true,
		"0810": true,
		"0106": true,
	}
	_, isOngoing := projectOngoingTable[shortcode]
	if isOngoing {
		return "ongoing"
	}
	return "finished"
}

// Loads a project from a JSON files
// Expects this file to be located in ./data/*.json
func loadProject(path string) Project {
	log.Printf("Loading: %v", path)

	// read json
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal json
	jsonMap := make(map[string]interface{})
	err2 := json.Unmarshal(byteValue, &jsonMap)
	if err2 != nil {
		log.Fatal(err)
	}

	// grab actual metadata from JSON
	// projMetadata, ok := jsonMap["projectsMetadata"].([]interface{})
	projectMap, ok := jsonMap["project"].(map[string]interface{})

	if ok {
		// projectMap := findProjectNode(projMetadata)
		id := projectMap["shortcode"].(string)
		name := projectMap["name"].(string)
		description := projectMap["teaserText"].(string)
		status := getStatus(id)
		return Project{
			ID:          id,
			Name:        name,
			Description: description,
			Status:      status,
			Metadata:    jsonMap,
		}
	} else {
		log.Fatal("Could not find project in JSON")
		return Project{}
	}
}

// Loads Projects
func loadProjectData() Projects {
	var res Projects

	// get all JOIN and JSON-LD files
	pathPrefix := "./services/metadata/backend/data/"
	paths, _ := filepath.Glob(pathPrefix + "*.json")

	// loop thought paths and prepare response
	for _, path := range paths {
		file := filepath.Base(path)

		// omit files which name starts with _
		if !strings.HasPrefix(file, "_") {
			// load JSON
			project := loadProject(path)
			res.dataJSON = append(res.dataJSON, project)

			// load potential RDFs
			filename := strings.TrimSuffix(file, filepath.Ext(file))
			pathLD := pathPrefix + filename + ".jsonld"
			pathTTL := pathPrefix + filename + ".ttl"
			pathXML := pathPrefix + filename + ".xml"

			// load JSON-LD
			byteValueLD, err := ioutil.ReadFile(pathLD)
			if err == nil {
				jsonMap := make([]map[string]interface{}, 100)
				err2 := json.Unmarshal(byteValueLD, &jsonMap)
				if err2 == nil {
					res.dataJSONLD = append(res.dataJSONLD, Project{
						ID:          project.ID,
						Name:        project.Name,
						Description: project.Description,
						Metadata:    jsonMap,
					})
				} else {
					log.Fatal(err2)
				}
			}

			// load Turtle
			byteValueTTL, err := ioutil.ReadFile(pathTTL)
			if err == nil {
				ttl := string(byteValueTTL)
				res.dataTTL = append(res.dataTTL, Project{
					ID:          project.ID,
					Name:        project.Name,
					Description: project.Description,
					Metadata:    ttl,
				})
			}

			// load RDF/XML
			byteValueXML, err := ioutil.ReadFile(pathXML)
			if err == nil {
				xml := string(byteValueXML)
				res.dataXML = append(res.dataXML, Project{
					ID:          project.ID,
					Name:        project.Name,
					Description: project.Description,
					Metadata:    xml,
				})
			}
		}
	}

	return res
}

// Gets projects on route: /projects
// request parameters are only provided for JSON requests
func getProjects(w http.ResponseWriter, r *http.Request) {
	var matches []Project

	// depending on the Accept header, return JSON or RDF data
	accept := r.Header.Get("Accept")
	if accept == "application/ld+json" {
		w.Header().Set("Content-Type", "application/json")
		matches = make([]Project, len(projects.dataJSONLD))
		copy(matches, projects.dataJSONLD)
		log.Printf("JSON-LD request for: %v", r.URL)
	} else if accept == "text/turtle" {
		w.Header().Set("Content-Type", "application/json")
		matches = make([]Project, len(projects.dataTTL))
		copy(matches, projects.dataTTL)
		log.Printf("TURTLE request for: %v", r.URL)
	} else if accept == "application/rdf+xml" {
		w.Header().Set("Content-Type", "application/json")
		matches = make([]Project, len(projects.dataTTL))
		copy(matches, projects.dataXML)
		log.Printf("RDF-XML request for: %v", r.URL)
	} else {
		w.Header().Set("Content-Type", "application/json")

		// request parameters
		query := r.URL.Query().Get("q")
		// TODO: does page start at 0 or 1?
		page, _ := strconv.Atoi(r.URL.Query().Get("_page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("_limit"))
		filter := r.URL.Query().Get("filter")

		matches = make([]Project, len(projects.dataJSON))

		if query == "" {
			// no search query all projects are matches
			copy(matches, projects.dataJSON)
		} else {
			// reduce projects by search
			matches = searchProjects(query)
		}

		matches = filterProjectsByStatus(matches, filter)

		w.Header().Set("X-Total-Count", strconv.Itoa(len(matches)))

		// paginate
		if len(matches) > 1 && len(matches) > limit && page > 0 && limit > 0 {
			max := len(matches)
			start := (page - 1) * limit
			if start > max {
				start = max
			}
			end := page * limit
			if end > max {
				end = max
			}

			matches = matches[start:end]
		}

		log.Printf("JSON request for: %v", r.URL)
	}

	json.NewEncoder(w).Encode(matches)
}

// Gets a single project on route /projects/:id
func getProject(w http.ResponseWriter, r *http.Request) {
	var data []Project
	params := mux.Vars(r)

	// depending on the Accept header, return JSON or RDF data
	accept := r.Header.Get("Accept")
	if accept == "application/ld+json" {
		data = projects.dataJSONLD
		log.Printf("JSON-LD request for: %v", r.URL)
		for _, item := range data {
			for item.ID == params["id"] {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(item.Metadata)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		message := fmt.Sprintf("No JSON-LD serialization available for project %v", params["id"])
		w.Write([]byte(message))
		return
	} else if accept == "text/turtle" {
		data = projects.dataTTL
		log.Printf("TURTLE request for: %v", r.URL)
		for _, item := range data {
			for item.ID == params["id"] {
				w.Header().Set("Content-Type", "text/turtle")
				w.Write([]byte(item.Metadata.(string)))
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		message := fmt.Sprintf("No Turtle serialization available for project %v", params["id"])
		w.Write([]byte(message))
		return
	} else if accept == "application/rdf+xml" {
		data = projects.dataXML
		log.Printf("RDF-XML request for: %v", r.URL)
		for _, item := range data {
			for item.ID == params["id"] {
				w.Header().Set("Content-Type", "application/rdf+xml")
				w.Write([]byte(item.Metadata.(string)))
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		message := fmt.Sprintf("No RDF-XML serialization available for project %v", params["id"])
		w.Write([]byte(message))
		return
	} else {
		data = projects.dataJSON
		log.Printf("JSON request for: %v", r.URL)
		for _, item := range data {
			for item.ID == params["id"] {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(item.Metadata)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		message := fmt.Sprintf("No project %v available", params["id"])
		w.Write([]byte(message))
		return
	}
}

// getSitemap returns the sitemap.xml containing routes to all project pages.
// Route /sitemap.xml
func getSitemap(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request for: %v", r.URL)

	sm := sitemap.New()
	sm.Add(&sitemap.URL{
		Loc:        "https://meta.dasch.swiss/",
		ChangeFreq: sitemap.Weekly,
	})

	for _, item := range projects.dataJSON {
		projectUrl := fmt.Sprintf("https://meta.dasch.swiss/projects/%s/", item.ID)
		sm.Add(&sitemap.URL{
			Loc:        projectUrl,
			ChangeFreq: sitemap.Weekly,
		})
	}
	sm.WriteTo(w)
}

// getRobotsFile returns the robots.txt file containing the reference to the sitemap
// Route /robots.txt
func getRobotsFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request for: %v", r.URL)
	rf := "Sitemap: https://meta.dasch.swiss/sitemap.xml\nUser-agent: *\nDisallow:"
	io.WriteString(w, rf)
}

// servers the contnet version.txt file
// Route /version.txt
func getVersionFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request for: %v", r.URL)
	vf, err := ioutil.ReadFile("./version.txt")
	if err != nil {
		fmt.Println("Error creating", "version.txt")
		fmt.Println(err)
		return
	}
	io.WriteString(w, string(vf))
}

// handle SPA to serve always from right place, no matter of route
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("SPA Handler: %v", r.URL)

	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err2 := os.Stat(path)
	if err2 == nil {
		// file exists -> serve file
		// log.Printf("Serving from File Server: %v", path)
		http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
		return
	} else {
		// file does not exist, see where to go from here
		pattern := "/projects/?([0-9A-F]{4})?"
		match, _ := regexp.MatchString(pattern, path)
		if match {
			// file matches "/project/shortcode" pattern -> remove this section of the path
			re := regexp.MustCompile(pattern)
			s := re.ReplaceAllString(path, "/")
			_, err3 := os.Stat(s)
			if err3 == nil {
				// file exists after removing the section -> serve this file
				// log.Printf("Existis after changing: %v", s)
				http.ServeFile(w, r, s)
				return
			}
		}

		// file still not found, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
	}
}

func main() {
	// CORS header
	ch := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	// load Data
	projects = loadProjectData()
	log.Printf("Loaded %v files (%v JSON + %v JSON-LD + %v TTL + %v RDF-XML)",
		len(projects.dataJSON)+len(projects.dataJSONLD)+len(projects.dataTTL),
		len(projects.dataJSON),
		len(projects.dataJSONLD),
		len(projects.dataTTL),
		len(projects.dataXML))

	// init Router
	router := mux.NewRouter()

	// set up routes
	router.HandleFunc("/api/v1/projects", getProjects).Methods("GET")
	router.HandleFunc("/api/v1/projects/{id}", getProject).Methods("GET")
	router.HandleFunc("/robots.txt", getRobotsFile).Methods("GET")
	router.HandleFunc("/sitemap.xml", getSitemap).Methods("GET")
	router.HandleFunc("/version.txt", getVersionFile).Methods("GET")

	// init SPA handler
	spa := spaHandler{
		staticPath: "public",
		indexPath:  "index.html",
	}

	// apply SPA handler
	router.PathPrefix("/").Handler(spa)

	// init server
	srv := &http.Server{
		Handler:      ch(router),
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Metadata Server started @ port 3000")

	// run server
	log.Fatal(srv.ListenAndServe())
}
