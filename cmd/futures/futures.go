/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package futures

import (
	"fmt"

	"github.com/spf13/cobra"
)

// FuturesCmd represents the futures command
var FuturesCmd = &cobra.Command{
	Use:   "futures",
	Short: "futures - Retrieve and store data for Binance Futures",
	Long: `The futures command is used to retrieve and store data for Binance Futures.
	It provides several subcommands for working with different types of data, such as klines, long/short ratio, open interest, and more.
	The tool uses a PostgreSQL database to store the data and provides a simple and easy-to-use interface for working with the data.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("futures called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// futuresCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// futuresCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
