package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		// create HTTP request
		req, err := http.NewRequestWithContext(ctx.Request.Context(), http.MethodGet, "https://yahoo.com", nil)

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

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		ctx.Data(200, "text/html", data)
	})

	r.Run("localhost:8080")
}
