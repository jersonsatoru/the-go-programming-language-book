package curl

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Curl() {
	urls := os.Args[1:]
	for _, url := range urls {
		resp, errCurl := http.Get(url)
		if errCurl != nil {
			fmt.Fprintf(os.Stderr, "%v", errCurl)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}

		fmt.Printf("%s", b)
	}
}

func Curl2() {
	urls := os.Args[1:]
	for _, url := range urls {
		resp, errCurl := http.Get(url)
		if errCurl != nil {
			fmt.Fprintf(os.Stderr, "%v", errCurl)
		}

		b, err := io.Copy(os.Stderr, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}

		fmt.Printf("%b", b)
	}
}
