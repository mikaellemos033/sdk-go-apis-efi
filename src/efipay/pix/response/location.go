package response

type Location struct {
	ID      string `json:"id,omitempty"`
	PixURL  string `json:"location,omitempty"`
	Type    string `json:"tipoCob,omitempty"`
	Created string `json:"criacao,omitempty"`
}
