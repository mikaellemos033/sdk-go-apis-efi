package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay"
)

func main() {

	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	customer := map[string]interface{}{
		"name":         "Gorbadoc Oldbuck",
		"cpf":          "73504596295",
		"phone_number": "5144916523",
	}

	body := map[string]interface{}{
		"payment": map[string]interface{}{
			"banking_billet": map[string]interface{}{
				"expire_at": "2022-12-12",
				"customer":  customer,
			},
		},
	}

	res, err := efi.DefineSubscriptionPayMethod(1, body) // no lugar do 1 coloque o subscription_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
