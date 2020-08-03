// Code generated by go-bindata.
// sources:
// template.go.tpl
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateGoTpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\xdf\x4f\xdb\x30\x10\x7e\x8e\xff\x8a\x23\x42\x53\x8a\x42\x78\xaf\xd4\x07\x44\xb7\x09\x69\x94\xaa\x63\xe3\x01\x21\x88\xd2\x4b\x9a\xb5\xb5\x2b\xc7\xd9\x86\x2c\xff\xef\xd3\x39\x4e\x62\xfa\x83\x0a\xd6\x97\x3a\x67\xdf\xdd\x77\xdf\x9d\x3f\x5f\x5c\xc0\xf8\x16\x26\xb7\x77\xf0\x79\x7c\x7d\x77\x02\x57\x62\x8e\x50\x20\x47\x99\x2a\x9c\x27\x6c\x93\x66\xcb\xb4\x40\x90\x98\xa3\x44\x9e\x21\x63\xe5\x7a\x23\xa4\x82\x88\x05\x5a\x83\x4c\x79\x81\x90\x5c\x5b\x5b\x05\xc6\x30\x00\x00\xad\xcf\x21\xb9\x5c\x95\x29\x59\x20\xd4\x1a\x92\xe9\xb2\x98\xa6\x6a\x01\xc6\x84\xd6\x11\xf9\x9c\x4e\x0f\x18\xcb\x6b\x9e\x41\xc9\x4b\x15\x0d\x40\xb3\x40\x62\x51\x56\x0a\x65\x32\x73\x8b\x28\x94\x98\x87\x31\x4c\xf0\x0f\x05\xba\x7b\xd9\x20\x18\x33\xc3\xfc\xbe\x54\x8b\x2b\xc1\xf3\xb2\x18\xec\x75\x9b\xef\x71\x1b\x1f\x77\xe3\x82\x63\x18\x03\x7f\xe5\x37\x11\x1c\x07\xcc\x30\xa6\x35\x9c\x5a\xdb\x70\xd4\x6d\x5a\xe3\x74\x59\x58\x1b\xfd\x1b\xc3\x98\xa2\xbd\x26\x0f\x54\x4a\xd6\x99\xa2\xea\x26\xe9\x1a\xe9\xb3\xe4\x05\x0b\xc6\x98\x13\x57\xad\x4f\xe2\xe5\x83\xe7\x5f\x95\xe0\xc3\x30\x16\xeb\x52\xe1\x7a\xa3\x5e\xc2\x67\x4a\x6f\xc9\x7a\x83\x89\x28\x13\x3c\x87\x33\x57\xdf\xa1\xe0\x9a\x05\x82\xc0\x7e\xf2\x8c\x9a\x05\x16\xdd\x10\x28\x44\x42\xcb\x98\x05\x84\x71\x08\x8d\x69\x8c\x79\xcc\x02\x43\xa4\xa9\x5a\x72\x10\x07\x00\x8d\x3f\x04\xc8\x05\xf5\xac\xd3\x5a\x45\x3d\x96\x0e\x83\x6d\xc3\xef\x54\x42\x64\x67\x6d\x5d\x2b\xa8\x5e\x78\x96\xcc\xee\x6f\x6a\x85\x7f\xad\xf1\xc9\x0b\xf3\x5d\x09\x89\x30\x82\x75\xba\x79\x68\x98\x7f\xdc\x8f\x42\xf7\xe3\xb8\x85\x82\xf7\x5d\x8b\x61\x8e\xf9\x81\x32\xde\x28\x8f\x40\x95\xb9\xf5\x1d\x8d\x80\x97\x2b\x67\xa3\x9f\x35\xc2\xd6\xb0\xd9\x5d\xd3\x56\x98\x7c\x13\xd9\x32\x1a\xb0\x60\xa7\xb0\x07\xc2\xf6\x08\x23\x8a\xc2\x02\x3a\xfa\x83\xaf\xdc\x61\x47\x29\xed\x98\xdd\xc2\xbe\xe2\x6e\x61\x69\xbd\x52\xd5\xc7\xaa\xa3\xd4\xb3\x16\xa6\x88\x41\x2c\x69\xc2\x0e\xe0\x6d\x90\xce\x7a\xa8\x65\x4e\x0e\x34\x83\xdd\x70\xd1\xa4\x35\x94\x35\xa8\x4e\x5a\xde\x02\xaf\x30\xbb\xe5\x0f\xe5\x36\x8d\x6e\x56\xb6\xcc\xaf\x70\xd9\x73\xcd\x7d\xdd\x36\xbb\x9b\xab\x5b\xfa\x76\x45\x21\x7a\xdf\x4c\xb7\x98\x7a\xe9\xbc\x41\xb5\x10\x73\x2b\x9d\x36\x45\xf4\xd4\x49\x8c\x93\x1d\xeb\xff\xa5\xe6\x99\x15\x0f\x63\xa2\x4e\x63\x5d\x88\x4b\x59\x54\x70\x6e\x0c\x0b\x82\x27\x3f\x59\xdc\x1d\x24\xb5\xb5\x07\x06\xb0\xe3\x3d\xc3\xca\xf2\xdb\x04\xd0\x9a\xe6\x34\xf9\x99\xae\xea\xa6\x0a\x0d\x89\xcb\x4b\x6b\x5c\x55\x76\x69\xf3\x34\x1a\x7e\x34\xa3\x66\xc1\x4a\x14\x05\xca\xe4\x3e\x95\x3c\x0a\xd5\xa2\xac\xa0\xac\x80\x84\x16\x84\xbd\x4c\xa7\x3d\x7f\x6d\xe9\xe1\xc0\x7f\x60\x5a\x90\xdd\x03\x73\x18\x65\x73\x95\xfc\x2d\xf7\xd6\x38\x47\xf7\xe5\x9a\xc3\xac\x7a\x3b\x5b\x33\x04\x7e\x13\xdf\xaf\xdc\x47\xba\x2b\xe0\xcc\x2b\xf2\x1d\xbd\x6d\xd5\xc2\xbe\xad\x5e\x47\x3e\xde\xec\x3e\xe2\xb1\xfe\x79\x94\xf7\x8d\x68\xf8\x7b\x35\x06\xa7\x9e\xb4\x08\x27\xdb\xc2\x6a\x76\xf2\x5f\x75\xee\x83\xe5\x37\xee\x5f\x00\x00\x00\xff\xff\x4b\xe4\x22\xf7\xc6\x08\x00\x00"

func templateGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templateGoTpl,
		"template.go.tpl",
	)
}

func templateGoTpl() (*asset, error) {
	bytes, err := templateGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template.go.tpl", size: 2246, mode: os.FileMode(420), modTime: time.Unix(1596609141, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"template.go.tpl": templateGoTpl,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"template.go.tpl": &bintree{templateGoTpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
