/*
 *  Copyright 2021 Data and Service Center for the Humanities - DaSCH.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/api/presenter"
	projectEntity "github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity/project"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/service/project"
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func createProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error creating project"
		var input struct {
			ShortCode   string `json:"shortCode"`
			ShortName   string `json:"shortName"`
			LongName    string `json:"longName"`
			Description string `json:"description"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// convert input strings to value objects
		sc, err := valueobject.NewShortCode(input.ShortCode)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		sn, err := valueobject.NewShortName(input.ShortName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		ln, err := valueobject.NewLongName(input.LongName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		desc, err := valueobject.NewDescription(input.Description)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		id, err := service.CreateProject(ctx, sc, sn, ln, desc)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// get the project
		p, err := service.GetProject(ctx, id)
		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No data was returned"))
			return
		}

		toJ := &presenter.Project{
			ID:          id,
			ShortCode:   p.ShortCode().String(),
			ShortName:   p.ShortName().String(),
			LongName:    p.LongName().String(),
			Description: p.Description().String(),
			CreatedAt:   p.CreatedAt().String(),
			CreatedBy:   p.CreatedBy().String(),
			ChangedAt:   p.ChangedAt().String(),
			ChangedBy:   p.ChangedBy().String(),
			DeletedAt:   p.DeletedAt().String(),
			DeletedBy:   p.DeletedBy().String(),
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating project"
		var input struct {
			ShortCode   string `json:"shortCode"`
			ShortName   string `json:"shortName"`
			LongName    string `json:"longName"`
			Description string `json:"description"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		// get variables from request url
		vars := mux.Vars(r)

		// create empty Identifier
		uuid := valueobject.Identifier{}

		// create byte array from the provided id string
		b := []byte(vars["id"])

		// assign the value of the Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get the project
		p, err := service.GetProject(ctx, uuid)
		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No data was returned"))
			return
		}

		// convert input strings to value objects
		sc, err := valueobject.NewShortCode(input.ShortCode)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		sn, err := valueobject.NewShortName(input.ShortName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		ln, err := valueobject.NewLongName(input.LongName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		desc, err := valueobject.NewDescription(input.Description)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// update the project
		up, err := service.UpdateProject(ctx, uuid, sc, sn, ln, desc)
		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err == projectEntity.ErrNoPropertiesChanged {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Project{
			ID:          up.ID(),
			ShortCode:   up.ShortCode().String(),
			ShortName:   up.ShortName().String(),
			LongName:    up.LongName().String(),
			Description: up.Description().String(),
			CreatedAt:   up.CreatedAt().String(),
			CreatedBy:   up.CreatedBy().String(),
			ChangedAt:   up.ChangedAt().String(),
			ChangedBy:   up.ChangedBy().String(),
			DeletedAt:   up.DeletedAt().String(),
			DeletedBy:   up.DeletedBy().String(),
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get variables from request url
		vars := mux.Vars(r)

		// create empty Identifier
		uuid := valueobject.Identifier{}

		// create byte array from the provided id string
		b := []byte(vars["id"])

		// assign the value of the Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get the project
		p, err := service.GetProject(ctx, uuid)
		w.Header().Set("Content-Type", "application/json")

		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No p was returned"))
			return
		}

		toJ := &presenter.Project{
			ID:          p.ID(),
			ShortCode:   p.ShortCode().String(),
			ShortName:   p.ShortName().String(),
			LongName:    p.LongName().String(),
			Description: p.Description().String(),
			CreatedAt:   p.CreatedAt().String(),
			CreatedBy:   p.CreatedBy().String(),
			ChangedAt:   p.ChangedAt().String(),
			ChangedBy:   p.ChangedBy().String(),
			DeletedAt:   p.DeletedAt().String(),
			DeletedBy:   p.DeletedBy().String(),
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed encoding p to JSON"))
		}
	})
}

func deleteProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get variables from request url
		vars := mux.Vars(r)

		// create empty Identifier
		uuid := valueobject.Identifier{}

		// create byte array from the provided id string
		b := []byte(vars["id"])

		// assign the value of the Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// delete the project
		data, err := service.DeleteProject(ctx, uuid)
		w.Header().Set("Content-Type", "application/json")

		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Project has already been deleted"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No data was returned"))
			return
		}

		toJ := &presenter.DeleteProject{
			ID:        data.ID(),
			DeletedAt: data.ChangedAt().String(),
			DeletedBy: data.ChangedBy().String(),
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed encoding data to JSON"))
		}
	})
}

func listProjects(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ReturnDeletedProjects bool `json:"returnDeletedProjects"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			input.ReturnDeletedProjects = false // default to false if decoding fails (likely because it wasn't provided)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get all project ids
		ids, err := service.ListProjects(ctx, input.ReturnDeletedProjects)
		w.Header().Set("Content-Type", "application/json")

		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if ids == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No ids was returned"))
			return
		}

		var toJ []presenter.Project

		for _, id := range ids {
			p, err := service.GetProject(ctx, id)
			if err != nil && err == projectEntity.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("No project found for this uuid"))
				return
			}

			if err != nil && err != projectEntity.ErrNotFound {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("The server is not responding"))
				return
			}
			if p == nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("No data was returned"))
				return
			}

			toJ = append(toJ, presenter.Project{
				ID:          p.ID(),
				ShortCode:   p.ShortCode().String(),
				ShortName:   p.ShortName().String(),
				LongName:    p.LongName().String(),
				Description: p.Description().String(),
				CreatedAt:   p.CreatedAt().String(),
				CreatedBy:   p.CreatedBy().String(),
				ChangedAt:   p.ChangedAt().String(),
				ChangedBy:   p.ChangedBy().String(),
				DeletedAt:   p.DeletedAt().String(),
				DeletedBy:   p.DeletedBy().String(),
			})
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed encoding ids to JSON"))
		}
	})
}

//MakeProjectHandlers make url handlers
func MakeProjectHandlers(r *mux.Router, n negroni.Negroni, service project.UseCase) {

	r.Handle("/v1/project", n.With(
		negroni.Wrap(createProject(service)),
	)).Methods("POST", "OPTIONS").Name("createProject")

	r.Handle("/v1/project/{id}", n.With(
		negroni.Wrap(updateProject(service)),
	)).Methods("PUT", "OPTIONS").Name("updateProject")

	r.Handle("/v1/project/{id}", n.With(
		negroni.Wrap(getProject(service)),
	)).Methods("GET", "OPTIONS").Name("getProject")

	r.Handle("/v1/project/{id}", n.With(
		negroni.Wrap(deleteProject(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteProject")

	r.Handle("/v1/projects/all", n.With(
		negroni.Wrap(listProjects(service)),
	)).Methods("GET", "OPTIONS").Name("listProjects")
}