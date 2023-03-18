/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package symbols

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/seorlando33/binance-data-retriever/cmd/app"
)

var symbols []string
var all bool

// SymbolsCmd represents the symbols command
var SymbolsCmd = &cobra.Command{
	Use:   "symbols",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(symbols) > 0 && all {

			fmt.Println("You can not use flag symbols and all simultaneously")

		} else if all {

			fmt.Println("Retrieving and storing all the symbols")

			err := app.Service.Futures.Symbol.InsertSymbol()
			if err != nil {
				fmt.Println(err.Error())
			}

		} else {
			fmt.Printf("Retrieving and storing the symbols: %v.", symbols)
			err := app.Service.Futures.Symbol.InsertSymbol(symbols...)
			if err != nil {
				fmt.Println(err.Error())
			}

		}
	},
}

func init() {

	SymbolsCmd.Flags().StringSliceVarP(&symbols, "symbols", "s", nil, "Symbols to retrieve")

	SymbolsCmd.Flags().BoolVarP(&all, "all", "a", false, "Retrieve all the Symbols data")
}