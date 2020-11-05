package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/macaddressfetcher/api"
)

const (
	DefaultConfigFile = "../conf/config.json"
)

type (
	CliServer struct {
		Config *Config

		MacApi *api.MacAddressApi
	}

	Config struct {
		ApiKey string `json:"api_key"`
		Server struct {
			Port string `json:"port"`
		} `json:"server"`
	}
)

func NewConfig() (config *Config, err error) {

	var (
		configB []byte
	)

	config = &Config{}

  configB, err = ioutil.ReadFile(DefaultConfigFile)
	if err != nil {
		return
	}

   err = json.Unmarshal(configB, config)
	if err != nil {
		return
	}

	return
}

func NewCliServer() (cliServer *CliServer, err error) {

	cliServer = &CliServer{}

	if cliServer.Config, err = NewConfig(); err != nil {
		return
	}

	if cliServer.MacApi, err = api.NewMacAddressApi(cliServer.Config.ApiKey); err != nil {
		return
	}

	return
}

func (cliServer *CliServer) RunCommand(w http.ResponseWriter, r *http.Request) {

	var (
		macAddr string
		err     error
		macResp *api.MacAddressResp
	)

	macAddr = r.URL.Query().Get("mac")

	if macResp, err = cliServer.MacApi.FetchInfo(macAddr); err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintf(w, macResp.Vendor.CompanyName)

	return

}

func (cliServer *CliServer) Run() (err error) {

	log.Println("Starting CLI Server")

	http.HandleFunc("/command", cliServer.RunCommand)
	err = http.ListenAndServe(":"+cliServer.Config.Server.Port, nil)

	return
}
