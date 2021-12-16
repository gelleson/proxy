package http

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

type Config struct {
	Schema            string
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
		client: &fasthttp.Client{
			ReadTimeout:  time.Second * 40,
			WriteTimeout: time.Second * 40,
		},
		verbose: verbose,
		config:  config,
	}
}

func (p Proxy) Proxy(request *fasthttp.Request, response *fasthttp.Response) error {
	if p.config.Host != "" {
		request.URI().SetHost(p.config.Host)
	}

	if p.config.Schema != "" {
		request.URI().SetScheme(p.config.Schema)
	} else {
		request.URI().SetScheme("https")
	}

	if p.verbose {
		fmt.Println(request.Header.String())
		fmt.Println("")
		fmt.Println("Body: ")
		fmt.Println(string(request.Body()))
	}

	if err := p.client.Do(request, response); err != nil {
		return err
	}

	return nil
}
