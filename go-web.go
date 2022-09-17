package main

import (
	"fmt"
	config "go-web/mod/config"
	get "go-web/mod/get_handle"
	"log"
	"net/http"
)

func main() {
	// set everythign up
	// read the config
	data_list := config.Config_reader()

	// set verb for all files needing it
	get.Verb_update(data_list[8])

	fmt.Println(`[*] Starting up server...
goweb! a simple static web hosting service in golang!
VERSION: 0.1


HOW TO STORE:
	html -> go-web/templates/html
	css  -> go-web/templates/css
	img  -> go-web/templates/img
	vid  -> go-web/templates/vid (BEST TO USE EXTERNAL LINK!)


url: http://127.0.0.1:` + data_list[1] + `/index.html
	`)
	http.HandleFunc("/", get.Load_page)
	log.Fatal(http.ListenAndServe(":"+data_list[1], nil))
}

