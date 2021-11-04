
## Installation

    go get -d -v ./...
    go install -v ./...

## Available Scripts

In the project directory, you can run:

### `go run main`

Runs the app in the development mode.\
Open [http://localhost:4000](http://localhost:4000) to view it in the browser.

### `go test -run "^Test(G|P)"`

Launches the test runner for *Unit tests*

### `go test -run "(Contracts)$"`

Launches the test runner for *CDC Provider tests*\

**Note: Pact Broker is hosted at https://rmustafayev.pactflow.io**

### `go build .`

Builds the app for production

# Docker

### `docker build -t <image_name>:<tagname> .`

Builds the docker image from the app.

### `docker run -dp <port>:3000 <image_name>:<tagname>`

Hosts the application at the given port.

# CircleCI

## STEPS
* Run Unit Tests 
* Run CDC Tests 

**IF *`PASS`***

* Create Docker image of the app
* Deploy to GKE
