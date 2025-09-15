package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay"
)

func main() {

	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	body := map[string]interface{}{
		"name":     "My plan",
		"interval": 2,
		"repeats":  nil,
	}

	res, err := efi.CreatePlan(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
