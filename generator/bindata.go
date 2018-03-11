// Code generated by go-bindata.
// sources:
// generator/templates/comment.go.tmpl
// generator/templates/docs.md.tmpl
// generator/templates/vm_functions.go.tmpl
// DO NOT EDIT!

package generator

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

var _templatesCommentGoTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x41\x6f\xe2\x30\x10\x85\xef\xfc\x8a\x77\xc8\x61\x77\xb5\xe0\x3b\x15\x07\xa4\xa2\x56\x3d\x54\x15\xed\xbd\x32\xc9\xc4\x31\x25\x76\x34\x76\x28\x55\x94\xff\x5e\xd9\x4e\x81\x80\x38\xf4\x38\x33\x4f\x6f\xbe\x37\x23\x04\xba\x0e\xd9\xec\x59\xd6\x84\xbe\xc7\x34\x95\xf7\xe4\x72\xd6\x8d\xd7\xd6\xa0\xef\x27\x42\x4c\x84\xc0\x8b\xcc\x3f\xa4\xa2\xa1\x8a\xba\xa1\x75\xd2\x2c\x5b\x5f\x59\x1e\x8a\xc1\x2c\xf5\x4e\x9a\x27\xb9\x97\xc9\x7e\x68\x3c\x12\x13\xb4\x83\xaf\xe8\x6c\x88\x9a\x7c\x65\x0b\x38\xad\x8c\xf4\x2d\xd3\x3c\x68\x93\xe5\x9a\x72\xd2\x7b\xe2\x57\xcf\xda\xa8\xb3\xf5\xac\xda\x9a\x8c\x77\x17\xce\x12\x3b\xed\x3c\x6c\x19\x77\xc8\x1f\x15\x4a\xcb\x97\x5b\xcb\xd6\xe4\x21\xf7\x7c\xd2\x75\x53\xb0\x34\x8a\x90\xbd\xff\x47\x26\x31\x5f\x20\x9b\xad\x0e\x0d\xe5\x9e\x8a\x25\xab\xb7\xaf\x86\x5c\x5a\x0e\xfc\x8b\x64\xf2\x78\xc9\x3f\xa9\x7c\xb0\x41\x85\xbe\xff\x1b\xfd\xc8\x14\x27\xda\x35\xf9\x96\xcd\x6d\xd6\x52\xd3\xae\x70\xd0\x26\x32\x72\x54\xc3\x6e\xb6\x94\xfb\x2b\x3a\x1e\xd3\x25\xeb\x0b\x40\xbb\xd9\xce\x02\x15\x8f\x21\x79\x0c\x19\x19\xa7\x47\xc8\xd5\x41\xd6\xcd\x2e\x7c\x1d\x23\x4a\x03\x4a\x93\x40\x5a\xd9\x4f\x78\x8b\xd6\x11\x7c\xa5\xdd\xf1\x88\x81\x5d\xa5\xc3\xa6\xef\xed\x25\x07\x0c\x2c\x6e\xfc\xf1\xee\xf7\xb9\x84\xb8\x0e\x36\x8e\xf1\x1d\x00\x00\xff\xff\xc1\x69\x89\x83\xe5\x02\x00\x00"

func templatesCommentGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCommentGoTmpl,
		"templates/comment.go.tmpl",
	)
}

