/*
 * Digimon Service API
 *
 * 提供孵化數碼蛋與培育等數碼寶貝養成服務
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/",
		Index,
	},

	Route{
		"DigimonsDigimonIDFosterPost",
		strings.ToUpper("Post"),
		"/api/v1/digimons/{digimonID}/foster",
		DigimonsDigimonIDFosterPost,
	},

	Route{
		"DigimonsDigimonIDGet",
		strings.ToUpper("Get"),
		"/api/v1/digimons/{digimonID}",
		DigimonsDigimonIDGet,
	},

	Route{
		"DigimonsPost",
		strings.ToUpper("Post"),
		"/api/v1/digimons",
		DigimonsPost,
	},
}
