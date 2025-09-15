package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/open_finance"
)

func main() {

	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	res, err := efi.OfConfigDetail()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