func templatesCommentGoTmpl() (*asset, error) {
	bytes, err := templatesCommentGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/comment.go.tmpl", size: 741, mode: os.FileMode(420), modTime: time.Unix(1520657252, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDocsMdTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xc1\x4a\x03\x31\x14\x45\xf7\xf9\x8a\x0b\x33\xab\x60\xf2\x01\x82\x8b\x82\xd6\x8d\x56\xa8\xe2\xd6\xc6\xe4\x65\x88\xd8\xcc\x98\x64\x8a\x12\xf2\xef\x92\x74\x06\xc1\x8d\xb8\xcc\x7b\xb9\xe7\x3c\x6e\x87\x9d\x4a\xee\x44\xd8\xce\x5e\x27\x37\xfa\xc8\x58\xce\x08\xca\x0f\x84\xfe\xe5\x02\xbd\xc5\xe5\x15\x7a\x59\xf7\x11\xa5\x30\xd6\x75\x38\xe4\x8c\xde\xca\x3d\x69\x72\x27\x0a\x8f\x29\x38\x3f\xa0\x94\x43\x0b\xf7\x56\x5e\x53\xd4\xc1\x4d\x15\xb8\x64\x3a\x6c\xc2\x30\x1f\xc9\x27\xdc\xb9\x98\x7e\x49\x54\x93\x58\x79\xf3\x39\x91\x4e\x64\x36\x61\x78\xfa\x9a\xa8\x09\xc1\xc1\x79\xc5\x2a\xb9\x53\x47\x42\x29\x9c\x63\x19\xdc\x8e\xf5\x5b\x1d\xb1\x9c\x05\xc8\x9b\xd5\xb6\xa7\x34\x07\x4f\x06\x0f\xaf\x6f\xa4\x13\xb6\x8e\xde\x4d\xfc\x43\x7b\x0e\xfd\x98\x51\xa1\xce\x82\x3e\xaa\x6c\xd9\x8e\xcf\xf7\x48\x61\xa6\x7f\xdc\x76\x06\x91\x37\x10\xa5\xac\x97\xb6\x62\x84\x10\xad\xb3\xf5\xcd\xbe\x03\x00\x00\xff\xff\xb9\xcc\xb7\x2c\x91\x01\x00\x00"

func templatesDocsMdTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDocsMdTmpl,
		"templates/docs.md.tmpl",
	)
}

