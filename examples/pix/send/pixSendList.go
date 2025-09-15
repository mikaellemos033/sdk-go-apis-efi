package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const e2eid = "E00416968202105211756Rh0iPsaJ1RK"

	res, err := efi.PixSendList(e2eid)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
