/////User identity management system/////
package main

import (
	"net/http"
	"vijay/project-1/api/router"
)

func main() {
	route := router.New()
	http.ListenAndServe(":8080", route)
}
