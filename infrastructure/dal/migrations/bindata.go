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

var __000001_init_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x41\x6e\xc3\x20\x10\x45\xf7\x48\xdc\xe1\xef\x6c\x4b\xde\x34\xdb\x1c\x26\x9a\x9a\x51\x3b\x0a\x19\xdc\x01\xd2\xe6\xf6\x95\x83\xe3\x24\x32\x0b\x04\xf3\xbe\xde\x87\xc9\x98\x0a\xa3\xd0\x67\x64\xd4\xcc\x96\xbd\xeb\xbd\x03\x70\xbf\x49\xc0\xb2\x53\x1c\x9f\xb3\x98\xbe\x44\x71\x25\x9b\xbe\xc9\xa0\xa9\x40\x6b\x7c\x0d\xcc\x94\xf3\x6f\xb2\xb0\xcb\x78\x37\x1c\xbd\xf3\x6e\x2d\xad\x2a\x3f\x95\x21\x1a\xf8\xaf\x75\x9f\x36\xff\xa9\xde\xc7\x4d\x9a\xb4\x61\xf4\x1b\x6f\x22\x8a\x85\xed\xfd\xf1\x4b\x9e\x42\xc0\x94\x34\x17\x23\xd1\xb2\xaa\xe7\x73\x83\xcb\x9a\x4d\x2e\x64\x37\x9c\xf9\xd6\x9c\x12\x9a\x50\x34\xb3\x15\x88\x96\xb4\xab\x1c\xdf\x3e\x37\xe0\x4a\xb1\xf2\xda\xd8\x77\x14\x2e\xa2\x1f\xdd\x88\xf5\xf4\xc8\x75\xc3\xf8\xa0\x87\x8d\x1e\x9e\xf4\xe8\xdd\x7f\x00\x00\x00\xff\xff\x48\x84\x83\x73\x84\x01\x00\x00")

func _000001_init_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_up_sql,
		"000001_init.up.sql",
	)
}

var __000002_init_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8d\x31\xaa\xc3\x40\x0c\x44\x7b\x83\xef\x30\xe5\xff\x90\x1b\xb8\x4d\x93\x03\xb8\x5e\x94\x48\x0e\xc2\xeb\xdd\x45\xab\x85\xf8\xf6\x21\x76\x11\x07\x12\x75\x62\x66\xde\xa3\xe8\x62\x70\xba\x46\x41\xab\x62\xb5\xef\x00\x80\x2d\x17\xdc\x72\xaa\x6e\xa4\xc9\xa1\x13\xe4\xa1\xd5\xeb\x5e\x0a\x65\x3e\x1d\x8a\x9f\xa9\xf2\xd0\x77\x7d\xf7\x83\x4c\xcc\xdb\x3b\x8e\x97\x33\x5a\x53\x46\xca\x8e\xd4\x62\xdc\xf3\xd7\x1d\xc4\x3c\x85\xdd\xa8\x0c\x96\x89\x5a\x74\xdc\x25\x05\xa3\xc4\x79\x09\x1b\xe0\xef\xff\xeb\xb4\xcc\xef\x69\x31\x5d\xc8\x56\xcc\xb2\x0e\xcf\x00\x00\x00\xff\xff\xf1\x42\x7f\x60\xf3\x00\x00\x00")

func _000002_init_up_sql() ([]byte, error) {
	return bindata_read(
		__000002_init_up_sql,
		"000002_init.up.sql",
	)
}

var _bindata_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindata_go() ([]byte, error) {
	return bindata_read(
		_bindata_go,
		"bindata.go",
	)
}

var _generate_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xca\x31\x0a\xc5\x20\x10\x04\xd0\x5e\xf0\x0e\x73\x01\xb5\xff\xb7\x99\x9f\x2c\x83\x48\x5c\x31\x7b\x7f\xd2\xa4\x48\xfd\xde\xe2\x31\x28\xc3\xd5\xb5\x19\xdd\xe7\x9d\x53\x4e\xad\xc9\x7f\xb2\x69\x9b\x61\x90\x97\x7f\x9f\x27\x83\x28\x6b\xe8\x73\x51\x1c\x2f\x55\x39\x50\x73\x7a\x02\x00\x00\xff\xff\x98\x70\xac\x16\x51\x00\x00\x00")

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
	"000002_init.up.sql": _000002_init_up_sql,
	"bindata.go": bindata_go,
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
	"000002_init.up.sql": &_bintree_t{_000002_init_up_sql, map[string]*_bintree_t{
	}},
	"bindata.go": &_bintree_t{bindata_go, map[string]*_bintree_t{
	}},
	"generate.go": &_bintree_t{generate_go, map[string]*_bintree_t{
	}},
}}
