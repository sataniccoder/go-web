package main

import (
	"fmt"
	admin "go-web/mod/admin_go"
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
	get.Verb_update(data_list[9])
	fmt.Println(`[*] Starting up server...
goweb! a simple static web hosting service in golang!
VERSION: 0.1


HOW TO STORE:
	html -> go-web/templates/html
	css  -> go-web/templates/css
	img  -> go-web/templates/img
	vid  -> go-web/templates/vid (BEST TO USE EXTERNAL LINK!)


url: http://127.0.0.1:` + data_list[1] + `/index.html
adming url: http://127.0.0.1:` + data_list[1] + `/admin/login.html
	`)
	http.HandleFunc("/", get.Load_page)
	// put POST url's here

	// add ADMIN url's here (can be get and post)
	// admin login page loader
	http.HandleFunc("/admin/login.html", admin.Serv_login_page)
	// get the details
	http.HandleFunc("/admin/main/admin_set", admin.Handle_login)

	// the main admin pages
	http.HandleFunc("/admin/admin_index", admin.Admin_page)

	// admin css handler
	http.HandleFunc("/admin/css/ ", admin.Css_handle)

	log.Fatal(http.ListenAndServe(":"+data_list[1], nil))
}

/*
NOTES:
add admin .css suport
Add cookie support for admin custmizeation
upgrade all get functions to use the servefile function

help:

to add POST url's

add another http.HandleFunc("/path/of/url", post.Post_function_name)
*/
