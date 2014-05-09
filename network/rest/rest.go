package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getUser(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	uid := params.Get("uid")
	fmt.Fprintf(writer, "GET user: uid='%s'", uid)
}

func modifyUser(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	uid := params.Get("uid")
	fmt.Fprintf(writer, "POST user: uid='%s'", uid)
}

func deleteUser(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	uid := params.Get("uid")
	fmt.Fprintf(writer, "DELETE user: uid='%s'", uid)
}

func addUser(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	uid := params.Get("uid")
	fmt.Fprint(writer, "PUT user: uid='%s'", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/", getUser)
	mux.Post("/user/", modifyUser)
	mux.Del("/user/", deleteUser)
	mux.Put("/user/", addUser)
	http.Handle("/", mux)
	http.ListenAndServe(":8888", nil)
}
