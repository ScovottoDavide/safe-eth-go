package eth

import (
	"fmt"
	"net/url"
)

type URI struct {
	address string
}

func (uri *URI) isValid() bool {
	_, err := url.ParseRequestURI(uri.address)
	if err != nil {
		fmt.Printf("Provided Invalid URI for the EthereumClient: %s", err)
		return false
	}
	return true
}
