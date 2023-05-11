package handlers

type Uploader interface {
	Upload(f string, bytes []byte) (string, error)
}
