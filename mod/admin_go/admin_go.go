package admingo

/*
admin standerd
all admin html pages must start with admin_
*/

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// WARNING: this will be changes for the futer
// but for now we will set the values like this
var login_pin string
var verb string
var cookie string

// simple coockie checker funciton
func Cookie_check(r *http.Request) bool {
	chk, er := r.Cookie("admin")

	if er != nil {
		return false
	} else if chk.Value == cookie {
		return true
	} else {
		return false
	}
}

// set the local data
func Set_data(user string, ver string) {
	login_pin = user
	cookie = user
	verb = ver
	//Cookie_gen()
}

// admin help functions
func Cookie_gen() {
	cookie = "admin_cookie"
}

// handle the login form
func Serv_login_page(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("templates/admin/html/admin_login.html")

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
		cookie_main := http.Cookie{Name: "admin", Value: cookie, Path: "/admin/", MaxAge: 1000} // Expires:}
		http.SetCookie(w, &cookie_main)
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
	if Cookie_check(r) {
		http.ServeFile(w, r, "templates/admin/html/admin_index.html")
	} else {
		http.Redirect(w, r, "/admin/login.html", 302)
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
	if Cookie_check(r) {
		http.ServeFile(w, r, "templates/admin/html/admin_file_view.html")
	} else {
		http.Redirect(w, r, "/admin/login.html", 302)
	}
}

// file download
func File_send(w http.ResponseWriter, r *http.Request) {
	if !Cookie_check(r) {
		http.Redirect(w, r, "/admin/login.html", 302)
	}
	// get the file name
	file_name := r.URL.Query().Get("file")

	// check if the file_name is null or not
	if file_name == "" {
		http.Redirect(w, r, "/admin/main/view_page", 302)
	}

	path := "templates/"
	// check what file type it is
	// and determin if it's admin or not and if it's html or css
	// we wont send img's or vid's as that would take up too much space and doesn't really matter
	if strings.Contains(file_name, ".html") {
		// its html but now check if its admin
		if strings.Contains(file_name, "admin") {
			path = path + "admin/html/" + file_name
			// its admin
		} else {
			// not admin
			path = path + "html/" + file_name
		}
	} else if strings.Contains(file_name, ".css") {
		// its css but now check if its admin
		if strings.Contains(file_name, "admin") {
			// its admin
			path = path + "admin/css/" + file_name
		} else {
			// not admin
			path = path + "css/" + file_name
		}

	}

	// read and send the file to the user
	send_file, err := os.Open(path)
	defer send_file.Close()
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found. Redicrecting...", 404)
		http.Redirect(w, r, "/admin/main/view_page", 302)
	}
	// max file size for start will be 1mb, html and css files should never reach this size tho
	file := make([]byte, 1024)
	send_file.Read(file)

	file_content := http.DetectContentType(file)

	// get the file size
	FileStat, _ := send_file.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)
	// send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+file_name)
	w.Header().Set("Content-Type", file_content)
	w.Header().Set("Content-Length", FileSize)
	// send the file
	// we read 1024 bytes from the file already, so we reset the offset back to 0
	send_file.Seek(0, 0)
	// send the file
	io.Copy(w, send_file)
}

// file downloader
func update_file(data string, name string) {
	if _, err := os.Stat(name); err == nil {
		file, err := os.Create(name)

		if err != nil {
			fmt.Println("[!!] ", err)
		} else {
			file.WriteString(data)
		}
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("[!!] " + name + " does not exsist!")
	}
}

func Donwload(w http.ResponseWriter, r *http.Request) {
	if !Cookie_check(r) {
		http.Redirect(w, r, "/admin/login.html", 302)
	}

	// send the upload form to the user
	http.ServeFile(w, r, "templates/admin/html/admin_send_file_form.html")
	// serv the templates/admin/css/main.css file to the user
	http.ServeFile(w, r, "templates/admin/css/main.css")

}

func Get_download(w http.ResponseWriter, r *http.Request) {
	if !Cookie_check(r) {
		http.Redirect(w, r, "/admin/login.html", 302)
	}

	// get the new file
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	// work out the file path
	file_name := handler.Filename

	path := "templates/"
	if strings.Contains(file_name, ".html") {
		// its html but now check if its admin
		tmp := strings.Split(file_name, ".html")
		file_name = tmp[0] + ".html"
		if strings.Contains(file_name, "admin") {
			path = path + "admin/html/" + file_name
			// its admin
		} else {
			// not admin
			path = path + "html/" + file_name
		}
	} else if strings.Contains(file_name, ".css") {
		// its css but now check if its admin
		tmp := strings.Split(file_name, ".css")
		file_name = tmp[0] + ".css"
		if strings.Contains(file_name, "admin") {
			path = path + "admin/html/"
			// its admin
		} else {
			// not admin
			path = path + "html/"
		}
	}
	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	//fmt.Printf("File Size: %+v\n", handler.Size)
	//fmt.Printf("MIME Header: %+v\n", handler.Header)
	// write the file to the server
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	// redirect the user to the admin page
	http.Redirect(w, r, "/admin/main/view_page", 302)

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

var html []string
var css []string
var admin_html []string
var admin_css []string

func File_code_gen(w http.ResponseWriter, r *http.Request) {
	if !Cookie_check(r) {
		http.Redirect(w, r, "/admin/login.html", 302)
	}

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
	// simple div code gen for the file viwer
	return `
	<button type="button" class="collapsible" onclick="pop_js()">` + file_name + `</button>
	<div class="content">
		<p>Download or Upload? Click the buttons below</p>
		<a href="/admin/files_upload/html/">Upload?</a>
		<a href="/admin/files_download/html/?file=` + file_name + `">Download?</a>
	</div>
	`
}

// TODO: make it simpler
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
				<a href="/admin/files_upload/html/(file name)">Upload?</a>
				<a href="/admin/files_download/html/(file name)">Download?</a>
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
