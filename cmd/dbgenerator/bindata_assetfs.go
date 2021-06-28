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

var _templateDbschemaGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1a\x5b\x6f\xdb\xba\xf9\x59\xfa\x15\xac\x51\x14\x54\xa1\x2a\x7d\x18\xf6\x90\xc2\xc3\x9c\xb6\x39\x0b\xb2\xa6\x45\xd2\xd3\x0d\x08\x82\x82\x96\x3e\xd9\x9a\x65\xd1\xa3\x68\x3b\x3e\x82\xfe\xfb\x40\x52\x94\x28\xeb\xea\x34\x38\xe8\x19\x0e\x50\x34\x16\xf9\xdd\xf9\xdd\x78\x39\x3b\x43\x7f\x5f\x40\x02\x8c\x70\x08\xd0\x07\x8a\x12\xca\x11\x04\x11\x47\x7c\x19\xa5\x28\x8c\x62\x70\xd1\x7e\x19\xf9\x4b\x14\xa5\x88\x6c\x39\x5d\x13\x1e\xf9\x24\x8e\x0f\xa8\xc2\x9b\x1f\x10\x5f\x82\x1e\xa0\xcc\xb3\xed\x0d\xf1\x57\x64\x01\x28\xcb\xbc\x2f\xea\xe7\x0d\x59\x43\x9e\xdb\x76\xb4\xde\x50\xc6\x11\xb6\xad\x49\xb8\xe6\x13\xdb\xca\x32\x46\x92\x05\xa0\x97\x2b\x17\xbd\xdc\xa1\xf3\x29\xf2\xae\x24\x4c\x9a\xe7\xb6\x35\xc9\xb2\x97\xbb\x3c\x97\x70\x90\x04\x82\x82\x35\x59\x44\x7c\xb9\x9d\x7b\x3e\x5d\x9f\xed\x61\xfe\xf8\x86\xd3\xcd\x59\x30\x9f\x74\xce\x9c\xc5\xd1\xfc\x2c\x24\x3e\xa7\xec\xd0\x05\x05\xfe\x92\xf6\xcd\x9d\x6d\x08\x23\xeb\x89\xed\xd8\x36\x3f\x6c\x00\xdd\xc5\x91\x0f\xdf\xb3\xcc\xbb\xe3\x6c\xeb\x73\xa5\x1e\xba\x7f\x78\x7d\x34\x64\xdb\xe1\x36\xf1\x11\x4e\xdb\x31\x1c\x74\x2b\xb4\xc7\x61\x82\x04\x1c\x5e\xa3\x42\x4e\xef\x13\x0d\x20\x76\x10\x30\x46\x19\xd2\x7f\x33\xdb\x0a\x29\x43\xdf\x5d\x24\x2d\xa5\x2c\x97\x8a\x61\x2b\x0a\x05\x8c\x18\x0d\x13\xbc\x73\xde\xc9\xaf\x17\x53\x94\x44\xb1\x9c\xb7\x18\xf0\x2d\x4b\xc4\xb0\x6d\x59\xb9\x2d\xfe\x15\x43\x49\x14\xdb\x23\x05\xbd\x25\x7b\x43\xd6\x63\x65\x7f\x26\x71\x7f\x61\x74\xbb\xb9\x38\xe0\x15\x1c\x2e\x23\x88\x03\x94\x72\x16\x25\x0b\x07\xad\xc9\xe6\x5e\xfd\x7e\x68\x2e\x97\xe0\x2d\xa5\xea\x85\xca\xf2\x52\x31\x46\xf7\xc7\xaa\x05\x6b\xb2\x91\x63\x74\xef\xcd\xd2\x4f\x64\x83\x1d\xdb\xb2\x76\x2b\x38\x48\x75\xd7\xdc\xbb\xdb\xb0\x28\xe1\x58\x00\xde\x6b\xf9\x1e\x1c\x65\x95\xef\x2e\x92\x70\xec\x5e\x60\x3c\xbc\x43\x2f\x0e\x85\x41\xd4\x00\x9a\xb6\x78\x99\x10\x48\x58\xc8\x00\x22\x9b\x0d\x24\x01\x2e\x06\xa4\xa0\x8e\x69\x45\x36\xc2\x86\xd7\x70\x18\xb0\xe0\x18\xfb\xfd\x6e\xd6\xab\x94\x67\x74\x7f\xaa\xae\xb3\xf4\xfa\xdb\xb1\xaa\x2e\xda\x91\x78\x0b\x75\xed\x65\x22\xf0\xee\x38\x65\x50\x6a\x6b\x8c\xfd\x0e\xda\xc9\xb9\x4a\xb2\x87\x53\x55\xfd\xca\x48\x92\x86\x94\xad\x31\x97\xbf\x80\xa5\xe6\x7a\x29\x65\xbe\x16\x53\x0e\xba\x7f\x68\x53\x79\x4d\x56\x80\x6b\x53\x6e\x0c\x09\x4e\x1d\x47\xe9\x1f\x05\x8f\xad\x16\x60\xf7\x51\xf0\x58\xac\x91\x36\x81\xd7\x22\x51\xd3\x59\x87\xd4\xba\x64\x74\xfd\xcf\x28\xe5\x38\x20\x9c\xa0\x28\xe1\xc0\x42\xe2\x43\x96\x3b\x1d\xd9\x3a\xb3\x2d\x69\xc5\xd4\x45\x74\x25\xe4\x14\x88\x1e\x6e\x86\x97\x63\x8b\xc8\x7c\x41\x57\x52\x03\x9d\xd2\x04\x6a\xa5\x9d\xc6\xad\xf1\x55\x81\xab\x8c\xf0\xaa\x2d\x66\xc5\xa4\x27\x04\xbf\xa5\x7b\x2c\x29\x7a\xd8\x58\x0a\x93\x98\xf0\x01\x2b\xad\x22\x3b\xd5\x31\xad\x02\x5f\x19\x2a\x95\x56\xab\x43\x29\x1d\x3d\xcf\x73\x6c\xcb\xae\x00\x73\xdb\x3e\x3b\x43\x0d\x9b\xe8\xef\xf7\x74\xbd\x86\x84\xe7\xb9\xaa\x77\xc7\x70\xa9\xfc\x10\x0a\xce\x49\x0a\x08\xa1\xb2\x70\x5d\x90\x14\x6c\x8b\xce\xff\x03\x3e\x4f\xdb\x2a\x62\x5b\xc9\x57\x00\x33\xce\x59\x34\xdf\x72\x48\x15\xd8\xcb\x1d\x7a\xa3\x7e\xa9\xda\xaf\x44\x7e\x83\x24\x4b\xe1\x0d\x3c\xa2\x89\xf6\x0b\xd2\x52\x8e\xa4\x5f\x61\x07\xbd\xd6\xc2\xc9\x01\x22\xf1\xa4\x27\x2b\x63\x10\x4f\x50\xf4\x0a\xe8\x2a\x82\xda\x48\xfe\x9a\x82\x72\xd2\x56\xa2\x4e\xbd\x7e\x0b\x1e\x05\xf1\x12\xcf\xa9\xd8\xf6\x73\xba\x03\xfe\x9e\x26\x1c\x1e\x39\xf6\xf9\x23\x12\x5d\x88\x57\x0c\x74\xb3\xa9\x23\x8d\xe6\xf5\x71\x07\x09\xff\x7c\x83\x69\x82\x3c\xcf\x9b\x53\x1a\x77\xb3\xa8\x60\x95\x4f\x9d\xc2\xe1\xf2\x12\xd3\x30\x1c\xcb\x43\x41\x9f\xc2\x45\xeb\xee\xd4\xcc\xd5\x5c\xea\x12\x6e\xcc\x12\x24\x57\x1f\xb0\x2f\xff\x88\xac\x32\x64\xfb\x0a\xfa\x94\x95\x16\x3f\x19\x4e\xc4\xff\x48\x01\x1e\x35\x82\xba\xfc\xf4\x30\x37\x68\x8c\x66\xad\x70\x1c\xd5\xd1\xb5\xb2\x6c\x1a\xaf\xc0\x19\xd4\xe9\x8b\xa8\x0d\x58\x56\x88\x2a\x58\xe4\x60\xaf\x1a\x06\xda\x68\x35\x14\xce\x7a\xaf\xf4\x08\xe6\xde\x2d\xa4\xdb\x98\x3b\xa8\xfc\xe9\x22\xc2\x16\xa9\xf0\xbc\x5a\x86\xae\xcb\x25\x04\x89\x42\xad\xa7\x22\xea\xa0\x69\xd5\x96\x96\x86\xb8\x81\x7d\x31\x2d\x44\xfe\x14\x05\x41\x0c\x7b\xc2\x00\xaf\xf7\x72\x64\xc6\x16\x29\x16\x1c\x95\xf7\xe6\xc7\x36\x3c\x01\x59\x67\x3d\x7f\xcb\x18\x24\x7c\x54\xe2\xbb\x81\x3d\x4e\xcb\x91\xb2\x9d\x29\xbc\x58\x19\xa1\x65\x11\xa2\x10\x89\x22\x5e\xb8\x2f\xfa\x1b\x7a\x6b\x6a\xad\xc1\x6f\x60\x2f\x31\x0c\x0e\xae\x42\xb9\x7f\xfb\xe0\xc8\x4c\x57\xcf\xa8\x35\x03\xf4\x51\xa9\xa2\x53\x44\x91\xd3\x4e\xab\xd7\x11\x3e\xab\xba\x83\x9d\x96\xca\x53\xae\xae\x2e\x4e\xcd\x85\x15\x7b\x8b\xda\x62\x15\xa0\xf7\xe7\x0f\xfd\x7c\xff\x5d\x31\xee\x6c\x36\x0a\xa2\xad\xf3\x98\x78\x25\x85\x01\x15\x6f\x60\x5f\x31\xd3\xd6\x94\xbb\x33\x66\x70\x79\xd5\xca\x26\xcb\xfb\x69\x5f\x25\x11\xaf\x88\xbf\x6e\xb7\xa1\x61\xc0\x8e\xbd\x88\x96\xa1\x84\x1c\xd4\x48\xc7\x5a\x33\x1e\x9b\x6e\xa3\x60\xf5\xc0\x07\x08\xc9\x36\xe6\x97\xea\x53\xc6\xcf\x55\x12\xc0\x23\x6e\x38\xd3\x1d\x70\xe5\x44\x47\x1e\xa5\x12\x77\x1c\x83\x0c\x2b\x4c\x64\x82\xfb\x5e\x4c\x28\x1f\x25\x43\xd9\x6e\x49\x19\xff\x8e\x5b\x72\xe6\x24\xcb\xbc\xaf\x64\x1e\x17\x27\x20\x93\x01\x3a\xf2\xab\x93\x90\x09\x3c\x40\xa9\xd0\xc1\xa0\x53\x25\x36\x9d\xf4\x5f\x34\xfc\xff\x5f\x11\x5f\x7e\x61\x10\x46\xa5\xf9\x0a\x58\x4c\xea\x51\x6c\x00\x96\x6d\x90\xd6\x92\xfd\x02\x1c\x13\x4f\xdb\x44\xe1\xf6\x97\xed\x2f\x33\xd1\x04\xe3\x94\x6e\x99\x0f\xc7\x67\x20\x2d\xb5\xc2\x68\x73\x14\x4e\x55\xd1\x1d\x31\x2f\xf2\x46\x31\x51\xe5\x20\x52\x15\xc9\x62\xae\x50\x6e\x74\x99\x11\x8a\x3d\xa1\xc8\x60\x60\x4c\x1d\x89\x38\x65\xc7\x7c\x3e\x2d\x56\x43\xed\x2d\xcc\xbe\x47\xd8\x11\xab\xcd\x83\x40\x14\x80\xba\xbe\x29\x06\xa2\x2c\x08\x5d\x6e\xc1\xdf\x61\xe2\x78\x9f\x13\x90\xfb\x48\x45\x04\x4d\x91\xa2\x5b\x28\x25\x97\xed\xbf\x5b\x60\x07\x15\x53\xe7\x83\x04\x6d\x7d\x3c\x33\x45\x1f\x2e\xae\xbc\xcb\x88\xc1\x2d\x90\x20\x4a\x16\x98\xb8\xa8\x22\xd5\x3c\xb6\x31\x38\x2a\xfc\x0a\x58\x4b\x79\x24\xa4\xe6\x54\xd1\x68\x32\x86\xe0\x88\xaf\xe1\x89\xfd\x0b\x26\xf7\x83\x0c\xfc\x9d\xb9\x1f\x74\x51\xff\x1a\x6e\xc8\x02\x5c\x94\x46\xbf\x81\xc0\xea\x5a\x53\x49\xc0\x11\x10\x7f\xfd\x8b\x6b\xac\x6e\x14\x22\xc9\xb0\x56\x5c\xc4\x37\x22\x5e\x2d\xbd\x4a\x25\x7a\xd6\xbe\xac\x41\xed\x6b\xf5\x85\x2c\x00\x0b\x51\xe5\xd7\x5d\xf4\x1b\x60\x21\x72\xb5\x8e\x82\xab\xe3\x49\x0b\x38\x27\xf8\xc0\x38\xba\xb6\x71\x82\xf7\x24\x1f\x11\x9f\xae\x3a\xde\xcb\x6d\xcb\x17\x76\x2e\xc8\x19\x2e\xa3\x85\x6f\x3a\x49\xba\x8f\xb8\xbf\x2c\xce\x15\xc1\xdf\x79\x58\x6c\x53\x95\xe1\x7c\xe1\x5f\x2d\xb5\xeb\x5c\xec\xa0\x87\xbd\xcb\xed\xa8\xd0\xaf\x77\x72\x0f\x2e\xa9\x3f\x3b\x71\x83\x76\xbd\xa4\x8f\xa6\xbb\x73\x8e\x0f\x4a\xb5\x51\x07\x72\x5a\xfb\x41\xa9\x8b\xa2\x64\xb3\xe5\xb7\x74\x2f\x7d\xbf\xe5\x50\x64\xf8\x24\x75\x47\x18\x62\x82\x40\xab\xce\x65\xbb\x59\x32\x32\x3a\x4e\x81\x35\xed\xb0\x55\x09\x2f\xda\x4d\xdb\xca\x11\xc4\x29\x0c\xa3\xd5\x3a\x2c\xf3\x70\x89\xee\x53\xef\xd8\x0a\x03\x05\xab\xf5\x64\xf4\x34\x93\xfd\xc1\x0d\x56\xb7\xc0\x80\xb9\x46\x1e\xae\x8e\x30\xe0\xd1\x51\xe4\x1f\xc8\x60\x35\x1b\x98\xca\x0f\xd8\x4e\xe4\xc0\x8b\xc3\xe7\x30\x4c\xe1\xf4\x52\x46\x25\xda\xcf\x5f\xcc\x0a\xf5\x94\xb8\xcf\x5c\xd0\xc6\xd3\xfe\xb3\xa8\xfd\xbf\x14\xb5\x59\x10\x60\x07\xe1\xcd\xaa\x1e\x2c\xf5\x3e\x3c\xcb\xbc\x0b\x08\x29\x83\xab\x24\x05\xc6\x8b\x7c\xd1\xdb\x89\x6b\x69\xf1\xc4\x67\x40\x78\x94\x2c\x26\x2e\x22\xae\x58\x41\xa7\xba\xd5\x6c\xde\x63\x96\x5a\x6c\x56\x2e\xaa\x37\xf5\x02\x55\x3a\x25\x24\x81\x68\xe6\x95\x30\xc2\x41\xb2\xcc\x9b\x85\x1c\x58\x4d\x3c\xc3\x65\x5e\xbd\x3a\x49\x58\x08\x4c\x59\x47\x36\xd0\x1f\x83\xe8\x39\xb6\x3c\xa5\xa9\x7f\xdd\x04\x84\xeb\xd4\xfc\xc4\x64\xa1\x2d\xa5\x68\x55\xb9\xe7\x58\xeb\xad\x98\x2f\x97\xc8\xa4\xd2\xb7\x6f\x29\x09\x8d\x65\xdf\x47\xac\xd0\xe4\x48\x26\xbd\x12\x26\xe5\xc1\xd3\x54\x59\x2b\x06\x96\x22\x6c\xd6\xd7\x7a\x00\x8c\x59\xab\xd2\xfa\x9a\x69\x2a\x4d\xd0\x7e\x2d\x25\x2f\xc4\x04\xd0\xb9\x62\xe7\xda\x56\x7e\xaa\x52\xe9\x80\x56\xab\x5d\x0a\xbc\x83\xff\x68\xf7\xbb\x03\xae\x16\x2c\x98\xf1\x67\x70\x3f\x29\xd3\x91\x0b\xca\x2a\xf4\x9a\xd8\xd6\xba\xbc\xd2\x53\x60\xaa\x4f\x81\x20\xe2\xef\x69\xbc\x5d\x27\x29\xba\x7f\x50\xaa\xa8\xbb\x52\x5f\x8e\x56\x77\x89\x4a\x61\x19\xc8\x06\x4e\x79\xa7\x67\x0c\xba\x05\x6e\x47\x0c\x14\xe2\x99\x91\xf0\x6a\xed\xa2\x1a\x81\xe7\x8d\x8b\x23\xbb\x9c\x10\x1b\x35\x59\x65\x84\xf4\x8a\x3a\x70\x4b\xb7\x91\x49\xf4\x29\x89\xab\xbf\x66\x34\xf3\x77\x77\x7a\x90\x22\x14\x9d\x55\xf1\x20\x06\xb5\x24\xc2\x3e\x57\xac\x9d\x4b\x9b\x17\xbd\x63\xb2\x9c\x0c\xc5\x2e\xfe\x55\x51\xf9\x41\xfe\x6d\x85\x30\xff\xa1\xea\x15\x85\x68\xb3\x32\xfb\xa2\x46\x3d\xeb\xce\xa2\x96\xd9\xac\x8f\xa9\x83\xb2\x2e\xa3\xd2\x1b\xfb\xdd\xea\x03\xc4\xc0\xe1\x29\x6e\xd5\x59\x10\x15\xc9\x1f\xc8\x48\x85\x4c\x9d\x55\x30\x10\xf3\xcf\x5c\x05\x35\xcf\x93\x4a\x9f\x14\xe4\xf4\xd2\xf7\x9e\x6e\x93\xa7\x45\x72\x63\x37\xd3\x67\x46\xc5\x66\x40\x96\x8f\x8f\x51\xca\x87\xca\x55\xbb\x30\x73\x4a\xe3\xb1\xb2\x14\x6c\x06\x84\xb9\x05\xb1\xad\x71\x1a\x33\xca\xbd\xde\xa0\xc6\x23\x0b\x89\x20\x9f\x56\x10\x2f\xcb\x5e\xee\xbc\xeb\x3c\x47\x53\x24\x7f\x7e\xab\xbf\xb4\x18\x79\x1e\xae\xde\xef\xd0\x24\x56\xbb\x5a\xa9\xf6\x09\xef\xa4\x8a\xdd\x79\x85\x2f\xef\x5a\xd5\xfe\xbc\x55\x03\x79\xc5\x30\x4b\x02\x25\x85\xc2\xc9\xd5\xbb\xb7\x89\x52\x23\xcf\x27\xf2\xf5\x9b\x67\x6a\x55\xa9\x55\x3d\x28\xb2\xaa\x77\x5a\xaa\x6b\x2a\xab\xaf\xa1\x8e\xb1\x29\x53\x40\x32\xab\x9c\x26\x9a\xda\xff\x54\xe2\x9d\xa3\x01\x61\x0d\x69\xf3\x8e\xa7\x5d\x6d\x6b\xa1\x5b\x0e\x46\xf7\x1d\xbd\x92\xa3\x1f\x64\xae\xe0\xd0\x78\xbf\x24\xd0\x0c\x75\x57\x70\xf8\x51\x65\xaf\xa5\xb2\x5a\xb9\xca\xd5\xfe\xe1\xf9\x34\xd9\x01\xe3\x77\x9c\x88\xc2\xa0\x9c\x62\x96\x16\x93\x01\xe1\xe4\xeb\x61\x23\xf7\x8f\x42\x44\xa7\x8e\xf4\x31\x09\x3a\xec\x34\xd4\x6b\x62\xa1\x53\xad\xb4\x2b\x13\x1c\x87\x6a\x66\x97\x46\x10\xca\xae\xe0\xd0\xd8\x87\xb7\xdb\x57\xee\x6a\xa5\x7d\x57\x2e\xda\x19\x4f\x5e\xd5\xbb\x31\x4b\x5d\x40\x61\x35\x2b\x9f\x73\x09\xd9\x03\x75\x7f\x29\x91\x45\x9b\x88\x25\xe8\x6a\x85\x74\x87\x28\xc6\x6b\x87\x3e\x62\x48\xa2\x47\x21\x5a\x15\xcf\x45\xa5\x94\x45\xe8\xbd\x43\xc5\xda\x09\x2a\x53\xb4\x92\x9c\x8c\xaa\xa8\x86\x8d\xc7\x86\x2b\x38\x94\xe2\xe8\xa0\x54\xb6\x2f\x8f\xcb\xa4\x0c\x53\x65\xb1\xfb\xb7\x0f\x25\xb4\xb6\x94\x56\xf1\x44\x77\x69\x06\xc7\x0f\xf9\xcb\xae\xdb\x59\x0c\x6f\x19\xe1\x2e\xb3\x54\x84\xd2\x4f\x94\xd8\xae\x7f\xee\xc4\x76\x5d\x4b\x6c\x1d\xc2\x3e\x31\xb1\x5d\x10\xee\x2f\xbf\x91\x38\x92\xfb\x83\xbe\xcd\xa0\xf1\xee\x5c\x84\x86\x84\x34\xda\xc8\x62\x00\x11\x4f\x2d\xaf\xd3\xec\x53\x84\x72\x5e\x9d\x61\x75\x91\x5d\xec\x44\x07\xca\x72\x89\x38\xb0\x19\x37\x84\x6d\x8a\xd0\xca\x3d\x34\x4e\x8e\x85\x10\xff\x0b\x00\x00\xff\xff\x82\xc2\xf6\xeb\xbf\x31\x00\x00")

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

	info := bindataFileInfo{name: "template/dbschema.gotpl", size: 12735, mode: os.FileMode(420), modTime: time.Unix(1624886885, 0)}
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
