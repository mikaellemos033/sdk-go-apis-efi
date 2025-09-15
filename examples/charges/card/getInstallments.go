package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay"
)

func main() {

	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	res, err := efi.GetInstallments(20000, "visa") // total e bandeira

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
