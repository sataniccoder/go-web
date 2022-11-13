package main

import (
	"fmt"
	admin "go-web/mod/admin_go"
	config "go-web/mod/config"
	get "go-web/mod/get_handle"
	post "go-web/mod/post_go"
	pin "go-web/mod/util_go"
	"log"
	"net/http"
)

var verb string

func gen_new_pin(w http.ResponseWriter, r *http.Request) {
	// must be admin
	if admin.Cookie_check(r) {
		// admin we can now update the pin
		pi := pin.Gen_code()
		admin.Set_data(pi, verb)

		fmt.Fprintf(w, `
		<html>
		<head>
			<title>new pin!</title>
		</head>
		<body>
			<p>New Pin!</p>
			<p>PIN: `+pi+`</p>
			<button onclick="window.location.href='/admin/login.html'">login! (New pin)</button>
		</body>
	</html>
		
		`)
		fmt.Print("[\033[31m!!\033[0m]")
		fmt.Println(" Warning! new pin genrated! if this was not you press CTRL+C now!")
		fmt.Print("[\033[31m!!\033[0m]")
		fmt.Println(" or genrate a new pin! the new pin is (" + pi + ")")
	} else {
		http.Redirect(w, r, "/admin/login.html", 302)
	}
}

func main() {
	// set everythign up
	// setupt the admin pin
	admin_pin := pin.Gen_code()
	// read the config
	data_list := config.Config_reader()

	// set verb for all files needing it
	get.Verb_update(data_list[9])
	verb = data_list[9]
	admin.Set_data(admin_pin, verb)

	fmt.Print("\033[H\033[2J")
	fmt.Print(`

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
	fmt.Println()
	go http.HandleFunc("/", get.Load_page)
	// put POST url's here

	// add ADMIN url's here (can be get and post)
	// admin login page loader
	go http.HandleFunc("/admin/login.html", admin.Serv_login_page)
	// get the details
	go http.HandleFunc("/admin/main/admin_set", admin.Handle_login)

	// the main admin pages
	go http.HandleFunc("/admin/admin_index", admin.Admin_page)
	// genrate new pin
	go http.HandleFunc("/admin/new_pin", gen_new_pin)
	// admin file viewer functions
	go http.HandleFunc("/admin/main/view_page", admin.File_viewr)
	go http.HandleFunc("/admin/file_data_loader", admin.File_code_gen)

	// admin css handler
	go http.HandleFunc("/admin/css/ ", admin.Css_handle)

	// file download (client)
	go http.HandleFunc("/admin/files_download/html/", admin.File_send)
	go http.HandleFunc("/admin/files_upload/html/", admin.Donwload)
	go http.HandleFunc("/admin/upload", admin.Get_download)

	// user post url
	go http.HandleFunc("/post/", post.Main_Post)
	go log.Fatal(http.ListenAndServe(":"+data_list[1], nil))
}

/*
TODO: get file upload working


TODO: add documentation file
TODO: improve the admin UI to have the server stats on the main page

*/
