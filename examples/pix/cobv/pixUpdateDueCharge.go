package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const txid = "adssshdsjdsjeyccdyddsasdstxid01"

	body := map[string]interface{}{
		"valor": map[string]interface{}{
			"original": "1.00",
		},
		"solicitacaoPagador": "Enter the order number or identifier.",
	}

	res, err := efi.PixUpdateDueCharge(txid, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
