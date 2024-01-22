package shortener

type RedirectSerializer interface {
	Decode(data []byte) (*Redirect, error)
	Encode(data *Redirect) ([]byte, error)
}