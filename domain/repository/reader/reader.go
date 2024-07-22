package reader

type Reader interface {
	ReadFile(path string, target interface{}) error
}
