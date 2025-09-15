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
		"items": []map[string]interface{}{
			{
				"name":   "Product 1",
				"value":  1000,
				"amount": 2,
			},
		},
	}

	res, err := efi.CreateSubscription(1, body) // no lugar do 1 coloque o plan_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
