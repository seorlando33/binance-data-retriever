package cmd

import (
        "os"

        "github.com/seorlando33/binance-data-retriever/cmd/app"
        "github.com/seorlando33/binance-data-retriever/cmd/futures"
        "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
        Use:   "bdr",
        Short: "bdr - A CLI tool for retrieving and storing Binance data",
        Long: `bdr is a command-line tool written in Go for retrieving and storing data from the Binance exchange. 
        It supports various types of data, such as klines, long/short ratio, open interest, and more.
        The tool uses a PostgreSQL database to store the data and provides a simple and easy-to-use interface for working with the data.
        `,
}

func addSubCommands() {
        rootCmd.AddCommand(futures.FuturesCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
        err := rootCmd.Execute()
        if err != nil {
                os.Exit(1)
        }
}

func init() {

        rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

        app.Run()

        addSubCommands()
}