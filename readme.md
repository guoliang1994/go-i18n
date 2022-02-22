### go-i18n
用最简单的方法让你的go程序支持国际化，我们使用接口来支持多种驱动的扩展。

#### 目前支持的驱动
+ [x]  json 格式的语言文件
+ [x]  Embed 嵌入式文件系统 (使用 json格式)
+ [ ] toml 格式
+ [ ] yml 格式
+ [ ] xml 格式
+ [ ] sqlite
+ [ ] postgresql
+ [ ] redis
+ 其他你想支持的任何数据源

### 语言支持
+ [x] 中文
+ [x] English
+ [ ] other

### 如何使用
> import "github.com/guoliang1994/go-i18n"

语言文件示例
```json
{
  "install": {
    "success": "{{program}}安装成功",
    "error": "{{program}}安装失败, 失败原因:{{reason}}"
  },
  "uninstall": {
    "success": "{{program}}卸载成功",
    "error": "{{program}}卸载失败"
  },
  "start": {
    "success": "{{program}}启动成功",
    "error": "{{program}}启动失败"
  },
  "appName": "国际化支持",
  "needRoot": "需要root用户"
}
```

> 使用 json 驱动  
> 使用 T 方法获取语言, 第一个参数是路径，后面是动态参数，用于替换占位符号  
> 路径可以支持无限极，用 「.」 语法，就像 `JavaScript` 取对象属性
```golang

driver := driver.NewJsonFileI18nImpl("..\\lang")
i18n := go_i18n.NewI18N(go_i18n.Chinese, driver)
msg := i18n.T("install.error", i18n.T("appName"), i18n.T("needRoot"))

```

### 如何实现其他文件类型或数据库驱动
> 这是 json 的实现，你只需要实现 contract.I18NDriver 中的接口即可
```go
type JsonFileI18NImpl struct {
	langDir string
}

func NewJsonFileI18mImpl(langDir string) contract.I18NDriver {
	f := JsonFileI18NImpl{
		langDir: langDir,
	}
	return &f
}

func (Self *JsonFileI18NImpl) LoadLang(location string) []byte {
	fileName := fmt.Sprintf(Self.langDir + "\\%s.json", location)
	data, _ := ioutil.ReadFile(fileName)
	return data
}

```
> 欢迎各路大神提供不同的驱动