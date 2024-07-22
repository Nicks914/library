package router

import (
	"libary/controllers"
	"net/http"

	"strings"
)

func Routes(w http.ResponseWriter, r *http.Request) {
	// Trailing slash is a pain in the ass so we just drop it
	route := strings.Trim(r.URL.Path, "/")
	switch route {
	case "login":
		controllers.Login(w, r)

	case "home":
		controllers.HomeHandler(w, r)
	case "addBook":
		controllers.AddBookHandler(w, r)
	case "deleteBook":
		controllers.DeleteBookHandler(w, r)

	default:
		controllers.Forbidden(w, r)

	}
}
