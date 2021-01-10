package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://golang.org/pkg")
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}

		io.Copy(w, resp.Body)
	})
	http.ListenAndServe(":8080", nil)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
