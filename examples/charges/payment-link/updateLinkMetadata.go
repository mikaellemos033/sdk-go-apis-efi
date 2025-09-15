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
		"custom_id":        "Product 0001",
		"notification_url": "http://domain.com/notification",
	}

	res, err := efi.UpdateChargeMetadata(1, body) // no lugar do 1 coloque o id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
