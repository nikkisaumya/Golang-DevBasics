package main

import (
	"example.com/macaddressfetcher/cli"
  "fmt"
	"log"
	"os"
	"time"
  "flag"
)

func main() {

	var (
		cliServer *cli.CliServer
		cliClient *cli.CliClient

		err error
	)

	if cliServer, err = cli.NewCliServer(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	go func() {
		if err = cliServer.Run(); err != nil {
			log.Println(err)
			os.Exit(-1)
		}
	}()

	time.Sleep(2 * time.Second)

	if cliClient, err = cli.NewCliClient(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

  macAddr:= flag.String("mac","","mac")
  flag.Parse()

  if *macAddr == "" {
    fmt.Println("mac is mandatory param -mac=mac Addrress goes here")
    os.Exit(1)
  }

	cliClient.Run(*macAddr)
}
