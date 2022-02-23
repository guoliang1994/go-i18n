package contract

const PkgName = "github/guoliang1994/go-i18n"

type I18NDriver interface {
	LoadLang(location string) []byte
}
