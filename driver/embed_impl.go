package driver

import (
	"embed"
	"fmt"
	"github.com/guoliang1994/go-i18n/contract"
)

type EmbedI18NImpl struct {
	f       embed.FS
	dirName string
}

func NewEmbedI18NImpl(f embed.FS, dirName string) contract.I18NDriver {
	i18n := EmbedI18NImpl{
		f:       f,
		dirName: dirName,
	}
	return &i18n
}

func (Self *EmbedI18NImpl) LoadLang(location string) []byte {
	data, err := Self.f.ReadFile(Self.dirName + location + ".json")
	if err != nil {
		fmt.Println("go:embed load lang err,", err)
	}
	return data
}
