package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := NewRouter()

	router.GET("/", HandleHome)
	// Add the route handler
	// router.Handle("POST", "/register", HandleUserCreate)

	router.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)

	middleware := Middleware{}
	middleware.Add(router)
	log.Fatal(http.ListenAndServe(":3000", middleware))
}

// NewRouter creates a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	// TODO: need to be revisited -> outdated code
	router.NotFound = func(http.ResponseWriter, *http.Request, _ httprouter.Params) {}
	return router
}
