package migrations

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

var __000001_init_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x41\x0e\xc2\x30\x0c\x04\xef\x91\xf2\x87\x3d\xb6\x12\x3f\xe8\x63\x2a\xd3\x58\x60\x35\x38\xc5\x49\x80\xfe\x1e\x91\x94\x03\xe2\x68\xcf\x6a\x76\x17\x63\x2a\x8c\x42\xe7\xc8\xa8\x99\x2d\x7b\x37\x78\x07\x00\x12\x90\xd9\x84\xe2\xa9\xdf\x31\x5d\x44\xf1\x20\x5b\xae\x64\xd0\x54\xa0\x35\x7e\xe1\x46\x39\x3f\x93\x85\x3f\xee\xdd\x38\x79\xe7\xdd\x51\x54\x55\xee\x95\x21\x1a\xf8\xd5\xfb\xe6\xe6\x9d\x6b\x7b\x75\x59\xd2\x8e\x30\x34\xd6\x05\x14\x0b\xdb\xef\xd0\x4f\x96\x42\xc0\x92\x34\x17\x23\xd1\x72\x28\xb7\xb5\xc3\xb6\xcc\xe4\x46\xb6\x63\xe5\x1d\x83\x84\x71\x7a\x07\x00\x00\xff\xff\xb7\xbc\xa1\x80\xf3\x00\x00\x00")

func _000001_init_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_up_sql,
		"000001_init.up.sql",
	)
}

var _generate_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xca\xb1\x0d\xc5\x20\x0c\x04\xd0\x9e\x29\x6e\x01\xa0\xff\xdb\xdc\x4f\xac\x13\x42\xb1\x11\xf1\xfe\x4a\x93\x22\xf5\x7b\x8b\xc7\xa4\x0c\xd7\xd0\x66\x8e\xf0\xbb\x94\xde\x15\x3f\x99\xdb\x66\x1a\x14\xf5\x3f\xfc\x64\x12\x75\x4d\x7d\x26\x6a\xe0\xa5\xa6\x00\x5a\x79\x02\x00\x00\xff\xff\xe6\x24\xd8\x86\x4e\x00\x00\x00")

func generate_go() ([]byte, error) {
	return bindata_read(
		_generate_go,
		"generate.go",
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
	"000001_init.up.sql": _000001_init_up_sql,
	"generate.go": generate_go,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_init.up.sql": &_bintree_t{_000001_init_up_sql, map[string]*_bintree_t{
	}},
	"generate.go": &_bintree_t{generate_go, map[string]*_bintree_t{
	}},
}}
