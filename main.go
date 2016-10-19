package main

import (
	"net/http"
	"github.com/sKudryashov/social_event_api_prototype/app"
)

func main() {
	router := app.InitApp()
	http.ListenAndServe(":3030", router.Serve())
}



