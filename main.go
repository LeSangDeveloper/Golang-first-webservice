package main

import (
	"net/http"
	"sangle.com/pluralsight/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":7000", nil)
}
