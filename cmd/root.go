package cmd

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use:   "proxyd",
	Short: "Go based HTTP Proxy",
}

var (
	verbose = false
)

func init() {
	root.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "true | false")
}

func Execute() error {
	return root.Execute()
}
