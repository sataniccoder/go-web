package main

import (
	"fmt"
	admin "go-web/mod/admin_go"
	config "go-web/mod/config"
	get "go-web/mod/get_handle"
	pin "go-web/mod/util_go"
	"log"
	"net/http"
)

func main() {
	// set everythign up
	// setupt the admin pin
	admin_pin := pin.Gen_code()
	// read the config
	data_list := config.Config_reader()

	// set verb for all files needing it
	get.Verb_update(data_list[9])
	admin.Set_data(admin_pin, data_list[9])
	fmt.Println(`[*] Starting up server...
	 d888b   .d88b.         db   d8b   db d88888b d8888b. 
	88' Y8b .8P  Y8.        88   I8I   88 88'     88   8D 
	88      88    88        88   I8I   88 88ooooo 88oooY' 
	88  ooo 88    88 C8888D Y8   I8I   88 88~~~~~ 88~~~b. 
	88. ~8~  8b  d8'         8b d8'8b d8' 88.     88   8D 
	 Y888P    Y88P'           8b8' '8d8'  Y88888P Y8888P' 

					VERSION: 0.3


HOW TO STORE:
	html -> go-web/templates/html
	css  -> go-web/templates/css
	img  -> go-web/templates/img
	vid  -> go-web/templates/vid (BEST TO USE EXTERNAL LINK!)


url: http://127.0.0.1:` + data_list[1] + `/index.html
adming url: http://127.0.0.1:` + data_list[1] + `/admin/login.html
adming pin: ` + admin_pin + ` (you can request another one on the admin portal)
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
	// admin file viewer functions
	http.HandleFunc("/admin/main/view_page", admin.File_viewr)
	http.HandleFunc("/admin/file_data_loader", admin.File_code_gen)

	// admin css handler
	http.HandleFunc("/admin/css/ ", admin.Css_handle)

	log.Fatal(http.ListenAndServe(":"+data_list[1], nil))
}
