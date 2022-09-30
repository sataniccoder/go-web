package admingo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// WARNING: this will be changes for the futer
// but for now we will set the values like this
var login_user string
var login_pass string
var verb string

// set the local data
func Set_user(user string) {
	login_user = user
}
func Set_pass(pass string) {
	login_pass = pass
}
func Set_verb(ver string) {
	verb = ver
}

// admin help functions
func cookie_gen() string {
	return "admin_cookie"
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

	name := r.FormValue("name")
	pass := r.FormValue("pswd")

	if name == "test" && pass == "test" {
		// send the main admin.html page and send a login auth cookie
		cookie := http.Cookie{Name: "admin", Value: "admin_true", Path: "/", MaxAge: 1800} //, Expires: Nil}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/admin/admin_index", 302)

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

	fmt.Println("cookie: ", chk.Value)
	if chk.Value == "admin_true" {
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
