// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package web

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Templates statically implements the virtual filesystem provided to vfsgen.
var Templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
		},
		"/admin": &vfsgen۰DirInfo{
			name:    "admin",
			modTime: time.Date(2021, 2, 10, 16, 35, 17, 764090695, time.UTC),
		},
		"/admin/dashboard.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "dashboard.tmpl",
			modTime:          time.Date(2021, 2, 10, 18, 31, 30, 325113337, time.UTC),
			uncompressedSize: 393,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x8e\xb1\x6a\xc4\x30\x0c\x86\xf7\x3c\x85\xea\xdd\x09\xe9\xd4\xc1\x67\x28\xd7\x27\x28\x85\xce\x6e\xa4\x5e\x04\x8e\x1c\x1c\xdf\x65\x30\x7e\xf7\x12\x37\x29\x57\xb8\x49\x46\x9f\xff\xff\x53\xce\x80\xf4\xcd\x42\xa0\x12\x27\x4f\x0a\x4a\xc9\x99\xfb\x17\x01\xf5\x8a\x13\xcb\x9b\x5b\xc6\xaf\xe0\x22\x7e\x54\xbc\x51\x20\x41\x28\xa5\xb9\xcb\x0e\x41\x12\x49\xda\xd2\x8d\x41\xbe\xc1\xe0\xdd\xb2\x9c\xd4\xec\x2e\xa4\x47\x72\x48\x51\xd9\x06\xc0\x8c\x3d\x30\x9e\xd4\x4a\x7e\x08\x13\x29\xfb\xd8\xf5\xb9\xe3\x52\x4c\x37\xf6\xb6\x31\x1d\xf2\xcd\xfe\x6b\x8e\x61\xad\x8d\x00\xf7\xdb\x21\x78\x3d\xa1\xee\x9f\x0f\x36\x57\x5d\x0c\x61\x3a\x87\xab\xa4\x43\x38\xfb\x5d\xf9\xfe\x47\xa0\xad\x73\x53\xce\x7b\xf8\xea\x7f\x1f\x39\x47\x27\x17\x82\xf6\xec\x99\x24\x2d\xa5\xd4\x35\x18\xcf\x36\xe7\x76\x8b\x78\x3e\xbe\x92\xe0\xce\x4d\x77\x14\x1c\xf7\xd7\x01\xe6\x49\x6b\xe8\x62\x58\x41\x6b\xdb\xec\x89\x9f\x00\x00\x00\xff\xff\xfa\x5a\x53\xad\x89\x01\x00\x00"),
		},
		"/auth": &vfsgen۰DirInfo{
			name:    "auth",
			modTime: time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
		},
		"/auth/fallback_sign_in.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "fallback_sign_in.tmpl",
			modTime:          time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
			uncompressedSize: 445,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x74\x8f\xcb\x6a\xeb\x30\x10\x86\xf7\x7e\x8a\x39\xb3\x17\xc6\x67\x55\x82\x6d\xe8\xa6\xdb\x16\x1a\xe8\x5a\xb1\x26\xb1\xa8\x2e\x46\x1a\xc5\x2d\x42\xef\x5e\xec\xd8\x90\x42\xba\xd2\xe5\xfb\xe7\x9b\x99\x9c\x41\xd1\x59\x3b\x02\x64\xcd\x86\x10\x4a\xc9\x59\x37\x4f\x0e\xf0\x39\xf1\xf8\x22\x8d\x39\xc9\xe1\xf3\xb8\xc2\x85\x01\x39\x05\xa5\x54\x77\x95\x83\x77\x4c\x8e\x97\xda\xaa\x55\xfa\x0a\x83\x91\x31\x76\x38\xc9\x0b\x89\x91\xa4\xa2\x80\x7d\x05\xd0\x8e\x0d\x68\xd5\xe1\x4c\x66\xf0\x96\xb0\x7f\xd4\xe9\x63\x83\xa5\xb4\xf5\xd8\xf4\x55\x5b\x2b\x7d\xed\x7f\x79\x83\x9f\x57\x1f\xc0\xfd\xef\xe0\x8d\xb0\x4a\x34\xff\x77\x76\xf6\xc1\x82\x25\x1e\xbd\xea\xf0\xed\xf5\xfd\x88\x20\x07\xd6\xde\x75\x39\xa7\x60\x8e\x1e\x50\x26\x1e\x0f\xe7\xad\xf3\x21\xea\x8b\xd3\x0e\x4b\xb9\x09\x56\x89\x76\x53\x62\xe0\xef\x89\x3a\x64\xfa\x62\x04\x27\x2d\x75\x98\xe2\xb6\xd4\x83\xdc\x24\x63\x9c\x7d\x50\x7b\x76\x79\xff\x95\x8d\xe9\x64\x35\xef\x23\xd7\xcb\xcc\xfb\xfd\xb6\xf7\x7a\x40\xfb\x4f\x08\xa8\x83\x9f\x41\x88\xbe\xca\x99\x9c\x2a\xe5\x27\x00\x00\xff\xff\x62\x75\x88\x18\xbd\x01\x00\x00"),
		},
		"/base.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "base.tmpl",
			modTime:          time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
			uncompressedSize: 928,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x93\xc1\x8e\xd3\x30\x10\x86\xef\x7d\x8a\xa9\xc5\x35\xb5\xb8\xad\x22\x27\x12\xb0\x08\xb8\x6c\x11\xdd\x0b\xa7\x95\x13\x4f\x12\x6b\x27\x76\x65\x3b\x29\xab\xc8\xef\x8e\xe2\xa6\x29\x45\x48\xcb\x6d\x4f\x89\xe7\x9f\xf9\xc6\x33\xfa\x2d\xb6\xf7\xfb\x4f\x8f\x3f\xbf\x7f\x86\x2e\xf4\x54\x6e\xc4\xe5\x83\x52\x95\x1b\x00\xb1\xcd\x32\xa8\x6d\x7f\x24\x0c\x08\x95\xf4\x08\x01\xfb\x23\xc9\x80\x90\x65\x29\xa3\xc7\x20\xa1\xee\xa4\xf3\x18\x0a\x36\x84\x26\xbb\x63\xc0\x93\x14\x74\x20\x2c\xa7\xa9\x22\x5b\x3f\x03\x4b\x47\x06\xbb\x18\xbf\xd8\xec\x70\xf8\x08\x3f\xac\xed\xe1\x80\x6e\x44\x37\x4d\x68\x54\x8c\x82\x9f\x6b\x36\x82\x9f\xaf\x20\x2a\xab\x5e\x66\xd8\x4a\xc1\x5f\xc1\xc9\x44\x59\x6a\x36\x73\x2b\xa5\x47\xd0\xaa\x60\x5e\x13\xbd\x64\x46\x8e\x6c\x2e\x9a\xcb\xde\x0d\x1e\x1d\xe4\x05\x68\xff\x44\xb6\x6d\x51\x3d\x69\x13\xe3\xa2\xea\x06\x52\xc2\x12\x00\x10\xc7\x72\xf9\x03\xf8\x8a\x44\xf6\x82\xd8\x3d\xc8\x1e\x63\xdc\x5e\xf2\xf8\x9a\x28\x06\x4a\xbd\xaf\x5d\x53\x94\x74\x29\x24\x74\x0e\x9b\x82\x4d\xd3\xe0\xe8\xd1\x02\x33\x78\xf2\xb9\x1d\xd1\x8d\x1a\x4f\x2c\x46\x56\x4e\x93\x7e\x7f\x67\x80\x3d\xe0\xc9\xef\xaf\x82\xe0\xb2\x14\x9c\xf4\x6b\x40\xa9\x7a\x6d\x72\x25\x7d\x57\x59\xe9\xd4\x0d\xf2\xc3\xac\xdd\xff\x21\xfd\x37\x74\x08\x5d\xde\x48\xa2\x4a\xd6\xcf\x39\xd9\xd6\x0e\xe1\x96\x3c\x84\xee\xa0\x5b\xb3\x4f\xf1\xbf\xb1\x82\x0f\x74\x59\x3f\x92\xc7\xeb\x72\xdf\x74\x53\x37\x43\x79\xdd\x1a\x6d\xf2\xc6\xba\xfe\x9f\x93\x7d\x33\xaf\x0d\x96\xbc\x37\xc7\x94\x1e\xcb\xd5\x84\x35\x49\xef\x0b\x56\x5b\x13\xa4\x36\xe8\x56\x1f\xc2\xe2\xdf\x59\x41\x13\x18\xec\xe0\x6a\xe1\x15\x23\xf8\xd9\xf0\x82\xa7\x97\xf8\x3b\x00\x00\xff\xff\xbf\xbb\xfc\xb1\xa0\x03\x00\x00"),
		},
		"/error.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "error.tmpl",
			modTime:          time.Date(2021, 2, 9, 14, 51, 23, 670842269, time.UTC),
			uncompressedSize: 369,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x5c\xd0\xc1\x4e\xc4\x20\x10\x06\xe0\x7b\x9f\x62\x82\x17\x3d\xb0\x4d\x3d\x1a\xca\x61\xcd\x3e\x81\x4f\x30\x85\x59\x8b\xdb\x42\x33\x60\x4d\x43\x78\x77\xe3\xb2\xcd\xaa\xa7\x49\x7e\xe0\xfb\xc9\xe4\x6c\xe9\xec\x3c\x81\x48\x2e\x4d\x24\x4a\x61\xf2\x96\x18\x24\x9c\x98\x03\x43\xce\x87\xb7\x84\xe9\x33\xbe\x06\x4b\xa5\xe4\x4c\xde\x96\xd2\xdc\xdf\x99\xe0\x13\xf9\x24\x4a\x69\x94\x75\x2b\x98\x09\x63\xec\xc5\x82\xef\x24\x47\x42\x4b\x2c\x74\x03\xa0\xc6\x4e\x57\xf1\xe1\x1f\xa9\xe2\x8c\xd3\xa4\x41\xde\xbb\x4a\x51\x6d\x4d\x55\x3b\x76\xba\x51\xad\x75\xab\xfe\xe3\x73\xf8\xaa\xee\xaf\xcc\x84\x49\xc6\x59\x76\xcf\xd7\x13\x00\xb5\x30\x81\xb3\xbd\x20\xe6\x63\xb0\x9b\xd0\x39\x1f\x4e\xcc\x3f\xfc\xc2\xb4\x5f\xaa\x13\x40\x21\x8c\x4c\xe7\x5e\x7c\xe0\x8a\xd1\xb0\x5b\xd2\xcb\xe8\x62\x0a\xbc\x1d\x06\x34\x97\xc7\x27\xb1\x37\x0d\xc9\xc3\x90\xbc\x5c\xd8\xcd\xc8\x9b\xd0\x47\x34\x17\xd5\xe2\x8d\x6c\xaf\xe6\xfe\xeb\x3a\x6e\x9b\xfb\x0e\x00\x00\xff\xff\x6b\x12\x36\x5b\x71\x01\x00\x00"),
		},
		"/landing": &vfsgen۰DirInfo{
			name:    "landing",
			modTime: time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
		},
		"/landing/about.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "about.tmpl",
			modTime:          time.Date(2021, 2, 9, 14, 51, 23, 670842269, time.UTC),
			uncompressedSize: 155,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x4c\xcc\xb1\xaa\xc3\x30\x0c\x85\xe1\xdd\x4f\x71\xae\x77\x11\xb2\xfb\x1a\xba\x77\xec\x0b\xb8\x91\x1a\x1b\x82\x52\x12\x37\x1d\x84\xde\xbd\x64\x28\x74\x3a\xc3\xf9\xf8\xcd\xc0\xf2\x68\x2a\x88\xbd\xf5\x45\x22\xdc\xaf\x45\xb9\xe9\x0c\xc2\xe5\xbe\xbe\xba\x19\x44\x19\xee\xe1\x07\x4f\xab\x76\xd1\x7e\xf2\x90\xb8\x1d\x98\x96\xb2\xef\xff\xf1\x59\x66\xa1\x2a\x85\x65\x8b\x39\x00\xa9\x8e\xf9\x56\x05\xe5\x2c\xe1\x7c\xd3\x50\xc7\x1c\xd2\xc0\xed\xf8\x0e\xd2\x1f\x11\x86\x6d\x7d\x83\x28\x07\x33\x51\x76\xff\x04\x00\x00\xff\xff\xeb\xd0\x45\x14\x9b\x00\x00\x00"),
		},
		"/landing/index.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "index.tmpl",
			modTime:          time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
			uncompressedSize: 3993,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x57\x4b\x6f\x1c\x37\x12\xbe\xeb\x57\xd4\xce\x39\xa2\xe0\xdb\x1e\x64\x01\xc2\x66\x82\x18\x70\x94\xdd\xc8\xf6\x9e\x39\x64\xa9\xa7\x16\x7c\xb4\xc8\xe2\xe4\x20\xcc\x7f\x5f\x14\x5f\xd3\x8a\x6d\x21\xf0\xc9\x56\x37\x9b\x2c\x7e\xaf\xaa\x79\x79\x01\x8b\x4f\x14\x10\x76\x4c\xec\x70\x07\xe7\xf3\xcb\x0b\xbd\xfb\x67\x80\xdd\x47\x1d\x2c\x85\xe5\x53\x7d\x2e\x8f\x01\x83\x85\xf3\xf9\x6a\xf3\x91\x89\x81\x31\xb0\x7c\x76\x75\x6b\xe9\x04\xc6\xe9\x9c\xdf\xef\x56\xbd\xe0\xf5\x11\xb5\xc5\xb4\xbb\xbb\x02\xb8\x3d\xbe\x03\xb2\xef\x77\x7f\xa2\x33\xd1\xe3\xee\xee\x2f\x87\xfc\xb7\x3f\x3f\x9f\x6f\x6f\x8e\xef\xee\xae\xae\x6e\xc9\x2f\xf5\x0b\x17\x97\xb8\x83\x9c\xcc\xfb\xdd\x8d\xce\x19\x39\xdf\x90\x5f\x6e\x18\x33\x5f\x1f\x31\x79\x42\xb5\x86\x65\x27\x9f\xac\x77\x1f\x63\x42\x0f\xb4\xe6\xe2\xc1\x46\x17\x13\x64\x62\xd0\x1e\xf9\x27\x30\x31\x64\x34\x8c\x5c\x12\x68\x4b\x2b\x65\x43\x61\x01\x74\xc4\x3f\x41\x46\x0b\x36\x02\x52\xc9\x3e\x5a\x60\xf4\x6b\x4c\x40\xc1\x90\x25\x5b\x02\x43\x61\x70\xfa\x10\x13\x02\x72\xdb\x1a\xc1\xeb\x25\x68\xd0\x8e\x9e\x8b\x56\xf0\x88\x16\x9e\x74\x31\x74\x28\x19\xb8\xa4\x95\x32\x50\x00\x2c\xe0\x09\x0e\x74\xc0\x60\x8b\x57\x70\x0f\x3a\x99\x02\xa6\xa4\x5c\x32\x9c\x88\x35\x4a\x69\x4b\x91\xfd\x4a\xa2\x0c\xe9\x18\x83\x29\x19\x34\x06\xd4\x01\x4e\xe8\x6a\x91\x0a\xee\xe5\x24\xe4\xb9\x19\x60\x20\x0f\x4f\xda\x90\xa3\x4c\x19\x96\xa4\x4f\x64\x35\x04\x7c\x2e\x75\xd3\x93\x76\x8e\xb2\x82\x4f\xed\x3a\x31\x19\x92\x7a\x5c\x3c\xc4\xc4\x94\x01\x1d\x7a\x0c\x5c\x3c\x04\x3a\x1c\x81\xd1\xb9\x92\xc1\x47\x87\x99\x09\x15\x3c\x76\xf0\x00\x33\xc3\xea\xb4\xc1\xa4\xb9\x5e\x6a\xc1\xcc\x3a\x43\xfb\xdb\xaf\x98\x2c\xa1\x54\x28\x8b\xdf\x00\x1a\x6c\xa1\x0c\x9c\x28\xb3\x5c\x05\x72\x74\x8e\x0c\x71\xb1\x14\x5a\x0d\x83\x2f\x05\x9f\x1d\x27\x32\x58\x41\x24\x6d\x8a\xa3\x0c\xa1\x04\x23\x5c\x29\xf8\xdd\x52\xdc\x5c\xbd\x43\x37\xbe\x06\xaf\x73\xee\x9c\x9c\x8a\x5b\x0b\x6b\x46\xe1\x61\xb3\xa0\x7e\x60\xa2\xf7\xd1\x46\x78\x96\xba\x36\x17\xf9\x15\x83\x4d\x98\xa8\x73\xdd\xa9\xc6\xa5\x63\xe1\xbe\xab\x33\x05\x0f\x15\xfd\x46\xac\xe8\xa8\x64\x78\x2e\xda\xc3\x8a\xce\x61\x60\xcc\xf2\x3a\xa0\x81\xa0\xbd\x82\xff\xc8\xc1\x89\x44\x0b\xf9\x55\xb1\x51\x2e\x58\x58\xc1\xa7\xaa\x42\xd1\xe0\xe4\xed\x09\xcb\x42\x9a\xe1\x44\x27\xed\x45\x29\x0c\xba\x2c\x05\xd5\xed\xcd\x7a\x27\x3e\x78\x44\x0f\xa1\x38\xa7\x61\x3d\xea\x84\x9c\x34\x58\xd2\xfe\x72\xfd\x40\xd9\x35\x78\xea\x39\x3e\xa6\x03\x35\x10\x06\x22\xf5\x79\x97\xa0\x54\x76\x61\x52\xc1\x87\x00\x31\x05\x9d\xb0\xdd\xec\x44\x27\x4c\x49\x37\x7d\x65\xbd\x10\xb3\xa8\x6e\x2f\x84\x57\xc8\x86\x30\x4d\xf1\x90\xa3\x21\x61\x52\x73\x7c\x2e\xa8\xe0\xbe\x18\x8e\xa9\x6b\xf6\x2b\xd4\x14\xfc\x3b\x26\x66\xe2\xca\x41\xce\x1a\xc8\xf6\xb5\xcd\x77\x1e\x4e\x22\xd6\x43\x71\xc5\xf7\x6b\x1c\x9c\x64\x0a\x0f\x8b\x55\x70\x15\xfc\xde\x0a\x76\x68\xb8\x6c\x84\xf2\x96\xb2\x15\x7c\xe6\xe6\xb1\xb1\xe5\x29\xba\xc2\xab\x16\xf9\xa0\xc1\xa0\xb3\x82\x5f\x30\x75\x0f\x55\xc4\xb0\x4c\x76\xd6\x84\x4c\xdd\x5b\x0a\xee\xb9\xd3\x3c\xd0\xda\x58\x43\x8e\xda\xf3\x3c\xbb\x89\xb1\x8a\x2b\x97\xbc\x62\xb0\x94\x33\x2a\xf8\x4d\xdc\x72\x31\x81\x40\x91\x51\x34\xdb\x77\xa6\x00\xc7\xa1\x5b\x21\xb7\xbe\x1a\xd5\x34\x77\xa1\x85\xb5\xb8\x13\x05\x9d\xa6\x58\x1e\xc6\x7e\x55\x43\xe0\xb4\xd9\x54\x59\x19\x51\xf0\x39\x85\x0a\xfd\x04\xa0\xad\x72\x3a\x26\x14\x39\xc5\x20\x70\xeb\x03\x89\xe9\x07\xdd\x35\xe8\xb4\xa9\x77\xef\xb8\x17\xd9\xc7\x36\x9d\xf5\x78\xda\xfc\xaf\x07\x10\x59\xa0\xc0\x98\x6a\x5e\xfe\xaa\x0d\x1c\x65\x63\x09\x7f\xa1\x8a\x51\xc4\x6c\xb8\xf8\xcc\x15\xaa\x2a\xa3\xbf\x23\x49\x21\x67\x5e\x20\xbe\x0a\x10\x05\xf7\x06\x9e\x50\x8a\xb0\x51\xcc\x89\x7d\xc5\x2b\xd7\x56\x0f\xcd\x1d\x86\x53\x6a\x1e\xed\x67\x94\x4a\xe9\x0b\xa6\x26\x9c\x2e\xeb\xf1\x89\x36\xc0\xd3\xce\x4d\xec\x9d\x41\xb9\x88\x64\x47\x8d\x85\xa9\xed\x8c\x5e\x2a\xe1\x98\xc4\x00\x35\x58\x9f\x8b\x08\xd5\xc2\x1a\x13\xeb\xae\xad\xdf\xaa\x33\xba\x79\xe6\x5a\x4b\xbd\x22\xa7\x0d\x05\xd2\x4d\x57\x27\x74\x23\xa3\x86\xde\x7a\xfb\xe9\x58\xce\xc7\xf5\xb6\x14\xb6\xfd\x4a\xc1\x1f\xbd\x39\x79\x5d\x11\x1d\xbd\xaa\xf2\x3a\x3c\xdc\xa0\xff\x5f\xc9\x1c\x25\xe0\xb4\x9f\x5b\x93\xed\x18\xa2\x6b\x3c\x5d\x3a\xd0\x94\xa5\xa8\x65\x02\x56\x57\x17\x86\x13\x06\x0c\x9a\x37\x1a\x09\xe0\x51\x14\x75\x49\x4c\x2c\x0a\x7e\x2e\x24\x87\x74\x29\xd4\xca\x7b\x90\x56\x2c\x5f\x67\x62\xf3\x77\xa9\x6f\x06\x6b\x0d\xc4\x5a\x7b\xcb\xae\x96\xfc\x4f\xd3\xe5\xbd\x15\x29\xf8\xe3\xad\xd0\xde\x7f\x57\x69\x9b\x56\xa5\xe0\x0b\x3a\x78\x4a\x14\x16\x92\xba\x24\x8b\xa4\x40\x6f\x62\x12\x4d\xd4\xf3\x5b\xc9\xe3\x73\x40\x26\x49\xc5\x87\xba\xae\xea\xa9\xf1\xdd\xea\x5e\x4b\x2a\x79\x4e\x21\xd3\xe1\x7b\x91\xe2\x8c\xbc\x0b\x96\x95\xb5\x1e\x93\xfd\x30\x89\x02\x83\x0e\x53\x23\xac\x93\x79\xe1\xbd\xe6\xe5\xd6\x13\xb5\xad\x6f\xab\xb6\xb4\x04\xca\x99\x7c\x35\xe5\x56\xc9\x97\x9b\x16\xee\x59\x7d\x71\x43\x0b\x9f\x8b\xef\xe7\xec\xf1\xc6\x38\xa1\xe0\xc3\x0f\x44\x92\xa0\x17\x4c\xab\xec\x47\xa6\xa9\x7d\x8a\x23\x25\xb4\x69\xec\xf2\x66\xc0\xb2\x7a\xad\x23\x60\xf9\x7a\x7e\x11\x44\xda\xfd\x5e\xdd\xca\xd1\x01\x53\x14\xdd\xce\xf9\xb1\xf5\xe7\x7d\x16\xb6\x7a\x1b\x6d\x9e\xfc\xae\xbe\x15\x3c\x6e\x78\x6b\x51\x36\x07\x1a\x58\x53\xa4\xb0\x51\xb1\xc3\x58\x4d\x58\x2b\xee\x51\x12\x83\x82\x07\xd1\x58\x0f\xa6\x92\x0d\xad\xc4\x73\x2c\x6c\xba\xed\x93\xc4\xa5\x01\x3e\x50\x76\xb3\xd7\x3d\x95\x6c\xaa\xcd\x5b\xd9\x65\x66\xd7\x58\x30\x04\x35\xa4\xf9\x45\x27\xaa\x33\xa7\xa8\xa1\x45\xa6\x90\x33\x25\x2c\x69\xa0\xdb\x88\xd0\x32\x46\xde\x92\x9d\x79\xf4\xe5\x1b\xba\x1f\xf0\xbc\x6a\x9e\x9f\x79\x3c\x7e\x6b\x04\x50\xf0\xb1\x91\xd1\x12\x60\x88\xe9\x32\xdd\x3d\x7e\x6d\x8e\x81\x0c\x7d\x5f\xed\xd5\x0a\x72\x8e\x0e\x2c\x22\xef\xb9\xfd\xa9\xfd\x4a\x18\xe4\xce\xce\x21\x56\xec\xf2\x14\x7c\x0f\x47\x30\x49\xe7\x19\x91\x23\x7c\xe7\x20\xd0\xc7\x13\x05\xbf\xfc\x7d\x87\x0d\x61\x15\xaa\xbf\x78\x28\x48\x1b\x71\x98\x8b\xb6\x12\x37\x1e\x73\x4d\x97\xd7\x15\xca\x61\x7d\x3a\xab\x6c\x88\xb2\x66\x98\xee\xbf\x3d\xfa\xb4\x49\x46\x76\xaf\x28\x4d\x5b\xa0\xd8\x88\xc2\xc8\x9f\x7e\x52\x1f\xdc\xf7\x6f\xa7\xb2\x82\x9f\x47\xca\x6c\x59\xae\x03\xdc\x05\x65\x05\x1f\xa6\x01\x5a\x8c\xfa\x79\xc3\xaf\x32\xac\x2a\x64\x2a\x68\x42\x8d\x6f\xf4\x75\x05\xff\x9a\x1d\xb9\xb7\xf1\x21\x09\x69\xe2\xee\x1b\x4a\xea\xba\xbf\xb1\x74\x1a\xff\xc0\xed\x3f\xae\xaf\xe1\x26\xc5\x3f\xe1\xfa\xfa\xee\xea\xe5\x05\x83\x3d\x9f\xff\x1f\x00\x00\xff\xff\x6f\x91\xef\xb5\x99\x0f\x00\x00"),
		},
		"/news": &vfsgen۰DirInfo{
			name:    "news",
			modTime: time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
		},
		"/news/overview.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "overview.tmpl",
			modTime:          time.Date(2021, 2, 10, 16, 34, 42, 221227039, time.UTC),
			uncompressedSize: 390,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x54\x90\xbf\x6a\xc3\x30\x10\xc6\xf7\x7b\x8a\xeb\x91\xb1\x4a\x70\xa6\x12\x64\x41\xa1\x4b\x97\xd0\x21\xd0\x59\x58\x97\x58\x20\x4b\xc6\x52\xe3\xc0\xa1\x77\x2f\x0e\x0e\x24\x93\xe0\xf7\xfd\xd1\x27\x89\xa0\xe3\xb3\x8f\x8c\x54\x7c\x09\x4c\x58\xab\x88\x6f\x3e\x22\xd2\x91\xe7\x7c\xba\xc3\x85\x21\x47\x87\xb5\xc2\x53\xa2\x4b\xb1\x70\x2c\x4b\x06\xb4\xf3\x57\xec\x82\xcd\xb9\xa5\xd1\x5e\x58\xf5\x6c\x1d\x4f\x64\x00\x51\xf7\x0d\x7a\xd7\xd2\xcc\xa1\x4b\x03\x93\x79\xbe\xe1\x77\x85\xb5\xea\x5d\xdf\x18\xd0\x3b\xe7\xaf\xe6\xa5\x6f\x4a\x33\xbd\x92\x2e\x05\x35\x38\xd5\xec\xc9\x80\xc8\x64\xe3\x85\x71\xe3\xa3\xe3\xdb\x3b\x6e\x38\xf0\xc0\xb1\xe0\xa1\xc5\xed\x67\x08\x3f\x29\x97\x5c\x2b\x00\x68\x8b\xfd\xc4\xe7\x96\x44\xfe\xa6\x70\x4a\x48\x91\xe7\x7c\x18\x53\x2e\x84\xb4\xf8\xbe\xbf\x68\xed\xa9\x75\x9d\xbe\x37\x22\x8f\xca\xed\xd1\x0e\x7c\x1f\xba\x5f\x86\x5a\x03\xa0\xc7\x67\xfd\xc4\xb7\xb2\xe8\xa3\x01\x10\xe1\xe8\x96\x9f\x59\x1f\x74\x3f\x50\xbf\x29\x85\xbb\x29\xcd\xa8\x94\x79\x78\xfe\x03\x00\x00\xff\xff\x91\x0d\x8b\x0e\x86\x01\x00\x00"),
		},
		"/news/post.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "post.tmpl",
			modTime:          time.Date(2021, 2, 9, 14, 51, 23, 670842269, time.UTC),
			uncompressedSize: 388,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x74\xce\xb1\x6e\x84\x30\x0c\x80\xe1\x9d\xa7\x70\x3d\xb5\x43\x40\x74\xaa\x4e\x21\xd3\x2d\xb7\xd0\x0e\xbc\x40\x44\x7c\x25\x12\x97\xa0\x24\x05\xa4\xc8\xef\x5e\x41\xdb\x6b\x19\x3a\x25\xfa\x65\x7f\x72\xce\x60\xe8\x6a\x1d\x01\x26\x9b\x46\x42\x60\xce\xd9\xd6\x2f\x0e\xb0\xa5\x25\x76\x7b\x64\x3e\x41\xce\xe5\x9b\x8f\xa9\x6c\xf5\x8d\xb6\x19\x20\x67\x80\xb9\xf8\x23\xf4\xde\x25\x72\x69\x33\x0a\x69\xec\x0c\xfd\xa8\x63\x6c\x70\xd2\xef\x24\x06\xd2\x86\x02\xaa\x02\x40\x0e\xb5\x3a\x72\xb2\x1a\x6a\x55\xc8\xca\xd8\x59\x1d\x56\x83\x5f\xf0\x58\x7a\x3f\x8a\x9b\x11\xf5\xf3\xd6\xa7\xbb\xd3\xd1\x9a\x36\x67\xfa\x65\xf6\x07\xe4\x83\x10\x50\x05\xbf\x80\x10\xaa\x90\x1a\x86\x40\xd7\x06\x73\xfe\x08\x63\xe7\x01\x1d\x2d\xf1\xe4\x67\x0a\xb3\xa5\x05\x99\x11\xac\x69\xf0\x1e\xd4\xeb\xf7\x4f\x56\xfa\xdf\xf5\xc9\xc7\x84\x80\xdb\x21\x97\x33\xc2\xa3\x75\x3d\x7c\x9d\x75\x39\x3f\xfd\x90\x8e\xd6\x84\xaa\xa5\x35\x95\x65\xb9\x6b\x39\x93\x33\xcc\x9f\x01\x00\x00\xff\xff\xc9\x34\x6e\xa1\x84\x01\x00\x00"),
		},
		"/testing": &vfsgen۰DirInfo{
			name:    "testing",
			modTime: time.Date(2021, 2, 9, 14, 51, 23, 670842269, time.UTC),
		},
		"/testing/base.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "base.tmpl",
			modTime:          time.Date(2021, 2, 9, 14, 51, 23, 670842269, time.UTC),
			uncompressedSize: 243,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x54\x8f\xb1\xae\x83\x30\x0c\x45\x77\xbe\xc2\x2f\xfb\x23\x6b\x87\xc0\x52\x3a\xb7\x03\x4b\xc7\x90\x18\x11\x35\x04\x09\x0c\x6a\x15\xe5\xdf\x2b\x07\x5a\xa9\x53\x14\x5f\x9f\x73\x65\xf5\xd7\x5c\xcf\xed\xfd\x76\x81\x81\x46\x5f\x17\xea\xf3\xa0\xb6\x75\x01\xa0\x46\x24\x0d\x66\xd0\xf3\x82\x54\x89\x95\xfa\xff\x93\x00\x99\x23\x72\xe4\xb1\x8e\xb1\xf3\x93\x79\x80\xc8\x5f\x01\x65\x4a\x0d\xf6\x7a\xf5\x04\x2d\x4f\x62\xc4\x60\x53\x52\x72\x5f\x2f\x94\xdc\xdd\xaa\x9b\xec\x8b\x3d\x5f\x01\x3e\x69\xd6\x59\x70\x30\x5c\x62\xdd\x06\xc6\xeb\x65\xa9\x84\x99\x02\x69\x17\x70\x16\x8c\x31\x08\x07\xc9\x09\x06\x12\x50\xc2\x0f\x2c\xad\xdb\xb8\x71\xaf\x52\x32\x1f\xf7\x0e\x00\x00\xff\xff\xd6\x0a\x65\x40\xf3\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/admin"].(os.FileInfo),
		fs["/auth"].(os.FileInfo),
		fs["/base.tmpl"].(os.FileInfo),
		fs["/error.tmpl"].(os.FileInfo),
		fs["/landing"].(os.FileInfo),
		fs["/news"].(os.FileInfo),
		fs["/testing"].(os.FileInfo),
	}
	fs["/admin"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/admin/dashboard.tmpl"].(os.FileInfo),
	}
	fs["/auth"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/auth/fallback_sign_in.tmpl"].(os.FileInfo),
	}
	fs["/landing"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/landing/about.tmpl"].(os.FileInfo),
		fs["/landing/index.tmpl"].(os.FileInfo),
	}
	fs["/news"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/news/overview.tmpl"].(os.FileInfo),
		fs["/news/post.tmpl"].(os.FileInfo),
	}
	fs["/testing"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/testing/base.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
