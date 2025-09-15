package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	body := map[string]interface{}{}
	const id = "423"

	res, err := efi.PixUnlinkTxidLocation(id, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
