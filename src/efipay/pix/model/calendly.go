package model

type Calendly struct {
	CreatedAt string `json:"criacao,omitempty"`
	ExpireIn  int    `json:"expiracao,omitempty"`
}
