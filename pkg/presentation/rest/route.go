package rest

import (
	"drello-api/pkg/constants"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type handler func(http.ResponseWriter, *http.Request)

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logHTTPRequest(r)
	fn(w, r)
}

func HandleRequests() {
	fmt.Println("Listening on http://127.0.0.1:8080")

	r := mux.NewRouter()
	r.Handle(constants.Workspaces+"/{id:[0-9]+}", handler(workspaceHandler))
	r.Handle(constants.Workspaces, handler(workspacesHandler))
	log.Fatal(http.ListenAndServe(":8080", r))
}
