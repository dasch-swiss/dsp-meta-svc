# Repository Archived

We would like to inform you that this GitHub repository has been archived and is no longer maintained.

Please use the following repository for the latest updates and continued support: [dasch-swiss/dsp-meta](https://github.com/dasch-swiss/dsp-meta).

Thank you for your understanding and cooperation.

---

# Metadata service

Service for providing users with project and dataset specific metadata.

## Requirements

The following programs need to be installed to run the metadata browser locally:

- Go
- Node
- Yarn
- Bazel/Bazelisk
- Docker

Ensure that all dependencies are installt, running `yarn install`.

## Run locally

To run both front-end and back-end locally, there are two options:

- `make metadata` (requires Node, Yarn and Go installed)
- `make metadata-docker-run` (requires Docker and Bazel installed)

Neither of those commands will update automatically when changes are made to the data served by the backend.
In this case, manually stop and restart the service using the same command.

To auto-rebuild after changes were made to the frontend, you can in a separate terminal run `yarn run dev`.
However, the changes will only apply when you run the service without docker (i.e. `make metadata`),
if you run it inside docker, you will have to re-run the service for the changes to apply.

## Front-end

The front-end part is basing on [Svelte](https://svelte.dev). To run it, `yarn` and/or `make` need to be installed.

1. Clone the repository:

```
https://github.com/dasch-swiss/dasch-service-platform.git
```

2. Install the dependencies:

```bash
yarn install
```

or

```bash
make yarn
```

3. Start the application:

```bash
yarn run dev
```

This starts the application on the [localhost:5000](http://localhost:5000).

_Note that you will also need to have [Node.js](https://nodejs.org) installed._

## Server

For now, metadata is served from a simple server written in Go.

The server serves the frontend (static file serving on `./public/`) and the metadata on `http://localhost:3000/`.

The route `/` serves the frontend.

The routes `/projects` and `/projects/:id` form a simple metadata API.  
The server serves all data found in `./services/metadata/backend/data/*.json`, where the JSON file follows the data structure as currently provided by DSP-JS-LIB.  
**Note:** Files starting with underscore (`_`) are excluded. This provides a simple means to leave out files that are not supposed to be public.  
The server supports pagination and full text search.

To run the server locally, use the command `make metadata`.  
To run, build and publish a docker image of the server, use the commands `metadata-docker-run`, `metadata-docker-build` and `metadata-docker-publish`respectively. (`...-run` will build first.)

## Go dependencies

The Go dependencies are defined inside the `go.mod` and the corresponding `go.sum` files.
To be able to use external dependencies with Bazel, all external dependencies need to be registered with Bazel
so that they can be referenced in the `BUILD.bazel` files.

The steps to add an external dependency are as follows:

1. to add repository and version to `go.mod`, run `go get gihub.com/stretchr/testify`
   (exchange the name of the package with the one you would like to add)
1. from inside this directory, run `go mod download gihub.com/stretchr/testify`
   (exchange the name of the package with the one you would like to add)
1. from the root of the repository, run `make gen-go-deps`

Running the `make gen-go-deps` will regenerate the `deps.bzl` file found
in the root of the repository. This file is loaded inside the `WORKSPACE` file, so that Bazel
can go and fetch the external repositories and make them available to be used in rules.

From then on, you should be able to use the external dependency in a `BUILD.bazel` file like so:

```bazel
go_test(
    name = "entity_test",
    srcs = ["organization_test.go"],
    embed = [":entity"],
    deps = ["@com_github_stretchr_testify//assert"],
)
```
