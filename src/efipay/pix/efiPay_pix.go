package pix

type pix struct {
	endpoints
}

func NewEfiPay(configs map[string]interface{}) Pix {
	clientID := configs["client_id"].(string)
	clientSecret := configs["client_secret"].(string)
	CA := configs["CA"].(string)
	Key := configs["Key"].(string)
	sandbox := configs["sandbox"].(bool)
	timeout := configs["timeout"].(int)

	requestClient := newRequester(clientID, clientSecret, CA, Key, sandbox, timeout)
	efi := pix{}
	efi.requester = requestClient
	return &efi
}
