package main

import (
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request)  {
    // Display home page
    RenderTemplate(w, r, "index/home", nil)
}