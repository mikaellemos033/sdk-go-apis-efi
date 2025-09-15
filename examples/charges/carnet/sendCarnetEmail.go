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
		"email": "oldbuck@gerencianet.com.br",
	}

	res, err := efi.SendCarnetEmail(1, body) // no lugar do 1 coloque o carnet_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
