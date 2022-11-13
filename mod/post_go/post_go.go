package postgo

import (
	"fmt"
	"net/http"
)

func Main_Post(w http.ResponseWriter, r *http.Request) {
	// put your function url in here like this
	// the reason why we do this is to make it easier for the end user (you) to add post urls
	// it makes the code cleaner and means you only have to handle one go file instead of a bunch
	// plus it still aloows you too (if you really wanted too) add your own post urls in the go-web.go file
	// REMBER: every new post url you add you must recompile the code for it to be reconized, you can also do this in the admin pannel

	// you can remove this and add your own post urls here
	if r.URL.Path == "/post/example" {
		post_example(w, r)
	}

}

// you can remove this function and add your own post url functions here
func post_example(w http.ResponseWriter, r *http.Request) {
	// example post function
	fmt.Fprintf(w, "hello world")
}
