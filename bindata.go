// Code generated by go-bindata.
// sources:
// base.css
// base.html
// base.js
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

var _baseCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\xfd\x6e\x9b\x30\x10\xff\xbb\x79\x0a\x4b\xd5\xa4\x55\x2a\x88\x84\xb1\x0f\x22\xf5\x19\xf6\x0a\x06\x1b\x62\xc5\xd8\xc8\x38\x09\x6d\xd4\x77\xdf\xd9\x8e\xc1\x40\xe8\x3a\x69\xaa\xaa\xc0\xd9\xbe\xfb\x7d\xdc\x99\x83\x6e\xf8\x33\x2a\x24\x79\x45\xd7\xcd\xc3\x81\xb2\xfa\xa0\x73\xb4\x4d\x92\x2f\xfb\xcd\xc3\x85\x11\x7d\x18\xde\x0a\x5c\x1e\x6b\x25\x4f\x82\x44\xa5\xe4\x52\xe5\xe8\x31\x4d\xd3\xfd\xe6\x7d\x43\xd8\xd9\x9c\x6e\x31\x21\x4c\xd4\x39\x4a\x60\x77\x83\x55\xcd\x84\x7d\x7e\xdf\x60\xb3\x5c\x49\xa1\xa3\x0a\x37\x8c\xbf\xe6\xa8\xc3\xa2\x8b\x3a\xaa\x58\x05\x7b\x35\xed\x75\x44\x68\x29\x15\xd6\x4c\xc2\x21\xa8\x41\x15\x67\x82\xc2\xa2\xaf\x55\x55\x95\x49\xf5\x58\x42\x1a\x25\xb9\x2d\x28\x3b\xe6\x0e\x54\xac\xa7\xc4\x64\x92\x6d\x8e\x52\x83\x56\x39\x26\xf6\xd9\xe6\xc7\x9c\xd5\xb0\xb3\xa4\x42\x53\x15\x92\x33\x3b\x1a\x26\x22\x1f\xc8\x92\xb6\xdf\x8f\x5a\x64\x66\xfd\x2d\x62\x00\xa9\xcf\xd1\xce\x60\x88\x8b\x93\xd6\x52\x18\x08\x61\x6a\x5b\x72\x15\xc4\x44\x4c\xe0\x41\x98\x82\xf3\xf2\x4c\x55\xc5\xe5\x25\x47\xf8\xa4\xa5\xd5\xad\x8f\x86\xd2\xe3\xd6\x6e\x9d\xae\x63\x70\x2b\x15\x4a\xbf\x96\xd1\x46\x6e\x70\x5c\x60\x70\x2e\xbd\x6f\xb4\xaa\x0b\xfc\x35\x79\xb6\x7f\xf1\xcf\xa7\xfd\x1d\xdc\x84\x75\x2d\xc7\xe0\xac\x90\xd6\xb6\x41\xb1\xad\xa7\xe0\x44\xbb\x2e\x94\x88\x0b\xd9\x5f\x47\xd4\x5b\x27\xff\x48\x56\x51\x0e\x6d\x71\xa6\x81\x27\xbb\x50\x51\xf7\x12\x3a\x98\xb8\x14\x26\x14\x74\xb4\x8b\x8d\xd4\xbf\x07\x21\xbf\xcd\xc7\x80\x1a\x86\x57\x4e\x2b\xbd\xd6\x3f\x73\x05\x80\x88\xc0\x0d\x9d\xf6\x25\x2e\x3a\xc9\x4f\xda\x60\x2f\x24\xd0\x6f\x9c\x43\x7f\x1b\xac\x40\x6f\x94\xc4\xd9\x5d\xc1\x2f\x52\x91\xa8\x50\x14\x1f\x73\x64\x7f\x00\x21\xf7\xf1\x8b\xc2\xad\x0f\x9b\x80\x81\xa7\x89\xc1\x06\x69\x34\x2b\x31\xf7\x7c\x1a\x46\x08\xa7\x6b\x24\xe7\x17\xc2\xbf\xd6\x8c\x61\x88\x8f\xa6\xec\xa7\x38\xce\xaa\xdd\x43\x04\x29\x59\x53\x4f\x45\x0e\x1a\x64\x2a\xec\x34\x9f\x3b\xfa\xb2\x38\x1e\x78\x64\xe7\x29\x99\x99\x65\x9a\xc0\x3d\xad\xce\x18\xa4\x2e\x0f\xbf\x71\x4d\xaf\xf3\xd6\x19\xef\xc4\x71\x1c\xdd\xc9\xac\xed\x07\x2b\xc3\x59\x9c\x83\x7e\xe4\x30\x1f\x68\x99\xf8\xfe\x65\x30\x05\xfc\x91\x1c\x81\x25\xb3\x01\xff\xf1\xf4\xd9\x9e\x1f\xa7\x3c\x1d\xa0\xbe\xdc\xbe\x06\xb3\x6a\x6b\x7d\x67\x4e\x7d\xe0\xa8\xeb\x21\xb7\x1e\xcc\xee\xaf\xe1\x1e\xf3\x65\x5c\xe4\xff\xbb\xba\xac\x7d\x23\x14\x16\xbf\x85\x96\x9e\x0c\xb7\xa2\xc6\x05\x0f\x1a\xd4\x5e\x59\x93\x2f\xdb\xda\x84\xec\xb2\xec\xd9\xff\x27\x71\xfa\x14\x8a\xfe\xcd\xc9\x07\x4c\x0c\xc4\x81\x11\xc4\x2c\x27\x13\x1c\xc9\x99\xf6\x17\x82\xaa\xe1\x53\xdc\xb1\x37\x0a\x48\x1c\xf4\x29\xd0\xa8\xa4\x76\xae\xd7\x4d\xb3\xdb\xd0\xfc\x32\x5f\x4e\x9c\x9f\xff\xf9\xc2\x9f\x00\x00\x00\xff\xff\xaa\xe8\xac\x11\x78\x08\x00\x00")

