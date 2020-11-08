// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../datasource/migrations/1555589258_create_users.down.sql (19B)
// ../datasource/migrations/1555589258_create_users.up.sql (244B)
// ../datasource/migrations/1555589937_create_domains.down.sql (21B)
// ../datasource/migrations/1555589937_create_domains.up.sql (151B)
// ../datasource/migrations/1585465452_create_domain_users.down.sql (25B)
// ../datasource/migrations/1585465452_create_domain_users.up.sql (305B)

package migrationData

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __1555589258_create_usersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x50\x2a\x2d\x4e\x2d\x2a\x56\xb2\x06\x04\x00\x00\xff\xff\xc2\x23\x41\xc8\x13\x00\x00\x00")

func _1555589258_create_usersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1555589258_create_usersDownSql,
		"1555589258_create_users.down.sql",
	)
}

func _1555589258_create_usersDownSql() (*asset, error) {
	bytes, err := _1555589258_create_usersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1555589258_create_users.down.sql", size: 19, mode: os.FileMode(0666), modTime: time.Unix(1584012673, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3e, 0x16, 0xef, 0x60, 0x8a, 0x67, 0x69, 0x63, 0x7d, 0xcd, 0x2d, 0x38, 0xad, 0xc3, 0x7f, 0x3c, 0xd8, 0xfb, 0xeb, 0x6a, 0xfd, 0xf9, 0x41, 0x2c, 0x23, 0x70, 0x29, 0x6, 0x9a, 0xf6, 0xaa, 0xef}}
	return a, nil
}

var __1555589258_create_usersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xce\xb1\x6b\x83\x40\x14\xc7\xf1\x5d\xf0\x7f\xf8\xe1\xa4\xd0\xa1\x43\xb7\x4e\xb6\xbd\x41\x6a\x6d\x2b\xde\xe0\x24\x4f\xef\x41\x0e\x3c\x3d\xee\xce\x04\xff\xfb\x20\x31\x81\x24\xeb\xfb\x7d\xbe\xf0\x3e\x6b\x91\x37\x02\x4d\xfe\x51\x0a\x24\x8b\x67\xe7\x13\xa4\x71\x04\x00\x5a\x41\xca\xe2\x0b\xd5\x6f\x83\x4a\x96\x25\xfe\xea\xe2\x27\xaf\x5b\x7c\x8b\xf6\xe5\x42\xb6\x60\x22\xc3\x38\x92\x1b\x0e\xe4\xd2\xb7\xd7\x0c\xb2\x2a\xfe\xa5\xb8\x65\x3b\x65\x43\x7a\xbc\x73\x8f\x60\x1a\xdc\x6a\x03\xab\xce\x92\xf7\xa7\xd9\x29\xf4\x6b\x60\xda\x67\xbf\x58\x76\x1d\x29\xa3\x27\xf4\xf3\x3c\xee\xe7\xc1\x31\x6d\x0d\x05\x04\x6d\xd8\x07\x32\xf6\xfa\x9c\x55\x4f\x53\x1c\x65\xef\xe7\x00\x00\x00\xff\xff\x92\x5c\xd3\x71\xf4\x00\x00\x00")

func _1555589258_create_usersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1555589258_create_usersUpSql,
		"1555589258_create_users.up.sql",
	)
}

func _1555589258_create_usersUpSql() (*asset, error) {
	bytes, err := _1555589258_create_usersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1555589258_create_users.up.sql", size: 244, mode: os.FileMode(0666), modTime: time.Unix(1584012673, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x51, 0xc0, 0x73, 0x45, 0xde, 0xfb, 0x75, 0x14, 0x51, 0xc3, 0xeb, 0xd5, 0x8a, 0x2c, 0xdf, 0x12, 0x22, 0xdb, 0x71, 0x9e, 0xdb, 0x3c, 0x60, 0x3c, 0x41, 0xf6, 0xc, 0xc5, 0xf2, 0x53, 0xc7, 0xe3}}
	return a, nil
}

