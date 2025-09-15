package main

import (
	"fmt"

	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const id = "1"

	res, err := efi.PixUnlinkLocationRecurrenceAutomatic(id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