func baseCssBytes() ([]byte, error) {
	return bindataRead(
		_baseCss,
		"base.css",
	)
}

func baseCss() (*asset, error) {
	bytes, err := baseCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "base.css", size: 2168, mode: os.FileMode(420), modTime: time.Unix(1477141944, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _baseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\xcd\x6e\xd3\x40\x10\x3e\x07\x89\x77\x58\x7c\x71\x2a\xb5\x71\xb9\x21\x6a\xe7\x42\x41\xea\x05\x2a\xd1\x1e\x10\x42\xea\xc6\x9e\xd8\x4b\xd7\xbb\x66\x77\x92\x36\x8a\x2c\xf1\x14\xf0\x26\x3c\x10\x4f\xc2\xcc\xc6\x76\x9c\xaa\xa0\x72\x88\x66\xf2\xed\xcc\x7c\xf3\xeb\xf4\xc5\xf9\x87\x37\x57\x9f\x2e\xdf\x8a\x0a\x6b\x3d\x7f\xfe\x2c\x65\x29\xb4\x34\x65\x16\x7d\x95\x51\x40\x40\x16\x2c\x6b\x40\x29\xf2\x4a\x3a\x0f\x98\x45\xd7\x57\xef\x4e\x5e\x85\x77\x54\xa8\x61\x9e\x26\x3b\x49\x80\xcf\x9d\x6a\x90\xb4\xed\x76\xf6\x31\xe8\x6d\x4b\x70\x32\xe0\xa9\xc7\x8d\x06\x81\x9b\x06\xb2\x18\xe1\x1e\x93\xdc\xfb\xb8\x73\xe0\xa7\xce\x9e\xd5\x81\xd9\xc8\x1a\xb2\x68\xad\xe0\xae\xb1\x0e\x23\x91\x5b\x83\x60\x28\x93\x3b\x55\x60\x95\x15\xb0\x56\x39\x9c\x84\x3f\xc7\x42\x19\x85\x4a\xea\x13\x9f\x4b\x0d\xd9\xcb\xd9\xe9\xb1\xa8\x09\xab\x57\xf5\x01\x24\xef\xc7\xd0\x29\x63\x2b\x0f\x2e\x00\x72\x41\xd8\x06\x7c\xa8\x31\xe9\x9b\xb0\xb0\xc5\x86\xe4\x24\x2d\xd4\x5a\xa8\x22\x8b\xf5\xc2\xde\xc7\x22\xa4\x9a\x45\x85\xf2\x8d\x96\x9b\xd7\xc6\x1a\x38\x8b\x84\xb0\x26\xd7\x2a\xbf\xcd\x62\x57\x4f\x8f\xce\xb8\xc2\xc9\xc8\x13\x96\x18\x8b\x5c\x4b\xef\xb3\x38\xaf\x54\x5d\xc6\x7b\x87\xc6\xc1\xfa\xa2\x2e\xa7\xd2\x95\xab\x9a\xca\xf4\x9f\x4f\xbf\xf4\x11\x26\x69\xd3\xbb\x29\x63\xc0\xc5\xf3\xdf\x3f\xbf\xa7\x49\xb3\x0b\x9f\x50\xfc\x43\x22\xa7\xca\xea\xef\x4c\x86\xfa\xff\x1f\x4c\x3f\x7e\xfd\x83\x49\x73\xe8\xce\x93\xd4\xde\x37\xc0\x42\x6a\xcc\x62\x6a\x95\xcb\x59\x0c\xfc\x1e\x6d\xf3\x28\xf9\x10\x7f\xaf\xf4\x3c\x3c\x7b\x67\x75\x3c\x0f\x48\xc7\xb2\x58\x21\x5a\x43\xd8\x4e\xd9\xd7\x7b\x29\x4b\x18\x11\x96\x16\xed\x74\x29\xb5\x07\x66\xe2\x46\xa7\xc9\xce\xe5\x69\xae\xe8\x56\xc1\x93\x1b\xb7\xf7\x0c\x39\x6e\xb7\x6a\x29\x66\xe7\xca\xf9\xb6\xed\x63\x71\xbe\x85\x72\x5d\x76\xfb\x58\x68\xcb\x52\x03\xdb\x86\xd5\x60\xe5\x41\xb4\x71\xc5\x14\x81\x2e\x64\xbb\x75\x74\x99\x30\x50\xb0\x55\x2a\x45\xe5\x60\x99\xc5\x74\x3c\x6d\xcb\x36\x24\xd2\x44\x0e\x29\x81\x29\xda\xb6\x13\x7d\xe0\x21\xd0\x05\x42\xed\xf9\xdc\x26\x21\x77\xf8\x26\x66\x57\x74\x96\xe2\x46\xd5\x54\xfa\x4d\x78\x39\x68\x32\x6d\xfb\x7e\xe6\x1d\x58\xa9\xa2\x00\xee\xfc\x68\xe6\x3b\xec\xe2\x60\xd5\xf8\x56\xa6\x58\x29\x7f\xc4\xa0\xb6\x92\x97\xd3\xf3\xf2\x75\x60\x58\x11\x2a\xe0\xda\x69\x2e\xe5\xc1\x86\x75\x91\xf9\x43\x10\x1f\x94\xdd\x99\x93\xf6\x9e\xde\xc6\xd5\x8f\xb7\x87\x5a\x40\x33\x7f\x5a\x45\x5a\x99\x5b\xe2\x40\xfe\x08\x90\x70\xf4\x2b\x9e\x40\xc9\x46\x09\x5b\x27\x9d\xe7\x63\x49\xf0\x1c\x78\x02\x9d\x42\x43\xdf\x7d\x53\xfe\x04\x00\x00\xff\xff\x5c\x37\x80\x01\x8c\x05\x00\x00")

func baseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_baseHtml,
		"base.html",
	)
}

