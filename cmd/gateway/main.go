package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	gologging "github.com/devopsfaith/krakend-gologging"
	viper "github.com/devopsfaith/krakend-viper"
	"github.com/gin-gonic/gin"
	luratwirp "github.com/kyawmyintthein/lura-twirp"
	_ "github.com/kyawmyintthein/twirp-api-gateway-poc/pkg/plugins/bodymodifier"
	_ "github.com/kyawmyintthein/twirp-api-gateway-poc/pkg/plugins/header2body"
	_ "github.com/kyawmyintthein/twirp-api-gateway-poc/pkg/plugins/querystring2body"
	color "github.com/kyawmyintthein/twirp-api-gateway-poc/rpc/color"
	number "github.com/kyawmyintthein/twirp-api-gateway-poc/rpc/number"
	"golang.org/x/net/http2"

	"github.com/luraproject/lura/proxy"
	krakendgin "github.com/luraproject/lura/router/gin"
	"github.com/luraproject/lura/transport/http/client"
	"github.com/luraproject/lura/transport/http/server"
	"github.com/twitchtv/twirp"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	debug := flag.Bool("d", false, "Enable the debug")
	configFile := flag.String("c", "./configuration.json", "Path to the configuration filename")
	flag.Parse()

	parser := viper.New()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, err := gologging.NewLogger(serviceConfig.ExtraConfig, os.Stdout)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	logger.Debug("config:", serviceConfig)

	ctx, cancel := context.WithCancel(context.Background())

	//backendFactory := martian.NewBackendFactory(logger, client.DefaultHTTPRequestExecutor(client.NewHTTPClient))
	colorLuraClient, err := color.NewColorServiceLuraClient(&serviceConfig, "rpc.color.ColorService", &http.Client{
		Transport: transport2(),
	}, logger, twirp.WithClientPathPrefix("rz"))
	if err != nil {
		panic(err)
	}

	numberLuraClient, err := number.NewNumberServiceLuraClient(&serviceConfig, "rpc.number.NumberService", &http.Client{
		Transport: transport2(),
	}, logger, twirp.WithClientPathPrefix("rz"))
	if err != nil {
		panic(err)
	}
	luratwirp.RegisterTwirpStubs(logger, colorLuraClient, numberLuraClient)

	bf := luratwirp.NewTwirpProxy(logger, client.DefaultHTTPRequestExecutor(client.NewHTTPClient))
	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         gin.Default(),
		Logger:         logger,
		Middlewares:    []gin.HandlerFunc{},
		HandlerFactory: krakendgin.EndpointHandler,
		ProxyFactory:   proxy.NewDefaultFactory(bf, logger),
		RunServer:      server.RunServer,
	})
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	routerFactory.NewWithContext(ctx).Run(serviceConfig)

	cancel()
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
