//basic web server that serves file 03_static_page.html on directory static
// through the root endpoint

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
