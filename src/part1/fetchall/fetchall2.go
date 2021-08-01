package fetchall

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Fetchall2() {
	start := time.Now()
	ch := make(chan string)
	for _, arg := range os.Args[1:] {
		go fetchURL(arg, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("\n%.2f segundos\n", time.Since(start).Seconds())
}

func fetchURL(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
	}

	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
	}

	ch <- fmt.Sprintf("%d - %.2f secondsg", nbytes, time.Since(start).Seconds())
}
