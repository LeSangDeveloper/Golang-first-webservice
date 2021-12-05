package main

import (
	"net/http"
	"sangle.com/pluralsight/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
