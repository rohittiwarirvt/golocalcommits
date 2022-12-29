package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()

	// create HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://via.placeholder.com/2000x2000.png", nil)

	if err != nil {
		panic(err)
	}

	// perform http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// get data from HTTPresponse

	imageData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("downloaded image of size %d\n", len(imageData))
}