var __1555589937_create_domainsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x50\x4a\xc9\xcf\x4d\xcc\xcc\x2b\x56\xb2\x06\x04\x00\x00\xff\xff\xde\x49\xd2\xaf\x15\x00\x00\x00")

func _1555589937_create_domainsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1555589937_create_domainsDownSql,
		"1555589937_create_domains.down.sql",
	)
}

func _1555589937_create_domainsDownSql() (*asset, error) {
	bytes, err := _1555589937_create_domainsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1555589937_create_domains.down.sql", size: 21, mode: os.FileMode(0666), modTime: time.Unix(1584012673, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xee, 0x8d, 0x1d, 0x25, 0xa3, 0x7b, 0x22, 0x77, 0xe8, 0x9f, 0xb, 0xca, 0xe6, 0xc5, 0x8e, 0x2b, 0xa3, 0x9f, 0xcc, 0x60, 0xd9, 0xac, 0x27, 0xf4, 0x9b, 0x7c, 0x2f, 0xc1, 0xdc, 0x9b, 0x54, 0xf9}}
	return a, nil
}

var __1555589937_create_domainsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xca\x31\x0a\xc2\x30\x14\x06\xe0\xbd\xd0\x3b\xfc\x74\x52\xf0\x06\x4e\x51\xdf\x10\x8c\x51\x43\xde\xd0\x49\x1e\x4d\xc0\x0c\xa9\x25\x8d\x9e\xdf\xc1\xe2\xd2\xf9\xfb\x8e\x8e\x94\x27\x78\x75\x30\x84\x2e\xbc\xb2\xa4\x71\xee\xb0\x69\x1b\x00\x48\x01\xcc\xfa\x04\x7b\xf5\xb0\x6c\x0c\x6e\x4e\x5f\x94\xeb\x71\xa6\x7e\xf7\x2b\xa3\xe4\x88\x8f\x94\xe1\x29\x05\x6c\xf5\x9d\xe9\xdf\x97\x32\x94\x28\x35\x86\x87\x54\xd4\x94\xe3\x5c\x25\x4f\x0b\xbd\xa7\xb0\xa2\xb6\xd9\xee\xbf\x01\x00\x00\xff\xff\x40\x76\x25\xed\x97\x00\x00\x00")

func _1555589937_create_domainsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1555589937_create_domainsUpSql,
		"1555589937_create_domains.up.sql",
	)
}

func _1555589937_create_domainsUpSql() (*asset, error) {
	bytes, err := _1555589937_create_domainsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1555589937_create_domains.up.sql", size: 151, mode: os.FileMode(0666), modTime: time.Unix(1584012673, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0xd6, 0x4d, 0x4a, 0x31, 0xb9, 0x69, 0x46, 0x2e, 0x95, 0xc3, 0x1b, 0x6, 0x1a, 0xa1, 0xc3, 0x24, 0x5b, 0x17, 0x23, 0xe8, 0x99, 0xb4, 0x4, 0x52, 0x2, 0xe1, 0xb2, 0xce, 0xc5, 0x1d, 0x35}}
	return a, nil
}

var __1585465452_create_domain_usersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x50\x2a\x2d\x4e\x2d\x8a\x4f\xc9\xcf\x4d\xcc\xcc\x53\xb2\x06\x04\x00\x00\xff\xff\xcf\x77\xdc\x89\x19\x00\x00\x00")

func _1585465452_create_domain_usersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1585465452_create_domain_usersDownSql,
		"1585465452_create_domain_users.down.sql",
	)
}

func _1585465452_create_domain_usersDownSql() (*asset, error) {
	bytes, err := _1585465452_create_domain_usersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1585465452_create_domain_users.down.sql", size: 25, mode: os.FileMode(0666), modTime: time.Unix(1585466002, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xeb, 0xe6, 0x59, 0xfe, 0x7e, 0x5a, 0xec, 0xd4, 0xe9, 0xe3, 0x83, 0x8e, 0xfe, 0x66, 0xe3, 0x65, 0xee, 0xa3, 0x20, 0x6e, 0xbe, 0x87, 0xbc, 0x30, 0x3a, 0x31, 0x52, 0xc3, 0x43, 0xa3, 0x22, 0xae}}
	return a, nil
}

