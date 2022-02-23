package driver

import (
	"github.com/guoliang1994/go-i18n/contract"
)

type StringI18NImpl struct {
	langString map[string]string
}

func NewStringI18NImpl(langString map[string]string) contract.I18NDriver {
	i18n := StringI18NImpl{
		langString: langString,
	}
	return &i18n
}

func (Self *StringI18NImpl) LoadLang(location string) []byte {
	return []byte(Self.langString[location])
}
