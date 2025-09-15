package main

import (
	"fmt"

	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	body := map[string]interface{}{
		"idRec": "RN09089356202509040648092afd6",

		"calendario": map[string]interface{}{
			"dataExpiracaoSolicitacao": "2025-09-05T12:17:11.926Z",
		},

		"destinatario": map[string]interface{}{
			"agencia":          "0001",
			"conta":            "1234567",
			"cpf":              "12345678910",
			"ispbParticipante": "00416968",
		},
	}

	res, err := efi.PixCreateRequestRecurrenceAutomatic(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
