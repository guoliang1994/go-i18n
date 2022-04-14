package driver

import (
	"fmt"
	"gopkg.in/guoliang1994/go-i18n.v3/contract"
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

func (Self *JsonFileI18NImpl) LoadLang() []byte {
	fileName := fmt.Sprintf(Self.langDir + "lang.json")
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(contract.PkgName, ": json file load lang  err,", err)
	}
	return data
}
