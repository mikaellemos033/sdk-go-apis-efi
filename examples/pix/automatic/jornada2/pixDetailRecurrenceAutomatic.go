package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/examples/configs"
	"github.com/efipay/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const idRec = "RN0908935620250902c5bf68c68d4"

	res, err := efi.PixDetailRecurrenceAutomatic(idRec)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
