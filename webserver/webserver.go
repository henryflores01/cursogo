package webserver

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func MiWebserver() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}
