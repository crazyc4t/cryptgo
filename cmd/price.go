/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	gecko "github.com/superoo7/go-gecko/v3"
	"log"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Gives you the current price of your desired crypto in fiat!",
	Long: `The price command lets you know the price of a cryptocurrency in fiat, for example:

cryptgo price monero cad

Being the cryptocurrency monero and the desired fiat currency the canadian dollar.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			cryptoPrice(args[0], args[1])

		} else if len(args) > 2 {
			fmt.Println("Extra arguments, run cryptgo price --help  for more.")
		} else if len(args) == 0 {
			cmd.Help()
		} else {
			fmt.Println("Missing arguments, run cryptgo price --help  for more.")
		}
	},
}

func cryptoPrice(crypto, fiat string) {
	cg := gecko.NewClient(nil)
	price, err := cg.SimpleSinglePrice(crypto, fiat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%[1]s is worth %[2]g %[3]v", price.ID, price.MarketPrice, price.Currency))
}

func init() {
	rootCmd.AddCommand(priceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// priceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	priceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
