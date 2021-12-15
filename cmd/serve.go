package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
	"proxy/internal/proxy/http"
)

var (
	host          string
	setFromConfig bool
)

var serveHTTPCMD = &cobra.Command{
	Use:   "serve",
	Short: "start http proxy",
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New(fiber.Config{
			DisableStartupMessage: true,
		})

		proxy := http.New(verbose, http.Config{
			Host:              host,
			SetHostFromConfig: setFromConfig,
		})

		app.All("*", func(ctx *fiber.Ctx) error {
			request := ctx.Request()
			response := ctx.Response()
			return proxy.Proxy(request, response)
		})

		log.Fatalln(app.Listen(":33413"))
	},
}

func init() {
	root.AddCommand(serveHTTPCMD)

	serveHTTPCMD.PersistentFlags().StringVar(&host, "set-host", "", "localhost")
	serveHTTPCMD.PersistentFlags().BoolVar(&setFromConfig, "set-from-config", true, "localhost")
}
