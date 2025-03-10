package main

import (
	"fmt"
	"io"
	"net/http"
)

func headers_one(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func headers_two(w http.ResponseWriter, r *http.Request) {
	for field, content := range r.Header {
		fmt.Fprintf(w, "%v: %v\n", field, content)
	}
}

func head_with_body(w http.ResponseWriter, r *http.Request) {
	for field, content := range r.Header {
		fmt.Fprintf(w, "%v: %v\n", field, content)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error trying to read body of request", http.StatusInternalServerError)
		fmt.Fprintf(w, "Error trying to read body of request: %v\n", err)
		return
	} else {
		fmt.Fprintf(w, "\n%s\n", string(body))
	}

	r.Body.Close()

}

func main() {
	http.HandleFunc("/headers1", headers_one)
	http.HandleFunc("/headers2", headers_two)
	http.HandleFunc("/head_with_body", head_with_body)
	http.ListenAndServe(":9000", nil)
}
