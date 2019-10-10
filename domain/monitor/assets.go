// Code generated by go-bindata.
// sources:
// assets/nodes.gotemplate
// DO NOT EDIT!

package monitor

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

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
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

var _nodesGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xdb\x8e\xe3\x36\x0c\x7d\x9f\xaf\x60\xbd\xaf\x23\x3b\x41\xda\xe9\xac\xd7\x36\xb0\x28\x50\xf4\x69\x51\xa0\x5f\xc0\x58\xb4\xad\xae\x2c\x19\x92\x72\xab\x91\x7f\x2f\xe4\x4b\x26\xbe\x64\xba\x5d\xa0\x40\xfd\xc4\x88\x22\x79\x78\x48\x91\x49\x2a\x57\x4b\x90\xa8\xca\x34\x20\x15\x64\x4f\x49\x45\xc8\xb3\x27\x00\x80\xa4\x26\x87\x90\x57\x68\x2c\xb9\x34\x38\xb8\x82\xbd\x06\xd1\xbd\x4e\x61\x4d\x69\x70\x14\x74\x6a\xb4\x71\x01\xe4\x5a\x39\x52\x2e\x0d\x4e\x82\xbb\x2a\xe5\x74\x14\x39\xb1\xee\xc7\x33\x08\x25\x9c\x40\xc9\x6c\x8e\x92\xd2\x6d\xb8\xb9\xf9\x72\xc2\x49\xca\xda\x36\xfc\x83\x8c\xb7\x08\x7f\xd1\xaa\x10\x65\xf8\x05\x6b\xba\x5e\x81\x81\xd2\x9c\x6c\x12\xf5\xf7\x7a\x1b\xeb\x2e\xa3\xec\xbf\xbd\xe6\x17\x68\x6f\x3f\xfd\x57\x68\xe5\x58\x81\xb5\x90\x97\x18\x3e\x1b\x81\xf2\x19\x7e\x23\x79\x24\x27\x72\x7c\x06\x8b\xca\x32\x4b\x46\x14\x9f\x26\x66\x28\x45\xa9\x98\x70\x54\xdb\x18\x72\x52\x8e\xcc\xf4\x02\x17\xb6\x91\x78\x89\xa1\x90\x74\x9e\xaa\xfe\x3c\x58\x27\x8a\x0b\x1b\x78\x58\xda\x5f\x6f\x52\x98\xa3\xe1\x33\xc4\x35\x9a\x52\xa8\x18\xb6\x2f\xcd\xcc\x71\x83\x9c\x0b\x55\xc6\xf0\xd3\x5c\x53\x0b\xc5\x2a\x12\x65\xe5\x62\xd8\x6d\xe6\xda\xbd\x36\x9c\x0c\x33\xc8\xc5\xc1\xc6\xb0\x5b\xea\xcf\xcc\x56\xc8\xf5\x29\x86\x4d\x73\x86\x6d\x73\xf6\x97\xe0\xc3\xcf\xdd\xb7\x0a\xdc\x17\x7d\x8d\x6a\x2b\xfe\xa2\x18\x24\x9a\x92\x3e\x2d\xb5\xa7\x01\xe3\x5e\x4b\xbe\x9a\x1b\x73\xba\x59\x41\x38\x6a\xf7\xda\x39\x5d\xcf\x08\xb8\x43\x65\x0f\x7b\x45\xee\xbd\x16\xa8\xb5\xd2\xb6\xc1\x9c\xd6\xed\x1d\xba\x83\x65\x0d\x29\x1f\x6e\xee\x47\x6a\x74\x31\x18\x9f\xc2\x83\xc2\x2c\x80\xf7\xc5\xec\xb3\x62\xcb\xb2\xf5\x5a\xd3\x93\xb2\xd4\xcf\x0a\xb7\xe9\xaa\xb2\x59\x96\x3f\xd7\x52\x9b\x18\x3e\xec\x76\x2f\x2f\x1f\x3f\x4e\x75\x8e\xce\x8e\x71\xca\xb5\x41\x27\xb4\x8a\x41\x69\xb5\x9e\x7c\xa5\xad\x2b\x84\x5c\x94\xf5\x7f\x9b\xf6\x1e\xf3\xaf\xa5\xd1\x07\xc5\xd9\x7b\x0c\x0c\xba\x53\x25\x1c\x7d\x2f\x37\xfe\x2d\xa3\x50\x64\x66\xe4\xdc\xc6\xc0\x5e\xea\xfc\xeb\xd4\xfb\x91\x8c\x1f\x32\x92\x75\xc3\x24\x86\x5a\x70\x2e\x67\x08\xba\xa9\x18\xc3\xeb\x66\xf3\xa0\xa3\x9b\xc3\x5e\x8a\x7c\xad\xa3\xc7\xb7\xb4\xdd\x6c\xe0\x07\x51\xfb\xd9\x8b\xca\xad\x37\xc6\xeb\x8f\xb8\xa3\xed\x6a\x80\x6e\x84\x7f\xeb\x04\x7a\xec\x6f\xea\x90\x21\xe7\x86\xac\xfd\xd7\x2f\x71\x86\x2b\x83\x39\xb6\x7f\x28\xd9\x7b\x20\x17\xcc\x4d\xa7\xd0\x5a\xf0\xb8\xd2\xc7\x45\xcd\x17\x10\x0e\x8a\x93\x91\x62\xd9\x3a\x49\x34\x2c\xa8\x24\xea\x17\x6a\xe2\x37\x54\xf6\xf4\x94\x70\x71\x84\x5c\xa2\xb5\x69\x70\x6b\xad\x60\x58\x6a\x77\x3a\x8f\x23\x78\xdb\x6f\x09\x42\x65\xa8\x48\x83\x30\x8c\x02\xe0\xfa\xa4\xa4\x46\x9e\x06\x8f\x96\x66\xf0\x70\x9d\x26\x11\x66\xcf\x90\xd8\x06\xd5\x7d\xac\xb1\x6e\x13\xc3\xcf\xfd\x99\x37\xf2\xf7\x07\x94\x11\x17\xc7\x5e\x6c\x5b\x83\xaa\x24\x08\xbf\xf8\x05\x7d\x7d\xeb\xad\xb6\x15\x05\x84\xbf\x92\xcb\x2b\xe2\x77\xe7\xf3\x2c\xfd\x12\xbc\xcb\x72\x91\xad\x69\xf2\xc8\x2f\xff\xa8\x6d\x07\xf8\xd1\x38\xab\x82\xd1\xc7\xdb\xc1\x84\x96\x91\x87\xf1\xd0\xe7\xbd\x0c\x74\x87\xc5\xef\xb5\x2e\xfb\xde\xb2\x4f\xe1\xf7\xee\x11\x5e\xaf\x13\xbe\xfa\x97\x19\x64\xc0\xa0\x17\x07\x7a\xda\x96\x14\xf7\x64\xdd\x08\x7a\x14\xab\xdf\x56\x3d\xd7\x9d\xb8\x6a\x35\x3b\x6a\x5b\x92\x96\xbe\x83\xcd\xfb\xc0\x93\x35\x17\x64\x83\x10\x86\xe1\x37\xa0\x9e\x31\xf4\x1f\xe6\xa9\xc6\xa6\x19\xe5\xe1\x4a\x12\xf5\xcf\x28\x89\xfc\xff\xd6\xec\xef\x00\x00\x00\xff\xff\x6e\xb7\x7a\x24\xbe\x0a\x00\x00")

func nodesGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nodesGotemplate,
		"nodes.gotemplate",
	)
}

func nodesGotemplate() (*asset, error) {
	bytes, err := nodesGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes.gotemplate", size: 2750, mode: os.FileMode(420), modTime: time.Unix(1570516409, 0)}
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
	"nodes.gotemplate": nodesGotemplate,
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
	"nodes.gotemplate": &bintree{nodesGotemplate, map[string]*bintree{}},
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