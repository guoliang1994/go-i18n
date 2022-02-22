package test

import (
	"github.com/guoliang1994/go-i18n"
	"github.com/guoliang1994/go-i18n/driver"
	"testing"
)

func TestJsonFileImplGetLang(t *testing.T) {
	driver := driver.NewJsonFileI18nImpl("..\\lang")
	i18n := go_i18n.NewI18N(go_i18n.Chinese, driver)
	msg := i18n.T("install.error", i18n.T("appName"), i18n.T("needRoot"))
	assert := "国际化支持安装失败, 失败原因:需要root用户"
	if msg != assert {
		t.Fatal("want:", assert, "get", msg)
	}
}
