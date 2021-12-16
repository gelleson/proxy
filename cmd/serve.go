package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/spf13/cobra"
	"log"
	"proxy/internal/proxy/http"
	"time"
)

var (
	host          string
	listen        string
	setFromConfig bool
	proxyTimeout  time.Duration
	local         bool
)

var serveHTTPCMD = &cobra.Command{
	Use:   "serve",
	Short: "start http proxy",
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New(fiber.Config{
			DisableStartupMessage: true,
		})

		config := http.Config{
			Host:              host,
			SetHostFromConfig: setFromConfig,
		}

		if local {
			config.Schema = "http"
		} else {
			config.Schema = "https"
		}

		proxy := http.New(verbose, config)

		app.All("*", timeout.New(func(ctx *fiber.Ctx) error {
			request := ctx.Request()
			response := ctx.Response()
			return proxy.Proxy(request, response)
		}, proxyTimeout))

		log.Fatalln(app.Listen(listen))
	},
}

func init() {
	root.AddCommand(serveHTTPCMD)

	serveHTTPCMD.PersistentFlags().StringVar(&host, "set-host", "", "localhost")
	serveHTTPCMD.PersistentFlags().BoolVarP(&local, "local", "d", false, "")
	serveHTTPCMD.PersistentFlags().StringVarP(&listen, "listen", "l", ":33413", "localhost")
	serveHTTPCMD.PersistentFlags().BoolVar(&setFromConfig, "set-from-config", true, "localhost")
	serveHTTPCMD.PersistentFlags().DurationVar(&proxyTimeout, "proxy-timeout", time.Second*10, "10m")
}
