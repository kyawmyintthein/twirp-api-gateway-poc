package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyawmyintthein/twirp-api-gateway-poc/rpc/color"
	"github.com/twitchtv/twirp"
	"golang.org/x/net/http2"
)

func main() {
	r := gin.Default()
	colorService := color.NewColorServiceProtobufClient("https://localhost:8081", &http.Client{
		Transport: transport2(),
	}, twirp.WithClientPathPrefix("rz"))
	r.GET("/color/random", func(c *gin.Context) {
		resp, err := colorService.GetRandomColor(context.Background(), &color.GetRandomColorRequest{Count: 10})
		if err != nil {
			c.JSON(400, gin.H{"message": "bad request"})
			return
		}
		c.JSON(200, resp)
		return
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func transport2() *http2.Transport {
	return &http2.Transport{
		TLSClientConfig:    tlsConfig(),
		DisableCompression: true,
		AllowHTTP:          false,
	}
}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("../../server.crt")
	if err != nil {
		log.Fatal(err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}
}
