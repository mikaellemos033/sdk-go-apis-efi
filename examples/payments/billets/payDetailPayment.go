package main

import (
	"fmt"
	"github.com/mikaellemos033/sdk-go-apis-efi/examples/configs"
	"github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/payments"
)

func main() {

	credentials := configs.Credentials
	efi := payments.NewEfiPay(credentials)

	params := map[string]string{
		"idPagamento": "1",
		// "status" : "REALIZADO", // EM_PROCESSAMENTO, AGENDADO, LIQUIDADO, CANCELADO, NAO_REALIZADO
	}

	res, err := efi.PayDetailPayment(params)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
