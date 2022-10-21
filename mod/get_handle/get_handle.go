package gethandle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// add all needed global verbales in here
var verb string

// get/post function handlers, this section handles all http server requests
func Verb_update(verbrose string) {
	verb = verbrose
}

// load_page, loads page from the templates folder
func Load_page(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if path == "" {
		// means it's the index path
		if verb == "1" {
			fmt.Println("GET page: index.html")
		}
		http.ServeFile(w, r, "templates/html/index.html")
	}

	// uses verb as verbrose, allows for cleaner and easier to understand output
	// to edit this change the config.conf setting
	if verb == "1" {
		fmt.Println("GET page: ", path)
	}

	// return html files
	// this if statement is the same for the others just replace "html" with what ever the .Contains is looking for
	if strings.Contains(path, ".html") {
		if verb == "1" {
			fmt.Println("loading html...")
		}

		path = "templates/html/" + path
		http.ServeFile(w, r, path)

	} else if strings.Contains(path, ".css") {
		if verb == "1" {
			fmt.Println("loading css...")
		}
		if strings.Contains(path, "admin/") {
			path = "templates/" + path
		} else {
			path = "templates/css/" + path
		}

		fmt.Println(path)
		http.ServeFile(w, r, path)

	} else if strings.Contains(path, "img") {
		path = "templates/" + path

		if verb == "1" {
			fmt.Println("loading image...")
		}

		content, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Println("error: ", err)
			fmt.Fprintf(w, "the page your looking for doesn't exist!")
		} else {
			w.Write(content)
		}
	} else if strings.Contains(path, "vid") {
		// this one is abit diffrent, it uses http.ServeFile to give the user thier vidoe data
		path = "templates/" + path

		if verb == "1" {
			fmt.Println("loading vidoe...")
		}

		http.ServeFile(w, r, path)
	} else {
		fmt.Fprintf(w, "the page your looking for doesn't exist!")
	}
}
