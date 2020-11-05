package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DefaultMacApi = "https://api.macaddress.io/v1?apiKey=%s&output=json&search=%s"
)

type (
	MacAddressApi struct {
		apiKey string
	}

	MacAddressResp struct {
		Vendor struct {
			CompanyName string `json:"companyName"`
		} `json:"vendorDetails"`
	}
)

func NewMacAddressApi(apiKey string) (macApi *MacAddressApi, err error) {

	macApi = &MacAddressApi{
		apiKey: apiKey,
	}

	return
}

func (macApi *MacAddressApi) FetchInfo(macAddress string) (macResp *MacAddressResp, err error) {

	var (
		resp *http.Response
	)

	macResp = &MacAddressResp{}

	if resp, err = http.Get(fmt.Sprintf(DefaultMacApi, macApi.apiKey, macAddress)); err != nil {
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(macResp); err != nil {
		return
	}

	return
}
