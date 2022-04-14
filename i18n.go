package i18n

import (
	"fmt"
	"github.com/dlclark/regexp2"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/guoliang1994/go-i18n.v3/contract"
	_ "io/ioutil"
	"strings"
)

type I18N struct {
	location    Location
	lang        jsoniter.RawMessage
	driverChain []contract.I18NDriver // 加载的所有语言
}

func NewI18N(location Location) *I18N {
	i18n := I18N{
		location: location,
		lang:     jsoniter.RawMessage{},
	}
	return &i18n
}

//T get lang
func (Self *I18N) T(keyPath string, placeholder ...string) string {
	Self.dealLang()
	pathArr := strings.Split(keyPath, ".")
	pathArr = append(pathArr, string(Self.location))
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

func (Self *I18N) AddLang(driver contract.I18NDriver) *I18N {
	Self.driverChain = append(Self.driverChain, driver)
	return Self
}

func (Self *I18N) dealLang() {
	for _, d := range Self.driverChain {
		var lang1 map[string]interface{}
		var lang2 map[string]interface{}

		if len(Self.lang) > 0 {
			err := jsoniter.Unmarshal(Self.lang, &lang2)
			if err != nil {
				panic("deal Lang error" + err.Error())
			}
		}

		langBytes := d.LoadLang()
		err := jsoniter.Unmarshal(langBytes, &lang1)
		if err != nil {
			panic("deal Lang error" + err.Error())
		}
		mergeLang := JsonMerge(lang2, lang1)
		Self.lang, _ = jsoniter.Marshal(mergeLang)
	}
}

//ChangeLocation the function can change location when the program is running
func (Self *I18N) ChangeLocation(location Location) *I18N {
	Self.location = location
	Self.dealLang()
	return Self
}