func baseHtml() (*asset, error) {
	bytes, err := baseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "base.html", size: 1420, mode: os.FileMode(420), modTime: time.Unix(1484920956, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _baseJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x5f\x6f\xdb\x36\x10\x7f\xb6\x3f\x05\x67\x60\x91\x54\xdb\x8a\x5c\xb4\x2f\x4d\xbc\x60\xfd\xb3\xb5\x40\xdb\x15\x6b\x1e\x06\x18\x79\xa0\x25\x5a\x62\x26\x91\x0a\x49\xc7\x31\xd6\x7c\xf7\xdd\x91\x94\x2c\x39\x6a\xd6\x0e\x03\xf6\x24\x52\xbc\x3b\xde\xfd\xee\xee\x77\xbc\xa5\x8a\x64\x64\x49\x36\xb4\xd4\xec\x6c\x8c\xdb\x52\xd2\x8c\x8b\xbc\xff\x53\xb9\xcf\x0d\xfc\xcd\x64\xba\xad\x98\x30\x71\x29\x53\x6a\xb8\x14\xb1\x66\x54\xa5\x85\x93\xd0\x20\x91\xcc\x48\x0a\x9f\xe7\xc9\xd9\x98\x6f\xc2\x1b\xf2\xc3\x92\x4c\x26\x11\xf9\x6b\x3c\xc2\x6d\x5c\x51\x93\x16\xe1\xa9\x5e\x86\xf1\x93\xe8\xe4\xe2\x34\x8a\xac\x56\x4d\x95\x66\xef\x84\x19\x92\x58\x2d\xae\xa2\xb3\xbe\x7a\xda\x51\x4f\x07\xd5\xd3\x23\xf5\xfb\xf1\x66\x2b\x52\xf4\x98\xe4\xd2\xc8\x70\xdd\xb8\xb4\x46\x07\xa6\x9d\xc8\x72\x66\xde\x94\x0c\x97\xfa\xe5\xfe\x55\x49\xb5\xfe\x48\x2b\x16\x06\x6b\x79\x17\x44\x71\xc9\x44\x6e\x20\xdc\x11\x03\x78\x40\x73\xbe\x24\x29\xec\x1a\x7c\x34\x39\x07\x08\xc8\x05\x09\x02\xf2\x82\x04\x17\x7a\x19\x90\x29\xd1\xde\x7d\x04\x23\x08\x22\x90\x84\x0b\xc1\x6f\x84\xa9\x95\x3d\x49\xad\x6c\xda\xd8\x46\x73\xc7\x32\x17\x07\x19\x9b\xad\xa1\x8c\x8c\x47\xa3\x07\xff\x40\xae\x8c\x6b\x05\x81\xa7\xb2\x04\x03\xc1\xe9\x29\xda\x29\xe3\x42\x6a\x63\x17\x35\x35\x85\x80\x38\x61\x73\xd3\x43\x4b\x17\xef\xaa\x3c\xb4\x68\xe1\x95\x75\xf7\xca\x66\xe1\xf1\x8a\x75\xaa\x64\x59\x5e\xca\x9a\x7c\xf9\x72\x90\x5a\xcb\x6c\x7f\x38\xf2\xae\x1b\xcc\x1a\x99\x93\x1d\x17\x99\xdc\xc5\x5c\x08\xa6\xde\x32\x9e\x17\x86\x3c\x21\x49\xfc\xdc\x8b\xad\xad\xd8\x74\x58\x6c\x61\xc5\x00\x58\x83\x30\x25\xe4\xe4\x04\xe5\x61\x65\xbd\xb5\xfa\x94\x7c\x43\x66\x0b\x9e\x65\x4c\x04\x58\x64\xa3\x8d\x54\x21\x77\x65\x5c\xc1\x87\x36\x09\x27\x9c\x9c\x13\x0a\x9f\xe9\x34\x02\xb1\x11\x5d\xf1\xab\x58\x2b\x2c\x3e\xbb\xa4\xa5\x41\x75\xc5\xcc\x56\x41\x0a\xee\x5d\x6d\x91\x9f\x48\x38\x84\x83\x0f\xa1\x8b\xd2\x30\x96\x4e\x30\xc2\x2b\xb1\x35\x3f\xb2\x3b\xf3\x89\xe6\x2c\x8c\x3c\x3e\xdf\x1b\x9f\x85\xbe\xc2\x24\xae\xae\x60\xdb\x44\xdb\x84\x39\x5f\x9c\x01\x8e\x50\xc2\x1c\x22\x9d\xcf\xf1\x5a\x0b\xef\xb9\x0b\x52\x6e\x36\x9a\x81\x03\x0a\x6f\x73\x1b\x4c\x36\xe0\xfe\xd8\xf1\x39\x71\xcd\x36\x1a\xc1\xcd\x71\xbd\xd5\x45\x68\xc5\x7b\xee\x5e\xd2\xdc\x39\xcb\xab\x3c\x88\x56\x89\x6d\x79\x0f\x73\xda\x84\x02\x9e\xda\x73\x3c\xba\x1f\xf7\x72\x45\xef\x60\x81\xf6\xbb\xf9\x82\xbf\x6d\xc6\xe0\xcc\xa6\x8c\x99\x9f\x8d\x51\x7c\xbd\x35\x70\x19\x24\x30\x98\x11\x7f\x94\xf7\x8e\x20\xa1\x41\xd4\x27\x0e\xa5\xb1\x15\xb8\x63\x8e\x63\x4b\x66\x5f\x32\xb0\x15\xc0\x9d\xf3\xc2\x66\xed\x05\x59\x24\xc9\x8f\x67\xe8\xc5\x7c\xc7\x33\x53\xf8\x1f\x41\xdf\x6a\x3f\xaf\x9e\x95\x3c\x0f\x47\xc4\x15\x14\x04\x7c\x60\x66\xa3\xb6\xac\xd3\x45\xc0\x5e\xdf\x4d\x5e\x0d\x5d\x99\x41\xba\xc2\x4a\x56\x70\x2a\xd8\x8e\xfc\xf1\xe1\xfd\x5b\x63\xea\xdf\xd9\xcd\x96\x69\x63\xeb\x4e\xc5\xb2\x66\x22\x9c\xfc\xfa\xe6\x72\x32\x23\xed\x20\x18\xa0\x97\xe6\xa8\x61\x99\x56\xb4\x43\x36\x33\x1b\x8f\xb7\x2b\x14\xa3\xd9\x5e\x1b\x6a\x58\x5a\x50\x91\xb3\xe5\x4e\x71\xc3\x5e\x42\xe7\x58\x01\xcd\x44\x16\x8a\x6d\x59\xf6\x21\x6c\x85\x5a\xfc\x54\x6c\x2d\x7d\x46\x4b\x48\x09\xcf\xb0\x48\x41\x1f\xf6\x5b\x8d\x3f\x9e\x26\x1d\x96\xa8\x7d\xb0\xaf\x7f\xfb\xf0\x09\x27\x89\xb2\x71\xda\x23\x89\x04\x14\xdb\xf9\xf2\x8b\x92\xd5\x67\x48\xb8\xc8\xad\x79\x5d\x4b\xa1\xd9\x25\x64\x6e\x46\x26\x06\x3e\xa7\x85\xa9\xca\x89\xd5\x94\xae\xd9\x15\xab\xe4\x2d\x7b\x55\xf0\x32\x0b\x65\x27\x37\x2f\xf7\xef\xb2\x30\x48\xa5\x30\xd0\xe3\xb6\xca\x0e\x6c\x25\x1f\xcf\x61\xcb\x51\xcb\x64\x56\x2d\x3b\xf4\x74\x5e\x1d\xc8\xa9\x4f\x39\xb4\x86\x74\x65\xce\x0b\xdf\x4f\x52\xb0\x8f\x32\x63\xa1\x85\x1e\x6d\xde\x7b\xd4\x3c\x40\x30\xa7\x9e\x25\xcf\x1a\xd6\xe9\x3d\x08\xba\xfd\x50\x39\xbc\x8f\x18\x0e\x3b\x21\x86\xb8\xd5\xa6\x94\x3b\x6c\x59\xba\x35\x12\x7b\x76\xa0\x4a\x2d\x12\x93\x12\x42\x9b\x44\x5e\x33\xe3\xba\x2e\xe9\x1e\x15\x05\xf8\x19\xf4\x9b\x05\x24\x43\xd9\x4e\xa3\xdd\x30\xfd\x75\xad\xc2\xbd\xbb\x87\x96\xb9\x28\xb9\xb5\x6d\xcd\x60\x78\xbb\xc7\xf8\xa8\xa1\xca\x65\xe2\x88\x25\x3f\xe0\xee\x20\xcf\x0f\xe3\x40\xfa\x59\xf0\x4f\xa8\x78\x56\xee\x0f\x5b\x23\xeb\x90\xd9\xf0\x58\x8c\x9b\x4f\x4a\xd6\x34\xb7\x5d\x13\xf6\x6b\x5e\x40\xc5\x21\x1d\x39\x69\xaf\x38\xf4\x2e\xf8\x5a\x39\x95\x4d\x64\xa8\xc1\xbf\xaa\x31\x08\xc4\x35\x88\x2f\xdc\x84\xe4\x9d\xd1\x71\x8d\x9c\x0b\x1f\x87\x09\x14\x14\x5f\x5d\x7b\x58\xe0\x01\x02\xa4\x8e\x6b\x3f\x0b\x9a\x2d\x9a\x58\x5d\x4f\x17\xed\x08\xb5\xf4\xd7\x1c\xba\xd7\x23\xfe\x1d\xf9\x27\x40\xf3\x90\x08\x93\x99\xd7\x1b\x18\x3b\xc7\x33\xc8\x0d\x93\x96\x4b\xa1\xdc\x3b\x48\xd6\x8a\xdd\xfe\xaf\x48\xb6\x10\x3e\xb5\xd3\x97\x5c\x03\x84\xed\xf0\xfd\x66\x08\xe7\xff\x12\xc2\xf9\x7f\x00\xa1\x91\x79\x5e\xb2\xd7\x5c\xe9\xc3\x4b\x31\x7b\xa4\x37\x83\x0c\x44\x03\xff\xaa\x0f\xb3\xa3\xfe\x84\x67\x91\x6b\xfd\x08\x83\xf6\x4b\x7c\xd6\x3e\xec\xe3\x35\xcc\x94\x3f\x03\xff\x64\x1e\x16\x69\x49\xa4\x75\x86\x66\xd9\x9b\x5b\x58\xbc\xe7\xda\x30\x78\x50\x86\x01\x10\xff\x2b\x60\x63\xfc\x07\x7c\xc7\x32\x18\xe4\x4d\x6c\x47\x14\xf7\x50\xd7\xa1\x09\x1a\xf6\xa5\x1c\x75\x3b\xff\xf8\x4d\xf7\x50\x19\xe9\xb5\xa3\x7a\x3f\x73\x34\x0b\xcb\xbf\x03\x00\x00\xff\xff\xf6\x67\x46\x6e\x9d\x0d\x00\x00")

func baseJsBytes() ([]byte, error) {
	return bindataRead(
		_baseJs,
		"base.js",
	)
}

func baseJs() (*asset, error) {
	bytes, err := baseJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "base.js", size: 3485, mode: os.FileMode(420), modTime: time.Unix(1484920599, 0)}
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
	"base.css": baseCss,
	"base.html": baseHtml,
	"base.js": baseJs,
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
	"base.css": &bintree{baseCss, map[string]*bintree{}},
	"base.html": &bintree{baseHtml, map[string]*bintree{}},
	"base.js": &bintree{baseJs, map[string]*bintree{}},
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

