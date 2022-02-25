package driver

import (
	"gopkg.in/guoliang1994/go-i18n.v2/contract"
)

type BytesI18NImpl struct {
	langString map[string][]byte
}

func NewBytesI18NImpl(langString map[string][]byte) contract.I18NDriver {
	i18n := BytesI18NImpl{
		langString: langString,
	}
	return &i18n
}

func (Self *BytesI18NImpl) LoadLang(location string) []byte {
	return Self.langString[location]
}
