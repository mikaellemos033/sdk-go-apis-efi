package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const chave = "" //sua chave Pix Efí

	body := map[string]interface{}{

		"webhookUrl": "https://seu_webhook",
	}

	res, err := efi.PixConfigWebhook(chave, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
