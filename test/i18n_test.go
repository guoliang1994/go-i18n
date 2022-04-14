package test

import (
	"gopkg.in/guoliang1994/go-i18n.v3"
	"gopkg.in/guoliang1994/go-i18n.v3/driver"
	"testing"
)

//func TestJsonFileImplGetLang(t *testing.T) {
//	drv := driver.NewJsonFileI18nImpl("lang")
//	lang := i18n.NewI18N(i18n.Chinese, drv)
//	msg := lang.T("install.error", lang.T("appName"), lang.T("needRoot"))
//	assert := "国际化支持安装失败, 失败原因:需要root用户"
//	if msg != assert {
//		t.Fatal("want:", assert, "get", msg)
//	}
//}
//
//func TestAddLang(t *testing.T) {
//	driver1 := driver.NewJsonFileI18nImpl("lang")
//	driver2 := driver.NewJsonFileI18nImpl("lang2")
//	lang := i18n.NewI18N(i18n.Chinese, driver1)
//	lang.AddLang(driver2)
//	msg := lang.T("hello")
//	assert := "world"
//	if msg != assert {
//		t.Fatal("want:", assert, "get", msg)
//	}
//}
//
//func TestAddLangAndLevelLang(t *testing.T) {
//	driver1 := driver.NewJsonFileI18nImpl("lang")
//	driver2 := driver.NewJsonFileI18nImpl("lang2")
//	lang := i18n.NewI18N(i18n.Chinese, driver1)
//	lang.AddLang(driver2)
//	msg := lang.T("nice.too.meet")
//	assert := "you"
//	if msg != assert {
//		t.Fatal("want:", assert, "get", msg)
//	}
//}
//
//func TestStringDriver(t *testing.T) {
//	langMap := make(map[string]string)
//	langMap[i18n.Chinese] = `{"nice":"你好字节"}`
//	langMap[i18n.English] = `{"nice":"nice byte"}`
//	driver1 := driver.NewStringI18NImpl(langMap)
//	lang := i18n.NewI18N(i18n.English, driver1)
//	msg := lang.T("nice")
//	assert := "nice byte"
//	if msg != assert {
//		t.Fatal("want:", assert, "get", msg)
//	}
//}
//
//func TestBytesDriver(t *testing.T) {
//	langMap := make(map[string][]byte)
//	langMap[i18n.Chinese] = []byte(`{"nice":"你好字节"}`)
//	langMap[i18n.English] = []byte(`{"nice":"nice byte"}`)
//	driver1 := driver.NewBytesI18NImpl(langMap)
//	lang := i18n.NewI18N(i18n.English, driver1)
//	msg := lang.T("nice")
//	assert := "nice byte"
//	if msg != assert {
//		t.Fatal("want:", assert, "get", msg)
//	}
//}
//
//func TestGoBindataDriver(t *testing.T) {
//	driver1 := driver.NewGoBindataI18NImpl(Asset, "lang/")
//	lang := i18n.NewI18N(i18n.English, driver1)
//	msg := lang.T("install.success")
//	assert := "{{program}} install success"
//	if msg != assert {
//		t.Fatal("want:", assert, "get", msg)
//	}
//}
//
////go:embed lang
//var f embed.FS
//
//func TestEmbedDriver(t *testing.T) {
//	driver1 := driver.NewEmbedI18NImpl(f, "lang/")
//	lang := i18n.NewI18N(i18n.English, driver1)
//	msg := lang.T("install.success")
//	assert := "{{program}} install success"
//	if msg != assert {
//		t.Fatal("want:", assert, "get", msg)
//	}
//}
//
//
//func TestDynamicAddLang(t *testing.T) {
//	lang := i18n.NewI18N(i18n.English, nil)
//	driver1 := driver.NewJsonFileI18nImpl("lang")
//	msg := lang.AddLang(driver1).T("install.success")
//	msg = lang.ChangeLocation(i18n.Chinese).T("install.success")
//	assert := "world"
//	if msg != assert {
//		t.Fatal("want:", assert, "get", msg)
//	}
//}

func TestDynamicAddLang(t *testing.T) {
	lang := i18n.NewI18N(i18n.Zh_CN)
	driver1 := driver.NewJsonFileI18nImpl("lang/")
	driver2 := driver.NewJsonFileI18nImpl("lang2/")
	//driver2 := driver.NewGoBindataI18NImpl(Asset, "lang/")
	msg := lang.AddLang(driver1).AddLang(driver2).T("install.success", "np")
	assert := "np安装成功"
	if msg != assert {
		t.Fatal("want:", assert, "get", msg)
	}
	msg = lang.ChangeLocation(i18n.En_WW).T("install.success", "np")
	assert = "np install success"
	if msg != assert {
		t.Fatal("want:", assert, "get", msg)
	}
	msg = lang.ChangeLocation(i18n.En_WW).T("extend.success")
	assert = "扩展语言成功"
	if msg != assert {
		t.Fatal("want:", assert, "get", msg)
	}
}
