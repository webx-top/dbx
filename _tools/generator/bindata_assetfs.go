// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package main generated by go-bindata.// sources:
// template/dbschema.gotpl
// template/dbschema_init.gotpl
// template/model.gotpl
// template/model_base.gotpl
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _templateDbschemaGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1a\x5b\x6f\xdb\xba\xf9\x59\xfa\x15\xac\x51\x14\x54\xa1\x2a\x7d\x18\xf6\x90\x22\xc3\x9c\xb6\x39\x0b\xb2\xa6\x45\xd2\xd3\x0d\x08\x82\x82\x96\x3e\x39\x9a\x65\xd1\xa3\x68\x3b\x3e\x82\xfe\xfb\xc0\x9b\x44\x59\x92\x25\xa7\xc1\x41\xcf\x70\x80\xa2\xb1\xc8\xef\xce\xef\xc6\xcb\xc9\x09\xfa\xfb\x1c\x32\x60\x84\x43\x84\x3e\x50\x94\x51\x8e\x20\x4a\x38\xe2\x0f\x49\x8e\xe2\x24\x05\x1f\x6d\x1f\x92\xf0\x01\x25\x39\x22\x6b\x4e\x97\x84\x27\x21\x49\xd3\x1d\xaa\xf1\x66\x3b\xc4\x1f\xc0\x0c\x50\x16\xb8\xee\x8a\x84\x0b\x32\x07\x54\x14\xc1\x17\xf5\xf3\x9a\x2c\xa1\x2c\x5d\x37\x59\xae\x28\xe3\x08\xbb\xce\x24\x5e\xf2\x89\xeb\x14\x05\x23\xd9\x1c\xd0\xcb\x85\x8f\x5e\x6e\xd0\xe9\x19\x0a\x2e\x25\x4c\x5e\x96\xae\x33\x29\x8a\x97\x9b\xb2\x94\x70\x90\x45\x82\x82\x33\x99\x27\xfc\x61\x3d\x0b\x42\xba\x3c\xd9\xc2\xec\xf1\x0d\xa7\xab\x93\x68\x36\xe9\x9d\x39\x49\x93\xd9\x49\x4c\x42\x4e\xd9\xae\x0f\x0a\xc2\x07\x7a\x68\xee\x64\x45\x18\x59\x4e\x5c\xcf\x75\xf9\x6e\x05\xe8\x36\x4d\x42\xf8\x5e\x14\xc1\x2d\x67\xeb\x90\x2b\xf5\xd0\xdd\xfd\xeb\xbd\x21\xd7\x8d\xd7\x59\x88\x70\xde\x8d\xe1\xa1\x1b\xa1\x3d\x8e\x33\x24\xe0\xf0\x12\x69\x39\x83\x4f\x34\x82\xd4\x43\xc0\x18\x65\xc8\xfc\x2d\x5c\x27\xa6\x0c\x7d\xf7\x91\xb4\x94\xb2\x5c\x2e\x86\x9d\x24\x16\x30\x62\x34\xce\xf0\xc6\x7b\x27\xbf\x5e\x9c\xa1\x2c\x49\xe5\xbc\xc3\x80\xaf\x59\x26\x86\x5d\xc7\x29\x5d\xf1\x4f\x0f\x65\x49\xea\x8e\x14\xf4\x86\x6c\x2d\x59\xf7\x95\xfd\x99\xc4\xfd\x85\xd1\xf5\xea\x7c\x87\x17\xb0\xbb\x48\x20\x8d\x50\xce\x59\x92\xcd\x3d\xb4\x24\xab\x3b\xf5\xfb\xbe\xbd\x5c\x82\xb7\x94\xea\x20\x54\x51\x56\x8a\x31\xba\xdd\x57\x2d\x5a\x92\x95\x1c\xa3\xdb\x60\x9a\x7f\x22\x2b\xec\xb9\x8e\xb3\x59\xc0\x4e\xaa\xbb\xe4\xc1\xed\x8a\x25\x19\xc7\x02\xf0\xce\xc8\x77\xef\x29\xab\x7c\xf7\x91\x84\x63\x77\x02\xe3\xfe\x1d\x7a\xb1\xd3\x06\x51\x03\xe8\xac\xc3\xcb\x84\x40\xc2\x42\x16\x10\x59\xad\x20\x8b\xb0\x1e\x90\x82\x7a\xb6\x15\xd9\x08\x1b\x5e\xc1\x6e\xc0\x82\x63\xec\xf7\xbb\x59\xaf\x56\x9e\xd1\xed\xb1\xba\x4e\xf3\xab\x6f\xfb\xaa\xfa\x68\x43\xd2\x35\x34\xb5\x97\x89\x20\xb8\xe5\x94\x41\xa5\xad\x35\xf6\x3b\x68\x27\xe7\x6a\xc9\xee\x8f\x55\xf5\x2b\x23\x59\x1e\x53\xb6\xc4\x5c\xfe\x02\x96\xdb\xeb\xa5\x94\xf9\xaa\xa7\x3c\x74\x77\xdf\xa5\xf2\x92\x2c\x00\x37\xa6\xfc\x14\x32\x9c\x7b\x9e\xd2\x3f\x89\x1e\x3b\x2d\xc0\xee\x92\xe8\x51\xaf\x91\x31\x41\xd0\x21\x51\xdb\x59\x87\xd4\xba\x60\x74\xf9\xcf\x24\xe7\x38\x22\x9c\xa0\x24\xe3\xc0\x62\x12\x42\x51\x7a\x3d\xd9\xba\x70\x1d\x69\xc5\xdc\x47\x74\x21\xe4\x14\x88\x01\x6e\x87\x97\xe7\x8a\xc8\x7c\x41\x17\x52\x03\x93\xd2\x04\x6a\xad\x9d\xc1\x6d\xf0\x55\x81\xab\x8c\xf0\xaa\x2b\x66\xc5\x64\x20\x04\xbf\xa1\x5b\x2c\x29\x06\xd8\x5a\x0a\x9b\x98\xf0\x01\x27\xaf\x23\x3b\x37\x31\xad\x02\x5f\x19\x2a\x97\x56\x6b\x42\x29\x1d\x83\x20\xf0\x5c\xc7\xad\x01\x4b\xd7\x3d\x39\x41\x2d\x9b\x98\xef\xf7\x74\xb9\x84\x8c\x97\xa5\xaa\x77\xfb\x70\xb9\xfc\x10\x0a\xce\x48\x0e\x08\xa1\xaa\x70\x9d\x93\x1c\x5c\x87\xce\xfe\x03\x21\xcf\xbb\x2a\x62\x57\xc9\x57\x00\x53\xce\x59\x32\x5b\x73\xc8\x15\xd8\xcb\x0d\x7a\xa3\x7e\xa9\xda\xaf\x44\x7e\x83\x24\x4b\xe1\x0d\x3c\xa1\x99\xf1\x0b\xd2\x51\x8e\xa4\x5f\x61\x0f\xbd\x36\xc2\xc9\x01\x22\xf1\xa4\x27\x2b\x63\x90\x40\x50\x0c\x34\x74\x1d\x41\x5d\x24\x7f\xcd\x41\x39\x69\x27\x51\xaf\x59\xbf\x05\x0f\x4d\xbc\xc2\xf3\x6a\xb6\x87\x39\xdd\x02\x7f\x4f\x33\x0e\x8f\x1c\x87\xfc\x11\x89\x2e\x24\xd0\x03\xfd\x6c\x9a\x48\xa3\x79\x7d\xdc\x40\xc6\x3f\x5f\x63\x9a\xa1\x20\x08\x66\x94\xa6\xfd\x2c\x6a\x58\xe5\x53\xc7\x70\xb8\xb8\xc0\x34\x8e\xc7\xf2\x50\xd0\xc7\x70\x31\xba\x7b\x0d\x73\xb5\x97\xba\x82\x1b\xb3\x04\xd9\xe5\x07\x1c\xca\x3f\x22\xab\x0c\xd9\xbe\x86\x3e\x66\xa5\xc5\x4f\x86\x33\xf1\x3f\xd2\x79\x4e\xd7\x1b\xf3\xf7\x00\x57\x0b\x79\x34\x4f\x85\xe3\xa9\x56\xae\xc9\xab\x6d\x2e\x0d\x3c\xa8\xc5\x17\x51\x0d\xb0\xac\x09\x75\x78\xc8\xc1\x83\xf2\x5b\x68\xa3\xe5\x57\x38\xcb\xad\x52\x20\x9a\x05\x37\x90\xaf\x53\xee\xa1\xea\xa7\x8f\x08\x9b\xe7\xc2\xd7\x1a\x39\xb9\x29\x97\x10\x24\x89\x8d\x9e\x8a\xa8\x87\xce\xea\x46\xb4\x32\xc4\x35\x6c\xf5\xb4\x10\xf9\x53\x12\x45\x29\x6c\x09\x03\xbc\xdc\xca\x91\x29\x9b\xe7\x58\x70\x54\xfe\x5a\xee\xdb\xf0\x08\x64\x93\xe7\xc2\x35\x63\x90\xf1\x51\xa9\xee\x1a\xb6\x38\xaf\x46\xaa\x06\x46\xfb\xad\x32\x42\xc7\x22\x24\x31\x12\x65\x5b\x3b\x2c\xfa\x1b\x7a\x6b\x6b\x6d\xc0\xaf\x61\x2b\x31\x2c\x0e\xbe\x42\xb9\x7b\x7b\xef\xc9\xdc\xd6\xcc\xa1\x0d\x03\x1c\xa2\x52\xc7\xa3\x88\x1b\xaf\x9b\xd6\x41\x47\xf8\xac\x2a\x0d\xf6\x3a\x6a\x4d\xb5\xba\xa6\x1c\xb5\x17\x56\xec\x26\x1a\x8b\xa5\x41\xef\x4e\xef\x0f\xf3\xfd\x77\xcd\xb8\xb7\xbd\xd0\x44\x3b\xe7\x31\x09\x2a\x0a\x03\x2a\x5e\xc3\xb6\x66\x66\xac\x29\xf7\x63\xcc\xe2\xf2\xaa\x93\x4d\x51\x1e\xa6\x7d\x99\x25\xbc\x26\xfe\xba\xdb\x86\x96\x01\x7b\x76\x1f\x46\x86\x0a\x72\x50\x23\x13\x6b\xed\x78\x6c\xbb\x8d\x82\x35\x03\x1f\x20\x26\xeb\x94\x5f\xa8\x4f\x19\x3f\x97\x59\x04\x8f\xb8\xe5\x4c\xb7\xc0\x95\x13\xed\x79\x94\x4a\xd5\x69\x0a\x32\xac\x30\x91\x09\xee\xbb\x9e\x50\x3e\x4a\x86\xb2\xdd\x03\x65\xfc\x3b\xee\xc8\x99\x93\xa2\x08\xbe\x92\x59\xaa\xcf\x3c\x26\x03\x74\xe4\x57\x2f\x21\x1b\x78\x80\x92\xd6\xc1\xa2\x53\x27\x36\x93\xed\x5f\xb4\xfc\xff\x5f\x09\x7f\xf8\xc2\x20\x4e\x2a\xf3\x69\x58\x4c\x02\xa3\x63\x33\x9c\x2d\x8c\xaa\x03\x32\xea\xb2\x5f\x80\xdb\x88\x98\x0c\xf9\xf6\xfb\x2f\x53\xd1\xff\xe2\x9c\xae\x59\x08\xfb\xc7\x1f\x1d\x45\xc3\xea\x70\x14\x4e\x5d\xcc\x3d\x31\x2f\x12\x88\x9e\xa8\x93\x11\xa9\xcb\xa4\x9e\xd3\x5a\x8e\xae\x37\x42\xb1\x27\x54\x1b\x0c\x8c\xa9\xd3\x10\xaf\x6a\x96\x4f\xcf\xf4\xb2\xa8\x6d\x85\xdd\xf2\x08\x3b\x62\xb5\x6f\x10\x88\x02\xd0\x14\x3a\xc5\x40\xd4\x07\xa1\xcb\x0d\x84\x1b\x4c\xbc\xe0\x73\x06\x72\x0b\xa9\x88\xa0\x33\xa4\xe8\x6a\xa5\xe4\xb2\xfd\x77\x0d\x6c\xa7\x82\xeb\x74\x90\xa0\x6b\x4e\x66\xce\xd0\x87\xf3\xcb\xe0\x22\x61\x70\x03\x24\x4a\xb2\x39\x26\x3e\xaa\x49\xb5\x4f\x6c\x2c\x8e\x0a\xbf\x06\x36\x52\xee\x09\x69\x38\xd5\x34\xda\x8c\x21\xda\xe3\x6b\x79\xe2\xe1\x05\x93\x5b\x41\x06\xe1\xc6\xde\x0a\xfa\xe8\xf0\x1a\xae\xc8\x1c\x7c\x94\x27\xbf\x81\xc0\xea\x5b\x53\x49\xc0\x13\x10\x7f\xfd\x8b\x6f\xad\x6e\x12\x23\xc9\xb0\x51\x65\xc4\x37\x22\x41\x23\xcf\x4a\x25\x0e\xac\x7d\x55\x8c\xba\xd7\xea\x0b\x99\x03\x16\xa2\xca\xaf\xdb\xe4\x37\xc0\x42\xe4\x7a\x1d\x05\x57\x2f\x90\x16\xf0\x8e\xf0\x81\x71\x74\x5d\xeb\xf0\xee\x49\x3e\x22\x3e\x7d\x75\xb2\x57\xba\x4e\x28\xec\xac\xc9\x59\x2e\x63\x84\x6f\x3b\x49\xbe\x4d\x78\xf8\xa0\x8f\x14\x21\xdc\x04\x58\xec\x50\x95\xe1\x42\xe1\x5f\x1d\x45\xec\x54\x6c\x9e\x87\xbd\xcb\xef\x29\xd5\xaf\x37\x72\xfb\x2d\xa9\x3f\x3b\x71\x8b\x76\xb3\xb6\x8f\xa6\xbb\xf1\xf6\xcf\x48\x8d\x51\x07\x72\x5a\xf7\x19\xa9\x8f\x92\x6c\xb5\xe6\x37\x74\x2b\x7d\xbf\xe3\x3c\x64\xf8\x10\x75\x43\x18\x62\x82\x40\xa7\xce\x55\xdf\x59\x31\xb2\x5a\x4f\x81\x75\xd6\x63\xab\x0a\x5e\xf4\x9d\xae\x53\x22\x48\x73\x18\x46\x6b\xb4\x5a\xf6\xb9\x12\xdd\xe6\xc1\xbe\x15\x06\x0a\x56\xe7\xa1\xe8\x71\x26\xfb\x83\x1b\xac\x69\x81\x01\x73\x8d\x3c\x57\x1d\x61\xc0\xbd\x53\xc8\x3f\x90\xc1\x1a\x36\xb0\x95\x1f\xb0\x9d\xc8\x81\xe7\xbb\xcf\x71\x9c\xc3\xf1\xa5\x8c\x4a\xb4\x9f\xbf\x98\x69\xf5\x94\xb8\xcf\x5c\xd0\xc6\xd3\xfe\xb3\xa8\xfd\xbf\x14\xb5\x69\x14\x61\x0f\xe1\xd5\xa2\x19\x2c\xcd\x3e\xbc\x28\x82\x73\x88\x29\x83\xcb\x2c\x07\xc6\x75\xbe\x38\xd8\x89\x1b\x69\xf1\x24\x64\x40\x78\x92\xcd\x27\x3e\x22\xbe\x58\x41\xaf\xbe\xd0\x6c\x5f\x61\x56\x5a\xac\x16\x3e\x6a\x36\xf5\x02\x55\x3a\x25\x64\x91\x68\xe6\x95\x30\xc2\x41\x8a\x22\x98\xc6\x1c\x58\x43\x3c\xcb\x65\x5e\xbd\x3a\x4a\x58\x88\x6c\x59\x47\x36\xd0\x1f\xa3\xe4\x39\xb6\x3c\x95\xa9\x7f\x5d\x45\x84\x9b\xd4\xfc\xc4\x64\x61\x2c\xa5\x68\xd5\xb9\x67\x5f\xeb\xb5\x98\xaf\x96\xc8\xa6\x72\x68\xdf\x52\x11\x1a\xcb\xfe\x10\x31\xad\xc9\x9e\x4c\x66\x25\x6c\xca\x83\xc7\xaa\xb2\x56\x0c\x2c\x45\xdc\xae\xaf\xcd\x00\x18\xb3\x56\x95\xf5\x0d\xd3\x5c\x9a\xa0\xfb\x46\x4a\xde\x85\x09\xa0\x53\xc5\xce\x77\x9d\xf2\x58\xa5\xf2\x01\xad\x16\x9b\x1c\x78\x0f\xff\xd1\xee\x77\x0b\x5c\x2d\x58\x34\xe5\xcf\xe0\x7e\x52\xa6\x3d\x17\x94\x55\xe8\x35\x71\x9d\x65\x75\x9b\xa7\xc0\x54\x9f\x02\x51\xc2\xdf\xd3\x74\xbd\xcc\x72\x74\x77\xaf\x54\x51\xd7\xa4\xa1\x1c\xad\xaf\x11\x95\xc2\x32\x90\x2d\x9c\xea\x3a\xcf\x1a\xf4\x35\x6e\x4f\x0c\x68\xf1\xec\x48\x78\xb5\xf4\x51\x83\xc0\xf3\xc6\xc5\x9e\x5d\x8e\x88\x8d\x86\xac\x32\x42\x0e\x8a\x3a\x70\x41\xb7\x92\x49\xf4\x29\x89\xeb\x70\xcd\x68\xe7\xef\xfe\xf4\x20\x45\xd0\x9d\x95\x7e\x0b\x83\x3a\x12\xe1\x21\x57\x6c\x1c\x50\xdb\x77\xbc\x63\xb2\x9c\x0c\xc5\x3e\xfe\x75\x51\xf9\x41\xfe\x5d\x85\xb0\xfc\xa1\xea\x95\xc4\x68\xb5\xb0\xfb\xa2\x56\x3d\xeb\xcf\xa2\x8e\xdd\xac\x8f\xa9\x83\xb2\x2e\xa3\xca\x1b\x0f\xbb\xd5\x07\x48\x81\xc3\x53\xdc\xaa\xb7\x20\x2a\x92\x3f\x90\x91\xb4\x4c\xbd\x55\x30\x12\xf3\xcf\x5c\x05\x0d\xcf\xa3\x4a\x9f\x14\xe4\xf8\xd2\xf7\x9e\xae\xb3\xa7\x45\x72\x6b\x37\x73\xc8\x8c\x8a\xcd\x80\x2c\x1f\x1f\x93\x9c\x0f\x95\xab\x6e\x61\x66\x94\xa6\x63\x65\xd1\x6c\x06\x84\xb9\x01\xb1\xad\xf1\x5a\x33\xca\xbd\xde\xa0\xd6\xfb\x0a\x89\x20\x5f\x55\x90\xa0\x28\x5e\x6e\x82\xab\xb2\x44\x67\x48\xfe\xfc\xd6\x7c\x64\x31\xf2\x3c\x5c\x3d\xdd\xa1\x59\xaa\x76\xb5\x52\xed\x23\x9e\x48\xe9\xdd\x79\x8d\x2f\x2f\x5d\xd5\xfe\xbc\x53\x03\x79\xc5\x30\xcd\x22\x25\x85\xc2\x29\xd5\x93\xb7\x89\x52\xa3\x2c\x27\xf2\xe1\x5b\x60\x6b\x55\xab\x55\xbf\x25\x72\xea\x27\x5a\xaa\x6b\xaa\xaa\xaf\xa5\x8e\xb5\x29\x53\x40\x32\xab\x1c\x27\x9a\xda\xff\xd4\xe2\x9d\xa2\x01\x61\x2d\x69\xcb\x9e\x57\x5d\x5d\x6b\x61\x5a\x0e\x46\xb7\x3d\xbd\x92\x67\xde\x62\x2e\x60\xd7\x7a\xba\x24\xd0\x2c\x75\x17\xb0\xfb\x51\x65\xaf\xa4\xb2\x46\xb9\xda\xd5\xfe\x11\x84\x34\xdb\x00\xe3\xb7\x9c\x88\xc2\xa0\x9c\x62\x9a\xeb\xc9\x88\x70\xf2\x75\xb7\x92\xfb\x47\x21\xa2\xd7\x44\xfa\x98\x45\x3d\x76\x1a\xea\x35\xb1\xd0\xa9\x51\xda\x95\x09\xf6\x43\xb5\x70\x2b\x23\x08\x65\x17\xb0\x6b\xed\xc3\xbb\xed\x2b\x77\xb5\xd2\xbe\x0b\x1f\x6d\xac\xd7\xae\xea\xc9\x98\xa3\x2e\xa0\xb0\x9a\x95\x2f\xb9\x84\xec\x91\xba\xc8\x94\xc8\xa2\x4d\xc4\x12\x74\xb1\x40\xa6\x43\x14\xe3\x8d\x43\x1f\x31\x24\xd1\x93\x18\x2d\xf4\x4b\x51\x29\xa5\x0e\xbd\x77\x48\xaf\x9d\xa0\x72\x86\x16\x92\x93\x55\x15\xd5\xb0\xf5\xce\x70\x01\xbb\x4a\x1c\x13\x94\xca\xf6\xd5\x71\x99\x94\xe1\x4c\x59\xec\xee\xed\x7d\x05\x6d\x2c\x65\x54\x3c\xd2\x5d\xda\xc1\xf1\x43\xfe\xb2\xe9\x77\x16\xcb\x5b\x46\xb8\xcb\x34\x17\xa1\xf4\x13\x25\xb6\xab\x9f\x3b\xb1\x5d\x35\x12\x5b\x8f\xb0\x4f\x4c\x6c\xe7\x84\x87\x0f\xdf\x48\x9a\xc8\xfd\xc1\xa1\xcd\xa0\xf5\xe4\x5c\x84\x86\x84\xb4\xda\x48\x3d\x80\x48\xa0\x96\xb7\xf3\x7d\x4a\x93\x5b\x7d\x8b\xad\xb7\xa1\x03\x35\xb9\x42\x1c\xd8\x89\x5b\x92\xee\xf1\xef\x64\x1d\x5b\x67\xc6\x52\x82\xff\x05\x00\x00\xff\xff\x88\x13\xbe\x72\xb5\x31\x00\x00")

