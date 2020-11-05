package cli

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	CliClient struct {
		Config *Config

		Client *http.Client
	}
)

func NewCliClient() (cliClient *CliClient, err error) {

	cliClient = &CliClient{
		Client: &http.Client{},
	}

	if cliClient.Config, err = NewConfig(); err != nil {
		return
	}

	return
}

func (cliClient *CliClient) Run(macAddr string) (output string, err error) {

	var (
		resp *http.Response
		body []byte
	)

	if resp, err = cliClient.Client.Get("http://localhost:8090/command?mac=" + macAddr); err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return
}
