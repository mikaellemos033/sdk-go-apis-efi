package main

import (
	"fmt"

	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const idRec = "RN0908935620250904cafd5a29364"

	res, err := efi.PixDetailRecurrenceAutomatic(idRec)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
