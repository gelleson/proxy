package http

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type Config struct {
	Host              string
	SetHostFromConfig bool
}

type Proxy struct {
	verbose bool
	config  Config
	client  *fasthttp.Client
}

func New(verbose bool, config Config) *Proxy {
	return &Proxy{
		client:  &fasthttp.Client{},
		verbose: verbose,
		config:  config,
	}
}

func (p Proxy) Proxy(request *fasthttp.Request, response *fasthttp.Response) error {
	if p.config.Host != "" {
		request.URI().SetHost(p.config.Host)
	}

	if p.verbose {
		fmt.Println(request.Header.String())

		fmt.Println(string(request.Body()))
	}

	return p.client.Do(request, response)
}
