package driver

import (
	"fmt"
	"gopkg.in/guoliang1994/go-i18n.v1/contract"
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
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(contract.PkgName, ": json file load lang  err,", err)
	}
	return data
}
