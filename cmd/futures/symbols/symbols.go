/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package symbols

import (
	"fmt"
	"os"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/seorlando33/binance-data-retriever/cmd/app"
	"github.com/spf13/cobra"
)

var symbols []string
var all bool

// SymbolsCmd represents the symbols command
var SymbolsCmd = &cobra.Command{
	Use:   "symbols",
	Short: "Get the symbols",
	Long:  `Retrieve the symbols data from Binance Futures API`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(symbols) > 0 && all {

			fmt.Println("You can not use flag symbols and all simultaneously")

		} else if all {

			symbols, err := app.Service.Futures.Symbol.GetSymbols()
			if err != nil {
				fmt.Println(err.Error())
			}
			// send symbols to stdout
			err = sendToStdout(symbols)
			if err != nil {
				fmt.Println(err.Error())
			}

		} else {

			symbols, err := app.Service.Futures.Symbol.GetSymbols(symbols...)
			if err != nil {
				fmt.Println(err.Error())
			}
			// send symbols to stdout
			err = sendToStdout(symbols)
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

// stdout sender
func sendToStdout(symbols []futures.Symbol) error {
	// convert symbols to bytes
	body := []byte(fmt.Sprintf("%v", symbols))

	// send symbols to stdout
	_, err := os.Stdout.Write(body)

	return err
}
