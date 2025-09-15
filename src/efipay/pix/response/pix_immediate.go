package response

import "github.com/mikaellemos033/sdk-go-apis-efi/src/efipay/pix/model"

type PixImmediate struct {
	ID        string         `json:"txid,omitempty"`
	Status    string         `json:"status,omitempty"`
	Calendly  model.Calendly `json:"calendario,omitempty"`
	PixKey    string         `json:"chave,omitempty"`
	Requester string         `json:"solicitacaoPagador,omitempty"`
	PixURL    string         `json:"location,omitempty"`
	PixCode   string         `json:"pixCopiaECola,omitempty"`
	Location  Location       `json:"loc,omitempty"`
	Infos     []model.Info   `json:"infoAdicionais"`
}