func templateDbschemaGotplBytes() ([]byte, error) {
	return bindataRead(
		_templateDbschemaGotpl,
		"template/dbschema.gotpl",
	)
}

func templateDbschemaGotpl() (*asset, error) {
	bytes, err := templateDbschemaGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/dbschema.gotpl", size: 12725, mode: os.FileMode(420), modTime: time.Unix(1613550514, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDbschema_initGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcd\x8e\x9b\x30\x14\x85\xd7\xf1\x53\x1c\xb1\x02\x35\x85\x27\xa8\x54\x51\x36\xa8\x52\x14\x75\xd3\xb5\x81\x0b\x5c\x15\xec\xd4\x5c\x26\x41\x96\xdf\x7d\x64\x92\x49\x36\x33\x2b\xff\xe8\x3b\xf7\x3b\xb7\x28\xf0\x73\x20\x43\x4e\x0b\x75\xa8\x2c\x8c\x15\x50\xc7\x02\x8d\x9e\x27\x3a\xe2\x3a\x72\x3b\x82\x17\xe8\x55\xec\xac\x85\x5b\x3d\x4d\x1b\x5e\xa1\x66\x83\x8c\xf4\xf1\x61\x5d\xae\xd4\x45\xb7\xff\xf4\x40\xf0\x3e\x3f\xdf\xaf\x27\x3d\x53\x08\x4a\xf1\x7c\xb1\x4e\x90\xaa\x43\x32\xb0\x8c\x6b\x93\xb7\x76\x2e\xae\xd4\xdc\xbe\x8b\xbd\x14\x5d\x53\x4c\xdc\x14\xbd\x6e\xc5\xba\x2d\x51\x99\x52\x6f\xda\xe1\x2f\xcb\x78\x76\xd4\xf3\x0d\x3f\xd0\xaf\xa6\x4d\x45\x37\xd3\x3e\x14\x8b\x38\x36\x43\xf6\x38\xe1\xd5\xc1\x91\xac\xce\x20\x89\xf6\x3d\x14\x42\x82\x6f\x78\x46\x54\x50\xde\x73\x0f\xfa\x8f\xbc\x2a\x7f\xd3\x86\xbc\xa2\x5e\xaf\x93\xec\xaf\x10\x76\x67\x55\xd6\x51\x76\x6f\xf2\x02\x6a\xe5\x3d\x4d\x0b\x7d\x4a\x9d\xe8\x5a\x95\x75\x9a\x45\xc6\x74\x71\xdf\x58\x16\x6c\x58\xd2\xcc\x2b\x00\xd8\xcd\x86\xbe\x32\x47\xe4\xe9\x2c\xeb\x3f\x34\xf0\x22\xe4\xd2\xaa\xac\x8f\x71\xa1\x07\x97\x64\x8f\x61\x77\xcd\xc1\xfb\xbc\x36\x2c\xbf\x6c\x17\x8b\x05\xf5\x1e\x00\x00\xff\xff\x1e\xda\x5f\x0a\xd7\x01\x00\x00")

