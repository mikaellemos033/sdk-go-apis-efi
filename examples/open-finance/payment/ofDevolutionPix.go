package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/open_finance"
)

func main() {

	credentials := configs.Credentials
	efi := open_finance.NewEfiPay(credentials)

	body := map[string]interface{}{
		"identificadorPagamento": "urn:gerencianet:ea807997-9c28-47a7-8ebc-eeb672ea59f0",
		"valorDevolucao":         "0.01",
	}

	res, err := efi.OfDevolutionPix(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
