package http

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"time"
)

type Config struct {
	Schema            string
	Host              string
	SetHostFromConfig bool
}

type Proxy struct {
	verbose      bool
	config       Config
	client       *fasthttp.Client
	accessLogger io.Writer
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
		fmt.Println("Request: ")
		fmt.Println(request.Header.String())
		fmt.Println("")
		if len(request.Body()) > 0 {
			fmt.Println("Body: ")
			fmt.Println(string(request.Body()))
		} else {
			fmt.Println("Body: EMPTY")
		}
	}

	if p.verbose && p.accessLogger != nil {
		if _, err := p.accessLogger.Write([]byte(fmt.Sprintf("%s\n", request.Header.String()))); err != nil {
			return err
		}

		if _, err := p.accessLogger.Write([]byte(fmt.Sprintf("%s\n", request.Body()))); err != nil {
			return err
		}
	}

	if err := p.client.Do(request, response); err != nil {
		return err
	}

	if p.verbose {
		fmt.Println("--------------")
		fmt.Println()
		fmt.Println("Response: ")
		fmt.Println(response.Header.String())
		fmt.Println("")
		if len(response.Body()) > 0 {
			fmt.Println("Body: ")
			fmt.Println(string(response.Body()))
		} else {
			fmt.Println("Body: EMPTY")
		}
	}
	return nil
}
