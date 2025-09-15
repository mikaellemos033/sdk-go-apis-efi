package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/payments"
)

func main() {

	credentials := configs.Credentials
	efi := payments.NewEfiPay(credentials)

	params := map[string]string{
		"codBarras": "36400000000000000000000000000000000000000000000",
	}

	res, err := efi.PayDetailBarCode(params)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
