package main

import (
	"fmt"

	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const idSolicRec = "SC09089356202509035192e32902e"

	body := map[string]interface{}{
		"status": "CANCELADA",
	}

	res, err := efi.PixUpdateRequestRecurrenceAutomatic(idSolicRec, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
