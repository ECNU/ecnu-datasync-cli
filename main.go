package main

import (
	"ecnu-datasync-cli/g"
	"flag"
	"fmt"
)

func main() {

	version := flag.Bool("v", false, "show version")
	clientId := flag.String("c", "", "oauth2 client_id")
	clientSecret := flag.String("s", "", "oauth2 client_id")
	apiPath := flag.String("a", "", "api path")
	output := flag.String("o", "", "output file")
	cfg := flag.String("config", "", "configuration file")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		return
	}

	if *cfg != "" {
		g.ParseConfig(*cfg)
		if err := g.SyncWithConfig(); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	if *apiPath != "" && *clientSecret != "" && *clientId != "" && *output != "" {
		if err := g.SyncWithoutConfig(clientId, clientSecret, output, apiPath); err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	fmt.Println("参数填写不完整，请填写完整参数")
	fmt.Println("Usage :")
	flag.PrintDefaults()
}
