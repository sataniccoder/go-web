package main

import(
	"fmt"
//    "html" 
	"net/http"
	"log"
	"strings"
	"io/ioutil"
)


func load_page(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:] 

	fmt.Println("page: ", path)

	if strings.Contains(path, ".html"){
		fmt.Println("loading html...")
		path = "templates/html/"+path
		content, err := ioutil.ReadFile(path)

		if err != nil{
			fmt.Println("error: ", err)
		} else{
			fmt.Fprintf(w, string(content))
		}
	} else if strings.Contains(path, ".css"){
		fmt.Println("loading css...")
		path = "templates/css/"+path
		content, err := ioutil.ReadFile(path)

		if err != nil{
			fmt.Println("error: ", err)
		} else{
			fmt.Fprintf(w, string(content))
		}
	}
}

func main(){
	fmt.Println(`
goweb! a simple static web hosting service in golang!


HOW TO STORE:
	html -> go-web/templates/html
	css  -> go-web/templates/css

url: http://127.0.0.1:8080/index.html
	`)
	http.HandleFunc("/", load_page)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