func templateDbschema_initGotplBytes() ([]byte, error) {
	return bindataRead(
		_templateDbschema_initGotpl,
		"template/dbschema_init.gotpl",
	)
}

func templateDbschema_initGotpl() (*asset, error) {
	bytes, err := templateDbschema_initGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/dbschema_init.gotpl", size: 471, mode: os.FileMode(420), modTime: time.Unix(1613548584, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xcd\x6a\xc3\x30\x0c\xc7\xcf\xd2\x53\x88\x10\x46\x5b\x3a\xf7\x1e\xd8\x65\x3b\xed\x52\x0a\x7b\x02\xd7\xd3\x92\x52\xf2\x41\xa2\xb4\x19\x42\xef\x3e\x62\x77\x74\xcb\xa9\x37\x5b\xfa\x7f\xfc\xec\xce\x87\xb3\x2f\x99\x54\xdd\x21\x1d\xf7\xbe\x66\x33\xc4\x53\xdd\xb5\xbd\xd0\x0a\x21\x53\x75\x1f\xa1\xe2\xda\xdf\x24\x07\x2f\x95\x59\x86\x90\x95\x27\xa9\xc6\xa3\x0b\x6d\xbd\xbb\xf2\x71\x7a\x96\xb6\xdb\x71\xa8\xda\x0c\x41\xb5\xf7\x4d\xc9\x94\x9f\xb7\x94\x5f\xa8\x78\x21\xf7\x1e\x23\x07\xb3\x98\x99\x5f\x62\x86\x2a\x37\x9f\x66\xb8\x46\x00\x00\xfc\x1a\x9b\x40\x7b\xbe\xce\x9d\xd2\x8f\x41\x12\xcf\x2a\xc8\x44\x73\xb2\x7b\x6b\x1b\xe1\x49\xd6\xb4\x59\x48\x48\x11\x7a\x96\xb1\x6f\xe8\x69\xb1\x52\x04\x58\x8c\x8a\x24\xfa\xfb\xac\xb4\x70\x4b\xaf\x6d\x93\xfb\xd5\x0f\xfc\xeb\xa5\xe8\xbe\x4f\xf4\x46\x55\x50\x90\x69\x36\x18\x1a\xa2\x7c\x77\xf1\x67\xff\x51\x0e\xf1\x32\xc3\x6e\x1e\x03\x48\xc2\x7b\x17\x1a\xfe\x04\x00\x00\xff\xff\xe3\x2d\x77\xd3\xb6\x01\x00\x00")

