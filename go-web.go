package main

import(
	"fmt"
//    "html" 
	"net/http"
	"log"
	"strings"
	"io/ioutil"
)

// load_page, loads page from the templates folder
func load_page(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:] 

	fmt.Println("GET page: ", path)

	if strings.Contains(path, ".html"){
		fmt.Println("loading html...")
		path = "templates/html/"+path
		content, err := ioutil.ReadFile(path)

		if err != nil{
			fmt.Println("error: ", err)
			fmt.Fprintf(w, "the page your looking for doesn't exist!")
		} else{
			fmt.Fprintf(w, string(content))
		}
	} else if strings.Contains(path, ".css"){
		fmt.Println("loading css...")
		path = "templates/css/"+path
		content, err := ioutil.ReadFile(path)

		if err != nil{
			fmt.Println("error: ", err)
			fmt.Fprintf(w, "the page your looking for doesn't exist!")
		} else{
			fmt.Fprintf(w, string(content))
		}
	} else if strings.Contains(path, "img"){
		path = "templates/"+path

		fmt.Println("loading image...")
		content, err := ioutil.ReadFile(path)

		if err != nil{
			fmt.Println("error: ", err)
			fmt.Fprintf(w, "the page your looking for doesn't exist!")
		} else{
			w.Write(content)
		}
	} else if strings.Contains(path, "vid"){
		path = "templates/"+path
		fmt.Println("loading vidoe...")

		http.ServeFile(w, r, path)
	} else{
		fmt.Fprintf(w, "the page your looking for doesn't exist!")
	}
}


func main(){
	fmt.Println(`
goweb! a simple static web hosting service in golang!
VERSION: 0.1


HOW TO STORE:
	html -> go-web/templates/html
	css  -> go-web/templates/css
	img  -> go-web/templates/img
	vid  -> go-web/templates/vid (BEST TO USE EXTERNAL LINK!)

url: http://127.0.0.1:8080/index.html
	`)
	http.HandleFunc("/", load_page)
        log.Fatal(http.ListenAndServe(":8080", nil))

}

/*
config example:



*/
