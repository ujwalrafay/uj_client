/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uj the boss is back ")
		fmt.Println("args in run is", args)
		list(args)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func list(args1 []string) {
	fmt.Println("args in list are", args1)
	if len(args1) < 3 {
		fmt.Println("invalid number of args")
		fmt.Println("Usage--- uj_client api {hostname} {port} {query} ")
		return
	}
	method := args1[0]
	host := args1[1]
	port := args1[2]
	query := args1[3]

	url := "http://" + host + ":" + port + "/" + query
	fmt.Println(url)
	if strings.EqualFold(method, "get") {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("error while accesing using get", err)
		}
		resp_data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error while reading", err)
		}

		fmt.Println(string(resp_data))
	}
}