func templatesDocsMdTmpl() (*asset, error) {
	bytes, err := templatesDocsMdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/docs.md.tmpl", size: 401, mode: os.FileMode(420), modTime: time.Unix(1520657252, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVm_functionsGoTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x51\x6f\xa4\x36\x10\x7e\xe7\x57\x4c\xdd\x9c\x04\xd5\x06\xfa\xbc\xba\x9c\x94\x46\x49\x5f\xb2\xa7\x53\x12\x6d\x1f\x4e\xa7\x93\x17\x06\xd6\x3d\x63\x53\x33\x6c\xb2\x45\xfc\xf7\xca\x36\xb0\xb0\x49\xae\xdb\x56\x51\x1f\x56\x0b\xf6\xcc\x7c\x33\xdf\xcc\x67\x93\x24\xf0\x89\xa7\xdf\x78\x81\x80\xaa\x10\x0a\x41\xd4\x40\x5b\x84\xaa\x5f\xa5\x2d\x27\x10\x65\x25\xb1\x44\x45\x7e\xaf\xa8\x53\x23\x2a\x82\xb5\x30\xd4\x70\x09\x2b\x9e\x6e\xad\x6b\xb8\x5e\x45\x71\x90\x24\x60\x7f\xeb\x15\xdc\x34\x2a\x25\xa1\x55\x7d\x08\x80\xd9\x32\x68\x5b\x30\x5c\x15\x08\x67\x52\x6c\x3e\xf2\x12\x17\xee\x09\x96\x17\x70\x16\xdf\x8a\x4d\x0d\xe7\x5d\x17\x24\x89\x8d\x72\x2b\x36\x86\x9b\x3d\xb4\xed\x68\x0d\xe3\xe6\x24\xbe\x3a\xb2\x98\xa2\x7c\x5d\xc0\x59\xee\xa2\x4b\xb1\x89\xad\xd3\x80\x00\xce\x2b\x8f\xef\x30\x45\xb1\x43\x73\x4f\x46\xa8\x02\xba\x0e\xce\x61\x4b\x54\xd5\xcb\x24\x29\x74\xa6\xd3\x58\x9b\x22\x29\x04\x6d\x9b\x4d\x9c\xea\x32\x29\x50\xfd\x9c\x8a\x0c\x93\x9e\x8a\xc4\x93\x97\xfc\x78\xed\xfe\x63\x1f\x76\xc8\xb6\x6d\x01\x55\xe6\x30\x27\x8f\x49\x12\x0c\x24\x5b\xf3\x78\xe8\x43\xd7\x05\x81\x28\x2b\x6d\x08\xc2\x00\x80\x55\x9c\xb6\x49\x2e\x24\xda\x07\x16\xd8\xa5\x49\x26\x46\x6f\xd0\xd0\x37\x23\x4a\x54\x89\x26\xd2\x2c\x88\x82\x20\x6f\x54\x0a\x21\xc2\x4f\x3e\x9f\x08\xae\x0c\x72\xc2\xf5\x2a\x8c\xa0\x0d\x00\x30\x5e\xaf\xe0\x02\xac\x7d\xfc\x11\x1f\xc3\xc8\xad\x09\xf5\x3b\xa6\xb4\xe6\xa6\x76\x0b\x2f\x31\x38\xa1\xcf\x47\x89\xef\x91\x42\x36\xab\x97\x2d\x00\xe3\x5d\x39\x5b\xeb\xe3\x0d\xb5\x03\x7c\x5d\x00\x1a\x63\x63\xba\x28\x77\x8d\x0a\x77\xe5\x27\x83\x52\xf3\xcc\x1a\x8b\xdc\xed\xff\x70\x01\x4a\x48\x97\xb3\xc5\xbb\xd5\x45\x81\x26\xfe\x4d\xd0\xf6\x46\xa0\xcc\x42\x46\x86\xa7\xc8\x16\xc0\xc8\x34\xc8\xa2\xf8\x86\x13\x97\x79\xc8\xee\xf7\x8a\xf8\x93\x8d\xa1\x8d\x1d\x8f\xca\x87\x5e\xc2\xbb\x9a\x39\xe8\xf8\xda\x6e\x85\x91\x05\xeb\xfa\xf2\x05\x09\x2e\xc5\x9f\xe8\x61\xc2\x28\xe8\x82\x60\x32\x48\x22\x7b\x3a\x26\xa2\xeb\x9e\x73\x7d\x54\x7a\x98\x72\x29\x3d\xd5\xc3\xc0\x5e\x71\x29\x23\xbf\xb4\xe6\xb2\x41\x57\x9f\xc8\x41\xa2\x72\xd6\xf1\xa5\x29\x1a\x2b\x97\x5b\x51\x53\x04\x1f\x2c\x75\x12\x95\x0d\x7a\xfd\x54\x61\x4a\x98\x5d\x9a\xe2\x61\x5f\xa1\xcd\xe0\x3b\xec\xe4\x3d\xa0\x25\x68\xde\xa3\xe8\x7b\x24\x7a\x6a\xd8\x83\xd6\x50\x72\xb5\x07\xde\xa7\xe3\x84\xe6\x12\x64\x91\xc3\x34\x48\x8d\x51\x7d\x71\x5c\xd6\xe8\xca\x09\x07\x4e\x5f\x2f\xe9\xfd\xff\x58\x52\x8e\x8f\xff\xba\xa2\xc3\x30\x70\x3f\x0d\xdc\x4d\xc3\x0b\x45\xf8\x31\x6f\xdb\x73\x37\xca\x7f\xc0\x19\x8f\x7f\xd5\x76\x0b\xd8\xe7\x2f\x9b\x3d\x21\xb3\x75\x82\x8f\x79\xc6\x87\x2a\x7a\x45\x58\xd0\x07\xfd\xcb\x9e\xf0\x5e\x8a\x14\xe7\x14\x86\xce\x43\x64\x4f\x56\x58\x83\xb2\x64\xed\x8e\x0e\x80\x1d\x37\x47\x21\xfd\x5b\x8f\xee\x6c\x0c\x7f\xbc\x34\xc5\x24\xcc\xa8\xc6\xd7\x81\x6c\x89\xda\x50\xf8\x4f\xd4\xf9\x5f\x9a\x95\x87\xec\x4a\x37\x32\x03\xa5\x09\xd0\x61\x43\x6e\xed\x7b\x0d\xb3\x59\x91\xa7\xb5\xaf\x7e\x14\x94\x6e\x61\x67\x2b\x7d\xc6\x41\x1c\xd2\xbe\x42\x7f\x40\xa6\xbc\xc6\x63\xde\x96\x0e\x61\xd6\x51\x07\xee\x4e\x68\xe6\x89\x3d\xee\xe6\x05\x0c\x27\x77\x7c\x25\x91\xab\xf0\x05\xd4\x23\x98\x28\x1a\xe2\x1c\x7a\xfa\x3c\xec\x09\x71\xc6\x30\xe3\xa1\x9b\x61\xce\x1b\x49\xcb\xb7\xea\xd7\x30\x37\x60\x89\x84\x52\xd4\x25\xa7\x74\xbb\xb4\xdd\x73\xe2\x80\x77\xf5\x02\x0a\x4d\xf0\xee\xe1\xd0\xc0\x31\x61\xb6\x80\xdd\xa9\x22\x3c\xd4\x74\xfc\xe2\xae\x73\xeb\x7e\xb8\xcc\x9d\xa4\x5e\xb9\xe9\xbd\x18\xd6\xab\x3b\x24\x6b\x67\x1f\xea\x4a\xab\x1a\xdb\xb9\xdc\x8d\x97\xbb\x39\x92\xbb\x87\x9a\x2a\x7e\x36\x21\x66\xd4\xbc\xbb\x88\xc6\x21\x11\xb9\x4b\xd5\x8c\x0d\x9d\x29\xe9\x4d\x7a\xf3\x7e\xf0\xf7\x77\xe2\x07\x2f\xa2\x59\x16\xd3\x2b\xd1\xb5\xa1\x27\xe6\x33\x9b\x99\xb1\x2f\x70\xf1\xb2\xa3\xf3\xeb\xfc\xdc\xb6\x27\x04\x51\x42\x7a\x97\x97\x26\xfe\x64\xf4\xe7\x73\x3e\x7b\xd9\x95\x77\x48\x0b\xff\xe7\xf2\x1c\x3f\x3a\x1e\xb4\x9f\xab\x01\xa9\x3f\xd9\x26\x96\x6f\x7d\xc0\xf9\xf1\x81\x54\xab\x1d\x9a\xda\x36\x27\xe7\x42\xe2\x70\xc4\x1d\x32\x99\xb7\xe6\x6f\xf4\xd1\x6f\x3b\x6f\xfb\x09\xd3\xd3\xd1\x75\x7f\x05\x00\x00\xff\xff\xf3\xb3\x01\x5a\xee\x0b\x00\x00"

func templatesVm_functionsGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesVm_functionsGoTmpl,
		"templates/vm_functions.go.tmpl",
	)
}

func templatesVm_functionsGoTmpl() (*asset, error) {
	bytes, err := templatesVm_functionsGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vm_functions.go.tmpl", size: 3054, mode: os.FileMode(420), modTime: time.Unix(1520726271, 0)}
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
	"templates/comment.go.tmpl": templatesCommentGoTmpl,
	"templates/docs.md.tmpl": templatesDocsMdTmpl,
	"templates/vm_functions.go.tmpl": templatesVm_functionsGoTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"comment.go.tmpl": &bintree{templatesCommentGoTmpl, map[string]*bintree{}},
		"docs.md.tmpl": &bintree{templatesDocsMdTmpl, map[string]*bintree{}},
		"vm_functions.go.tmpl": &bintree{templatesVm_functionsGoTmpl, map[string]*bintree{}},
	}},
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

