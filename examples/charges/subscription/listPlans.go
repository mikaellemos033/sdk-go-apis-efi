package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay"
)

func main() {

	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	res, err := efi.ListPlans(20, 0) // limit e offset

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
