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
	location    string
	lang        map[string]jsoniter.RawMessage
	driverChain []contract.I18NDriver // 加载的所有语言
}

func NewI18N(location string) *I18N {
	i18n := I18N{
		location: location,
		lang:     make(map[string]jsoniter.RawMessage),
	}
	return &i18n
}

//T get lang
func (Self *I18N) T(keyPath string, placeholder ...string) string {
	//If the language is not loaded
	_, ok := Self.lang[Self.location]
	if !ok {
		Self.dealLang()
	}

	pathArr := strings.Split(keyPath, ".")
	var getter jsoniter.Any
	for _, path := range pathArr {
		if getter != nil {
			getter = jsoniter.Get([]byte(getter.ToString()), path)
		} else {
			getter = jsoniter.Get(Self.lang[Self.location], path)
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

func (Self *I18N) AddLang(driver contract.I18NDriver) *I18N {
	Self.driverChain = append(Self.driverChain, driver)
	return Self
}

func (Self *I18N) dealLang() {
	for _, d := range Self.driverChain {
		var lang1 map[string]interface{}
		var lang2 map[string]interface{}
		langBytes := d.LoadLang(Self.location)
		_ = jsoniter.Unmarshal(langBytes, &lang1)
		_ = jsoniter.Unmarshal(Self.lang[Self.location], &lang2)
		mergeLang := JsonMerge(lang1, lang2)
		Self.lang[Self.location], _ = jsoniter.Marshal(mergeLang)
	}
}

//ChangeLocation the function can change location when the program is running
func (Self *I18N) ChangeLocation(location string) *I18N {
	Self.location = location
	Self.dealLang()
	return Self
}
