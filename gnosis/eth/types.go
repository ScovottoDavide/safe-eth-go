package eth

import (
	"fmt"
	"net/url"
)

type URI struct {
	address string
}

func NewURI(uri string) *URI {
	new_uri := &URI{address: uri}
	if new_uri.isValid() {
		return new_uri
	}
	return nil
}

func (uri *URI) isValid() bool {
	_, err := url.ParseRequestURI(uri.address)
	if err != nil {
		fmt.Printf("Provided Invalid URI for the EthereumClient: %s", err)
		return false
	}
	return true
}

func (uri *URI) GetAddress() string {
	return uri.address
}
