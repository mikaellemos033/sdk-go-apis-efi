package request

import "github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix/model"

type PixImmediate struct {
	Amount    model.Amount   `json:"valor,omitempty"`
	Calendly  model.Calendly `json:"calendario,omitempty"`
	PixKey    string         `json:"chave,omitempty"`
	Infos     []model.Info   `json:"infosAdicionais,omitempty"`
	Requester string         `json:"solicitacaoPagador,omitempty"`
}
