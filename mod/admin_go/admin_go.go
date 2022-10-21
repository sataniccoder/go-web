package admingo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// WARNING: this will be changes for the futer
// but for now we will set the values like this
var login_pin string
var verb string
var cookie string

// set the local data
func Set_data(user string, ver string) {
	login_pin = user
	verb = ver
	cookie_gen()
}

// admin help functions
func cookie_gen() {
	cookie = "admin_cookie"
}

// handle the login form
func Serv_login_page(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("templates/admin/html/login.html")

	// if there is an error then the page either doesn't exisst ot there's an internal error
	// best to make the user belive it doesn't exsist and return the accchtual error
	if err != nil {
		fmt.Println("error: ", err)
		fmt.Fprintf(w, "the page your looking for doesn't exist!")
	} else {
		fmt.Fprintf(w, string(content))
	}
}

func Handle_login(w http.ResponseWriter, r *http.Request) {
	// this base path will be for the html/ and css/ in the admin menu
	//base_path := "templates/admin/"
	//

	if verb == "1" {
		fmt.Println("got login request!")
	}

	name := r.FormValue("pin")

	if name == login_pin {
		if verb == "1" {
			fmt.Println("pin!")
		}
		// send the main admin.html page and send a login auth cookie
		cookie_main := http.Cookie{Name: "admin", Value: cookie, Path: "/", MaxAge: 1800} //, Expires: Nil}
		http.SetCookie(w, &cookie_main)
		http.Redirect(w, r, "/admin/admin_index", 301)

		// this will be loaded in case the user doesn't get redirected
		redir := `
<p>you are being redirected...</p>
<p>if it takes longer then a couple of seonds then please click the button below :)</p>
<a href="/admin/admin_index">
	<button>go-web!</button>
</a>
		`

		fmt.Fprintf(w, redir)
	} else {
		error_pas := `
		<!DOCTYPE html>
		<html>
		   <head>
			  <title>HTML Meta Tag</title>
			  <meta http-equiv = "refresh" content = "3; url = /admin/login.html" />
		   </head>
		   <body>
			  <p>Redirecting to another URL</p>
		   </body>
		</html>
		`
		fmt.Fprintf(w, error_pas)
	}
}

// main admin page
func Admin_page(w http.ResponseWriter, r *http.Request) {
	chk, er := r.Cookie("admin")

	if er != nil {
		http.Redirect(w, r, "/admin/login.html", 302)
	}

	if chk.Value == cookie {
		http.ServeFile(w, r, "templates/admin/html/admin_index.html")
	} else {
		http.Redirect(w, r, "/admin/login.html", 301)
	}
}

// admin css support
func Css_handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	path += "/templates/" + path

	if verb == "1" {
		fmt.Println("path: ", path)
	}
}

// admin file viewer
func File_viewr(w http.ResponseWriter, r *http.Request) {
	chk, er := r.Cookie("admin")

	if er != nil {
		http.Redirect(w, r, "/admin/login.html", 302)
	}

	if chk.Value == cookie {
		http.ServeFile(w, r, "templates/admin/html/admin_file_view.html")
	} else {
		http.Redirect(w, r, "/admin/login.html", 301)
	}
}

var html []string
var css []string
var admin_html []string
var admin_css []string

func File_code_gen(w http.ResponseWriter, r *http.Request) {
	var code string

	html = Get_list(html, "templates/html")
	css = Get_list(css, "templates/css")
	admin_css = Get_list(admin_css, "templates/admin/css")
	admin_html = Get_list(admin_html, "templates/admin/html")

	code += Gen_html_code()

	fmt.Fprintf(w, code)
}

// code for html file
func Get_list(lis []string, fil string) []string {
	files, err := ioutil.ReadDir(fil)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {
			if !file.IsDir() {
				lis = append(lis, file.Name())
			}
		}
	}

	return lis
}

func Gen_div_code(file_name string) string {
	return `
	<button type="button" class="collapsible" onclick="pop_js()">` + file_name + `</button>
	<div class="content">
		<p>Download or Upload? Click the buttons below</p>
		<a href="/files_upload/html/` + file_name + `">Upload?</a>
		<a href="/files_download/html/` + file_name + `">Download?</a>
	</div>
	`
}

func Gen_html_code() string {
	var html_code string

	/*
		the html code should look like this for each list defined at the start of the program
		the esiest way i can think of doing this is chopping it up into smaller functions to make the code esier to read and maintain

		<button type="button" class="collapsible" onclick="pop_js()">type (admin at start if admin)</button>
		<div class="content">
			<button type="button" class="collapsible" onclick="pop_js()">(file name)</button>
			<div class="content">
				<p>Download or Upload? Click the buttons below</p>
				<a href="/files_upload/html/(file name)">Upload?</a>
				<a href="/files_download/html/(file name)">Download?</a>
			</div>
			+more for each file
		</div>
	*/

	// gen html code
	html_code += `<button type="button" class="collapsible" onclick="pop_js()">html</button>
	<div class="content">`
	for _, file := range html {
		html_code += Gen_div_code(file)
	}
	html_code += "</div>"
	// gen css code
	html_code += `<button type="button" class="collapsible" onclick="pop_js()">css</button>
	<div class="content">`
	for _, file := range css {
		html_code += Gen_div_code(file)
	}
	html_code += "</div>"
	// gen admin html
	html_code += `<button type="button" class="collapsible" onclick="pop_js()">admin</button>
	<div class="content">
	<button type="button" class="collapsible" onclick="pop_js()">html</button>
	<div class="content">
	`
	for _, file := range admin_html {
		html_code += Gen_div_code(file)
	}
	html_code += "</div>"
	html_code += `<button type="button" class="collapsible" onclick="pop_js()">css</button>
	<div class="content">
	`
	for _, file := range admin_html {
		html_code += Gen_div_code(file)
	}
	html_code += "</div></div>"

	// gen admin css

	return html_code
}
