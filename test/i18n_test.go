package test

import (
	"github.com/guoliang1994/go-i18n"
	"github.com/guoliang1994/go-i18n/driver"
	"testing"
)

func TestJsonFileImplGetLang(t *testing.T) {
	drv := driver.NewJsonFileI18nImpl("lang")
	lang := i18n.NewI18N(i18n.Chinese, drv)
	msg := lang.T("install.error", lang.T("appName"), lang.T("needRoot"))
	assert := "国际化支持安装失败, 失败原因:需要root用户"
	if msg != assert {
		t.Fatal("want:", assert, "get", msg)
	}
}

func TestAddLang(t *testing.T) {
	driver1 := driver.NewJsonFileI18nImpl("lang")
	driver2 := driver.NewJsonFileI18nImpl("lang2")
	lang := i18n.NewI18N(i18n.Chinese, driver1)
	lang.AddLang(driver2)
	msg := lang.T("hello")
	assert := "world"
	if msg != assert {
		t.Fatal("want:", assert, "get", msg)
	}
}

func TestAddLangAndLevelLang(t *testing.T) {
	driver1 := driver.NewJsonFileI18nImpl("lang")
	driver2 := driver.NewJsonFileI18nImpl("lang2")
	lang := i18n.NewI18N(i18n.Chinese, driver1)
	lang.AddLang(driver2)
	msg := lang.T("nice.too.meet")
	assert := "you"
	if msg != assert {
		t.Fatal("want:", assert, "get", msg)
	}
}