var __1585465452_create_domain_usersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\xbb\x6a\x84\x40\x14\xc6\xf1\x5e\xf0\x1d\x3e\xac\x14\x7c\x83\x54\x26\x99\x42\x62\x4c\x10\x2d\xac\xe4\xc4\x39\xc2\xc0\x5c\x64\x2e\xef\x1f\x50\x43\x16\x16\x96\x6d\xe7\xfb\xfd\xe7\xbc\x0d\xa2\x19\x05\xc6\xe6\xb5\x13\x28\x52\x60\xbf\x48\x67\x48\xd9\x02\x65\x9e\x01\xc0\xf1\xa6\x24\xa6\xa9\x7d\x47\xff\x35\xa2\x9f\xba\x0e\x9e\x37\xf6\x6c\x57\x0e\x07\x08\xa5\x92\x55\x7d\x06\x67\xff\x30\x39\xc9\x6d\xf4\x3d\xb4\x9f\xcd\x30\xe3\x43\xcc\xe5\x75\xb1\xfe\xff\xe9\x4f\x91\x34\xca\xe2\xc7\x39\x0d\xeb\x22\x6c\xd2\x1a\x92\x37\x4a\x3a\x62\x23\x1d\xf8\x72\x6c\x48\xe9\x27\xdc\xea\x99\x22\xcb\x85\x22\xa2\x32\x1c\x22\x99\xfd\x9a\xd2\x2e\xef\xa6\x3c\xab\x5e\x7e\x03\x00\x00\xff\xff\x71\xc0\xd9\xe0\x31\x01\x00\x00")

func _1585465452_create_domain_usersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1585465452_create_domain_usersUpSql,
		"1585465452_create_domain_users.up.sql",
	)
}

func _1585465452_create_domain_usersUpSql() (*asset, error) {
	bytes, err := _1585465452_create_domain_usersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1585465452_create_domain_users.up.sql", size: 305, mode: os.FileMode(0666), modTime: time.Unix(1585465593, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x33, 0xa4, 0x9, 0xca, 0xe1, 0x20, 0x58, 0x8e, 0x21, 0xdb, 0x2f, 0xe8, 0x19, 0xa0, 0x18, 0x71, 0x71, 0x6d, 0xe4, 0xb2, 0xe1, 0x5f, 0x11, 0xa6, 0xb3, 0xa6, 0xb3, 0x50, 0xab, 0xfe, 0x14, 0x13}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"1555589258_create_users.down.sql":        _1555589258_create_usersDownSql,
	"1555589258_create_users.up.sql":          _1555589258_create_usersUpSql,
	"1555589937_create_domains.down.sql":      _1555589937_create_domainsDownSql,
	"1555589937_create_domains.up.sql":        _1555589937_create_domainsUpSql,
	"1585465452_create_domain_users.down.sql": _1585465452_create_domain_usersDownSql,
	"1585465452_create_domain_users.up.sql":   _1585465452_create_domain_usersUpSql,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"1555589258_create_users.down.sql":        &bintree{_1555589258_create_usersDownSql, map[string]*bintree{}},
	"1555589258_create_users.up.sql":          &bintree{_1555589258_create_usersUpSql, map[string]*bintree{}},
	"1555589937_create_domains.down.sql":      &bintree{_1555589937_create_domainsDownSql, map[string]*bintree{}},
	"1555589937_create_domains.up.sql":        &bintree{_1555589937_create_domainsUpSql, map[string]*bintree{}},
	"1585465452_create_domain_users.down.sql": &bintree{_1585465452_create_domain_usersDownSql, map[string]*bintree{}},
	"1585465452_create_domain_users.up.sql":   &bintree{_1585465452_create_domain_usersUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
