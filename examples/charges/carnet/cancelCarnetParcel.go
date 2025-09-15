package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay"
)

func main() {

	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	res, err := efi.CancelCarnetParcel(1, 1) // no lugar dos 1s coloque o carnet_id e o numero da parcela respectivamente

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
