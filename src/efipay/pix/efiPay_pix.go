package pix

type pix struct {
	endpoints
}

func NewEfiPay(configs *Setup) Pix {
	clientID := configs.ClientId
	clientSecret := configs.ClientSecret
	CA := configs.CA
	Key := configs.Key
	sandbox := configs.Sandbox
	timeout := configs.Timeout

	requestClient := newRequester(clientID, clientSecret, CA, Key, sandbox, timeout)
	efi := pix{}
	efi.requester = requestClient
	return &efi
}
