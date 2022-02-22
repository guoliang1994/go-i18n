package driver

import (
	"fmt"
	"github.com/guoliang1994/go-i18n/contract"
	"io/ioutil"
)

type JsonFileI18NImpl struct {
	langDir string
}

func NewJsonFileI18nImpl(langDir string) contract.I18NDriver {
	f := JsonFileI18NImpl{
		langDir: langDir,
	}
	return &f
}

func (Self *JsonFileI18NImpl) LoadLang(location string) []byte {
	fileName := fmt.Sprintf(Self.langDir+"\\%s.json", location)
	data, _ := ioutil.ReadFile(fileName)
	return data
}
