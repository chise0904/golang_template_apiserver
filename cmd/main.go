package main

import (
	"github.com/spf13/cobra"
	apigateway "github.com/chise0904/golang_template_apiserver/cmd/gateway"
)

func main() {
	rootCmd := cobra.Command{}

	rootCmd.AddCommand(apigateway.GatewayCmd())

	rootCmd.Execute()
}
