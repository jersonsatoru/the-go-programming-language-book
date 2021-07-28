package fetchall

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func FetchAll() {
	ch := make(chan string)

	start := time.Now()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("\nSeconds %3.2f\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, errHttp := http.Get(url)
	if errHttp != nil {
		ch <- fmt.Sprintf("Seconds %3.2f - %v", time.Since(start).Seconds(), errHttp)
		return
	}

	defer resp.Body.Close()

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Seconds %3.2f - %v", time.Since(start).Seconds(), err)
		return
	}

	ch <- fmt.Sprintf("Seconds %3.2f - Length %7d", time.Since(start).Seconds(), nbytes)
}
