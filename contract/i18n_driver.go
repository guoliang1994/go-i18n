package contract

type I18NDriver interface {
	LoadLang(location string) []byte
}
