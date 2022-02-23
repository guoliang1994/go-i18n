package driver

import (
	"fmt"
	"github.com/guoliang1994/go-i18n/contract"
)

type GoBindataI18NImpl struct {
	f       func(name string) ([]byte, error)
	dirName string
}

func NewGoBindataI18NImpl(f func(name string) ([]byte, error), dirName string) contract.I18NDriver {
	i18n := GoBindataI18NImpl{
		f:       f,
		dirName: dirName,
	}
	return &i18n
}

func (Self *GoBindataI18NImpl) LoadLang(location string) []byte {
	data, err := Self.f(Self.dirName + location + ".json")
	if err != nil {
		fmt.Println(contract.PkgName, ": go-bindata load lang err,", err)
	}
	return data
}
