package pix

type Setup struct {
	ClientId     string
	ClientSecret string
	CA           []byte
	Key          []byte
	Sandbox      bool
	Timeout      int
}
