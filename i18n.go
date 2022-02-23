package i18n

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/guoliang1994/go-i18n/contract"
	jsoniter "github.com/json-iterator/go"
	_ "io/ioutil"
	"strings"
)

type I18N struct {
	location string
	lang     jsoniter.RawMessage
}

func NewI18N(location string, driver contract.I18NDriver) *I18N {
	i18n := I18N{
		location: location,
	}
	_ = jsoniter.Unmarshal(driver.LoadLang(location), &i18n.lang)
	return &i18n
}

//T get lang
func (Self *I18N) T(keyPath string, placeholder ...string) string {
	pathArr := strings.Split(keyPath, ".")
	var getter jsoniter.Any
	for _, path := range pathArr {
		if getter != nil {
			getter = jsoniter.Get([]byte(getter.ToString()), path)
		} else {
			getter = jsoniter.Get(Self.lang, path)
		}
	}
	if getter == nil {
		return ""
	}
	msg := getter.ToString()
	reg := regexp2.MustCompile(`(\{\{[\d\w\s]*\}\})`, 0)
	group, err := reg.FindStringMatch(msg)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// replace placeholder
	// example: {{program}} install success will be output nps install success
	for _, p := range placeholder {
		if group != nil {
			msg = strings.Replace(msg, group.String(), p, -1)
			group, _ = reg.FindNextMatch(group)
		}
	}

	return msg
}

func (Self *I18N) AddLang(driver contract.I18NDriver) {
	content := driver.LoadLang(Self.location)
	var lang1 map[string]interface{}
	var lang2 map[string]interface{}
	_ = jsoniter.Unmarshal(content, &lang1)
	_ = jsoniter.Unmarshal(Self.lang, &lang2)
	mergeLang := JsonMerge(lang1, lang2)
	Self.lang, _ = jsoniter.Marshal(mergeLang)
}