func templateModelGotplBytes() ([]byte, error) {
	return bindataRead(
		_templateModelGotpl,
		"template/model.gotpl",
	)
}

func templateModelGotpl() (*asset, error) {
	bytes, err := templateModelGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/model.gotpl", size: 438, mode: os.FileMode(420), modTime: time.Unix(1613549580, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateModel_baseGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\x8d\xc1\x0a\xc2\x30\x10\x44\xcf\xd9\xaf\x58\x4a\x0f\x0a\x9a\xde\x05\x4f\x9e\xbc\x88\xe0\x17\xa4\x71\x69\x4b\x69\x12\xd2\x6d\xad\x2c\xfb\xef\x62\x7a\x9b\x81\x99\xf7\x92\xf3\xa3\xeb\x08\x45\xec\x73\x8f\x0f\x37\x91\x2a\xc0\x30\xa5\x98\x19\x0f\x60\xaa\x6e\xe0\x7e\x69\xad\x8f\x53\xf3\xa1\x76\x3b\x73\x4c\x0d\xf9\x3e\x56\x60\x44\xb2\x0b\x1d\x61\x3d\x9e\xb0\x5e\xf1\x72\x45\x7b\x2f\xc7\x59\x15\x4c\x25\x52\xaf\xaa\x65\x47\xe1\xad\x0a\x47\x00\xfe\xa6\xe2\x7b\x71\x5e\x3c\xef\x3a\x9c\x4b\x41\x01\xf3\x07\xdb\x5b\x0c\x4c\x1b\x83\xc2\x2f\x00\x00\xff\xff\x13\xe0\x2c\x56\xa1\x00\x00\x00")

func templateModel_baseGotplBytes() ([]byte, error) {
	return bindataRead(
		_templateModel_baseGotpl,
		"template/model_base.gotpl",
	)
}

func templateModel_baseGotpl() (*asset, error) {
	bytes, err := templateModel_baseGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/model_base.gotpl", size: 161, mode: os.FileMode(420), modTime: time.Unix(1613544428, 0)}
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
	"template/dbschema.gotpl":      templateDbschemaGotpl,
	"template/dbschema_init.gotpl": templateDbschema_initGotpl,
	"template/model.gotpl":         templateModelGotpl,
	"template/model_base.gotpl":    templateModel_baseGotpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"dbschema.gotpl":      &bintree{templateDbschemaGotpl, map[string]*bintree{}},
		"dbschema_init.gotpl": &bintree{templateDbschema_initGotpl, map[string]*bintree{}},
		"model.gotpl":         &bintree{templateModelGotpl, map[string]*bintree{}},
		"model_base.gotpl":    &bintree{templateModel_baseGotpl, map[string]*bintree{}},
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
