package driver

import (
	"gopkg.in/guoliang1994/go-i18n.v1/contract"
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
