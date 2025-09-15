package main

import (
	"fmt"

	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const inicio = "2025-01-01T03:01:35Z"
	const fim = "2025-12-31T22:01:35Z"

	res, err := efi.PixListLocationRecurrenceAutomatic(inicio, fim)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
