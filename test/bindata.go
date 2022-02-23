package test

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _lang_en_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\x3d\x0a\xc3\x30\x0c\x46\xf7\x40\xee\xf0\xa1\x39\x4b\xb7\x92\x43\x74\xe8\x0d\x4c\xaa\x96\x40\x62\x05\xc9\x86\x82\xf1\xdd\x8b\x9d\xf4\x67\x09\xcd\x26\xe3\x27\xbd\x97\xda\x06\xa0\xd1\x5b\x70\xd3\x44\x3d\xea\x1b\x20\x8b\xc3\xc0\x66\xd4\x83\x52\x5a\x54\x1e\xea\xe6\x9c\xb1\x81\x78\x7f\x77\x1b\xce\xaa\xa2\x7b\xf0\xdd\x8d\x53\x07\x65\x67\xe2\x7b\xa4\xb4\x4e\x39\x53\x59\xce\xf5\x04\x45\x7f\xb8\xe1\x83\x1e\xab\xf8\xe2\xa5\x63\x47\x6f\xc1\x69\xf8\xaf\xae\xd8\x31\xed\x8a\x16\xe5\xaf\xc7\x2d\xcb\xc5\xcd\x5c\xe0\xf1\x74\xf6\xeb\x09\xf2\xcc\xb7\xab\x48\x09\xa8\x33\x54\x24\x20\x1a\x2b\xf8\xc9\x43\x0c\x4c\x6d\x93\x5f\x01\x00\x00\xff\xff\x7e\x7e\x7d\x49\xab\x01\x00\x00")

func lang_en_json() ([]byte, error) {
	return bindata_read(
		_lang_en_json,
		"lang/en.json",
	)
}

var _lang_zh_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\xe5\x52\x50\x50\xca\xcc\x2b\x2e\x49\xcc\xc9\x51\xb2\x52\x00\xf3\x15\x14\x94\x8a\x4b\x93\x93\x53\x8b\x8b\x95\xac\x14\x94\xaa\xab\x0b\x8a\xf2\xd3\x8b\x12\x73\x6b\x6b\x9f\xae\xeb\x7c\xb1\xb8\xf5\x59\xc7\x84\xa7\x5d\xf3\x95\x74\xa0\x4a\x53\x8b\x8a\xf2\x8b\xb0\x2a\x7c\xba\x64\xe3\x8b\x2d\x4b\x75\x14\x20\xf4\xd3\xbe\xf9\x4f\x67\x2f\xb0\xaa\xae\x2e\x4a\x4d\x2c\xce\xcf\xab\xad\x55\x02\x19\x50\x0b\x36\x46\xa9\x34\x8f\x68\x37\xf4\xee\x78\xb1\x77\x2f\x31\x6e\x00\x2b\x84\xd8\x8d\x6c\x55\x71\x49\x62\x51\x09\x61\x6b\x26\xac\x7f\xda\xb5\x82\x18\x6b\xc0\x0a\x31\xad\x49\x2c\x28\xf0\x4b\xcc\x4d\x05\xa9\x7e\x3a\x7b\xef\xcb\x99\xad\x4f\x7b\xa6\x3d\x9b\xb2\xfe\x59\x4f\x23\xc4\x38\xa5\xbc\xd4\xd4\x94\xa0\xfc\x7c\x90\x53\x94\x5e\xce\x69\x78\xb1\xac\xb1\x28\x3f\xbf\xe4\xf9\x94\x15\xcf\x3a\xb6\x2b\xf1\x72\xd5\x02\x02\x00\x00\xff\xff\x36\x34\x36\xca\x9b\x01\x00\x00")

func lang_zh_json() ([]byte, error) {
	return bindata_read(
		_lang_zh_json,
		"lang/zh.json",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"lang/en.json": lang_en_json,
	"lang/zh.json": lang_zh_json,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"lang/en.json": &_bintree_t{lang_en_json, map[string]*_bintree_t{}},
	"lang/zh.json": &_bintree_t{lang_zh_json, map[string]*_bintree_t{}},
}}
