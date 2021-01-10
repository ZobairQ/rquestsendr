package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link

}

/*
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
*/
