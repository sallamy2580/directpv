// // This file is part of MinIO DirectPV
// // Copyright (c) 2021, 2022 MinIO, Inc.
// //
// // This program is free software: you can redistribute it and/or modify
// // it under the terms of the GNU Affero General Public License as published by
// // the Free Software Foundation, either version 3 of the License, or
// // (at your option) any later version.
// //
// // This program is distributed in the hope that it will be useful,
// // but WITHOUT ANY WARRANTY; without even the implied warranty of
// // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// // GNU Affero General Public License for more details.
// //
// // You should have received a copy of the GNU Affero General Public License
// // along with this program.  If not, see <http://www.gnu.org/licenses/>.

package installer

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

var _config_crd_direct_csi_min_io_directcsidrives_yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x61\x73\x1b\xb7\xd1\xfe\xae\x5f\xb1\xc3\xf7\x9d\xb1\xe4\x92\x27\x4b\x6e\xd3\x84\x33\x1e\x8f\x2b\x47\x1d\x4d\xe2\x58\x63\x29\xe9\x4c\x4d\xb5\x01\xef\x96\x24\xa2\x3b\xe0\x02\xe0\x24\x31\x9d\xfe\xf7\xce\x02\x77\xe4\x91\xbc\x3b\x91\xac\x5d\xc7\xc9\xe2\x13\x79\x00\x16\xc0\x62\x77\xb1\xc0\xf3\x0c\x06\x83\x03\x91\xcb\x1f\xd0\x58\xa9\xd5\x10\x44\x2e\xf1\xc1\xa1\xa2\x7f\x36\xba\xfd\xd2\x46\x52\x1f\xdf\x9d\x1c\xdc\x4a\x95\x0c\xe1\xac\xb0\x4e\x67\xef\xd0\xea\xc2\xc4\xf8\x1a\x27\x52\x49\x27\xb5\x3a\xc8\xd0\x89\x44\x38\x31\x3c\x00\x10\x4a\x69\x27\xe8\xb3\xa5\xbf\x00\xb1\x56\xce\xe8\x34\x45\x33\x98\xa2\x8a\x6e\x8b\x31\x8e\x0b\x99\x26\x68\xbc\xf0\x6a\xe8\xbb\x67\xd1\x57\xd1\xe9\x01\x40\x6c\xd0\x77\xbf\x96\x19\x5a\x27\xb2\x7c\x08\xaa\x48\xd3\x03\x00\x25\x32\x1c\x42\x22\x0d\xc6\x2e\xb6\x32\x31\xf2\x0e\x6d\x14\xfe\x47\xb1\x95\x51\x26\x55\x24\xf5\x81\xcd\x31\xa6\xb1\xa7\x46\x17\x79\xd5\xa1\xde\x20\x88\x2a\xe7\x17\xd6\xf6\xda\x37\x3a\xbb\xba\x78\x4d\x52\x7d\x45\x2a\xad\xfb\xa6\xa1\xf2\x5b\x69\x9d\x6f\x90\xa7\x85\x11\xe9\xc6\x8c\x7c\x9d\x95\x6a\x5a\xa4\xc2\xac\xd7\x1e\x00\xd8\x58\xe7\x38\x84\xb3\xb4\xb0\x0e\xcd\x01\x40\xa9\x03\x3f\x9f\x41\xb9\xca\xbb\x13\x91\xe6\x33\x71\x12\x84\xc5\x33\xcc\x44\x98\x2e\x80\xce\x51\xbd\xba\xbc\xf8\xe1\xf9\xd5\xca\x67\x80\x04\x6d\x6c\x64\xee\xbc\x3e\x57\xe7\x0c\x09\x2a\xed\xd0\x82\x9f\x04\x9c\xbd\x7b\x0d\x7a\xfc\x13\xa9\x65\xd1\x3b\x37\x3a\x47\xe3\x64\xa5\x97\x50\x6a\xd6\x51\xfb\xba\x36\xd6\x13\x9a\x4e\x68\x05\x09\x99\x05\x5a\x70\x33\xac\x16\x86\x49\xb9\x02\xd0\x13\x70\x33\x69\xc1\x60\x6e\xd0\xa2\x0a\x86\xb2\x22\x18\xa8\x91\x50\xd5\xf4\xe0\x0a\x0d\x89\x01\x3b\xd3\x45\x9a\x90\x35\xdd\xa1\x71\x60\x30\xd6\x53\x25\x7f\x59\xc8\xb6\xe0\xb4\x1f\x34\x15\x0e\xcb\x0d\x5a\x16\xa9\x1c\x1a\x25\x52\xb8\x13\x69\x81\x7d\x10\x2a\x81\x4c\xcc\xc1\x20\x8d\x02\x85\xaa\xc9\xf3\x4d\x6c\x04\x6f\xb4\x41\x90\x6a\xa2\x87\x30\x73\x2e\xb7\xc3\xe3\xe3\xa9\x74\x95\x57\xc4\x3a\xcb\x0a\x25\xdd\xfc\xd8\x1b\xb8\x1c\x17\x4e\x1b\x7b\x9c\xe0\x1d\xa6\xc7\x56\x4e\x07\xc2\xc4\x33\xe9\x30\x76\x85\xc1\x63\x91\xcb\x81\x9f\xba\xf2\x9e\x11\x65\xc9\xff\x99\xd2\x8f\xec\x93\x95\xb9\xba\x39\x19\x87\x75\x46\xaa\x69\xad\xc2\x5b\x69\xc7\x0e\x90\xa1\x82\xb4\x20\xca\xae\x61\x15\x4b\x45\xd3\x27\xd2\xce\xbb\xaf\xaf\xae\xa1\x1a\xda\x6f\xc6\xba\xf6\xbd\xde\x97\x1d\xed\x72\x0b\x48\x61\x52\x4d\xd0\x84\x4d\x9c\x18\x9d\x79\x99\xa8\x92\x5c\x4b\xe5\xfc\x9f\x38\x95\xa8\xd6\xd5\x6f\x8b\x71\x26\x1d\xed\xfb\xcf\x05\x5a\x47\x7b\x15\xc1\x99\x0f\x15\x30\x46\x28\xf2\x44\x38\x4c\x22\xb8\x50\x70\x26\x32\x4c\xcf\x84\xc5\x8f\xbe\x01\xa4\x69\x3b\x20\xc5\x6e\xb7\x05\xf5\x28\xb7\xde\x38\x68\xad\x56\x51\xc5\xa0\x96\xfd\x5a\xf5\xce\xab\x1c\xe3\x35\x0f\xa5\xfe\x72\x22\x63\xef\x20\xd1\x8a\xa0\x66\x47\xf5\x43\x54\x52\xdf\xde\x2b\x4c\xd6\x6b\xd7\xa6\x40\x7b\x21\x0d\x26\x1b\xad\xc2\x8a\xc6\x5a\xa7\x28\xd6\x7d\xd3\x4f\xee\x5a\x48\xe5\x36\xa5\x8b\x24\xf1\xc7\x81\x48\x2f\x5b\x67\xd8\xa1\xde\x4e\x75\x52\x29\x8d\x07\x93\x73\x6d\x32\xd1\x30\x81\x95\xe5\xbd\x5b\x6d\xbd\xa6\xde\x49\xf8\x58\x8a\xf4\x46\x46\x1f\x36\x74\xdd\xad\x6f\x2a\x13\x99\xa2\x9d\x5b\x87\x59\x53\xed\x23\xab\x05\x9a\x48\x8c\x5d\x3d\x9b\xf7\x81\x4a\xa6\x0b\xe5\xde\xe6\xb5\xa3\x76\xbd\x48\x87\x59\x4b\xd5\xa3\x13\xab\x1a\x08\x63\xc4\xbc\xb1\xfe\x61\x40\x67\xb9\x51\xe8\xd0\x0e\xe8\xb0\x1c\x94\x3d\x9c\xce\x64\xdc\x36\x61\x1f\x29\xf6\x52\x55\x5e\x98\xe9\x5e\xaa\x6a\xb5\xa9\xca\x05\x56\x85\x0e\xd6\xfc\x68\x2b\x77\x77\xc2\x15\x76\x7b\x87\xf7\xcd\xd7\x6c\xb2\xd5\x08\xdb\x0d\x50\xa4\xa9\x8e\x29\x74\x9e\x89\x5c\xc4\xd2\xcd\x37\xd5\x13\x64\x0e\xe9\x04\xfc\xe2\x8f\x2d\xaa\xa1\xd3\x71\xea\x53\x91\x7a\x89\xb5\x0a\x0e\xdd\x60\x42\xad\x96\xb5\xb2\xe8\xde\x59\x25\xc2\x67\x81\x42\x2a\x5a\xb3\x13\x32\xb5\x34\x2f\xd0\x0a\x41\x50\xa4\x73\x21\x33\x40\x88\x0b\x63\x36\x8f\x8f\xa5\x8e\x71\x91\x42\xbc\xba\xbc\x80\x2a\x15\x8d\x60\x30\x18\xc0\x35\x7d\xb6\xce\x14\xb1\xa3\x93\x90\x16\xa5\x12\x4c\xfc\x48\x61\x47\x1b\xc5\x16\x96\x26\x41\x29\x87\x37\x75\x10\xe1\x1c\x9b\x48\x4c\x13\xc8\x85\x9b\x41\x14\x76\x37\x5a\x2a\x24\x02\x38\xd7\x06\xf0\x41\x64\x79\x8a\xfd\x56\x9b\x84\x73\xad\xcb\xbd\x0e\x13\xfb\x17\x1c\x1f\xc3\xbb\xc5\xd9\xea\x47\xd2\x63\x8b\xe6\x2e\xa4\xcc\x3e\xf9\x81\x89\xd6\x4f\xd6\xcf\xe5\x72\x4f\x82\x7e\x82\x2e\x22\x12\xf6\x8d\xd2\xf7\xaa\x69\x8a\x7e\x7c\x61\x70\x08\xa3\xde\xab\x3b\x21\x53\x31\x4e\x71\xd4\x6b\x9e\xec\xa8\x77\x69\xf4\xd4\xa0\xa5\xbc\x75\xd4\x0b\x09\xd2\xa8\xf7\x1a\xa7\x46\x24\x98\x8c\x7a\x34\xd4\x1f\x72\xe1\xe2\xd9\x1b\x34\x53\xfc\x06\xe7\x2f\xfc\x00\x8b\xcf\x57\xce\x08\x87\xd3\xf9\x8b\x8c\xea\x1b\x07\xa1\xb6\x14\x27\xae\xe7\x39\xbe\xc8\x44\xbe\xf8\xf0\x46\xe4\x0b\x81\x0b\x93\xb1\xf0\xfe\x86\x8e\xdc\xbb\x93\x68\xf1\xad\x51\xec\x8f\x3f\x59\xad\x86\xa3\xde\x72\xed\x7d\x9d\x91\x81\xe6\x6e\x3e\xea\xc1\xca\xec\x86\xa3\x9e\x9f\x5f\xf5\xbd\x5a\xcc\x70\xd4\xa3\xd1\x47\xbd\xc6\x11\x72\xa3\x9d\x1e\x17\x93\xe1\xa8\x37\x9e\x3b\xb4\xfd\x93\xbe\xc1\xbc\x4f\x99\xfa\x8b\xe5\xa8\xa3\xde\x8f\x30\x52\xb4\x28\xed\x66\x68\x82\x05\x59\xf8\x77\x93\xcc\xee\x33\x05\x20\x15\xd6\x5d\x1b\xa1\xac\xac\xae\x42\x6d\x61\x7c\xc5\xe1\x36\xbb\x91\x27\x84\xac\xd8\x3a\x70\xf4\xc1\xbb\x59\xa7\x42\xa9\xb8\x85\x14\xf2\x20\xca\xf4\xc8\x59\x83\x95\x51\xa6\x2d\x94\x5f\x64\x54\x7a\x5d\x48\xce\xc7\x08\xf7\x33\xec\x10\x3a\x43\x28\x54\x82\x26\x9d\x53\x3e\x1a\x2f\xa3\xc3\x4c\xa8\x29\x25\x80\x70\x41\xee\x2d\xbc\x03\x53\x72\x78\x4b\xd6\xdd\xa7\x8e\xed\x52\x0b\x5b\x25\xb7\x7e\x7d\x34\x03\xff\x8f\x22\x44\xf0\xe2\x52\xbc\xcf\x8f\xe3\x18\x73\x47\xae\xb0\x79\xd0\x87\x52\x05\x4c\x4a\x49\x07\x24\x71\xdf\xf3\x33\x43\x6b\x45\xdb\x89\xb5\xb6\x71\x65\xdb\x90\xc1\xcf\x8a\x4c\x28\x30\x28\x12\x9a\xe7\xb2\x4e\x25\x3e\x1f\x6c\x19\x2e\xc8\x0c\xc1\x55\x8c\x75\x11\xc2\xd8\x72\x1f\xcb\xad\xa2\x24\x7e\x8c\x14\xee\xbc\x83\x94\x0b\x68\x53\x46\x26\x1e\xbe\x45\x35\x75\xb3\x21\x3c\x3f\xfd\xf3\x17\x5f\xee\xab\x8b\x10\xe3\x30\xf9\x2b\x2a\x34\x3e\xd4\x6d\xa5\x96\xcd\x6e\xb5\x8b\x89\x5f\x5f\x54\x65\xe5\xd1\x74\xd1\xa6\xc3\xfe\xca\xe0\xbe\xb4\xbc\x7b\x61\xc1\xa2\x83\xb1\xb0\x98\x40\x91\x93\x9e\x28\xb4\x4b\x65\x9d\x50\x31\xf6\x41\x4e\x76\x1b\x44\xda\x2a\x4a\xa7\x73\x38\x39\xed\xc3\xb8\xdc\x8a\xcd\x18\xfd\xfe\xe1\x26\xda\x5c\x62\x97\xe4\xaf\xfa\x6b\xf3\x97\x16\x68\xab\xf5\xc4\xdb\x2b\xdc\x4b\x37\xa3\xeb\x9d\x3f\x53\xcb\x0b\x71\xd7\x99\x0a\xab\xe7\x2a\x2e\xd6\xfd\x98\x77\x34\xa7\x13\xa1\x64\x52\xc9\xac\xc8\x86\xf0\xac\xd3\x5c\x9a\xb3\x8e\x50\x0c\x0a\xbb\xa5\x8d\x84\xa6\xcb\x04\x43\x50\x70\x9d\x1a\x91\x51\x2a\x15\x83\x4c\xe8\xca\x37\x91\x68\xb6\x71\x20\x52\x41\x29\x90\xd2\x86\x15\x5d\x3f\xb1\x65\x14\xad\xb9\xd4\xa5\xd1\x49\x11\xa3\x69\x3e\xad\x21\xbc\x66\x54\x77\xb9\xda\xb6\xf9\xbb\xa7\xf7\xc5\xf0\x5e\x02\xf8\x40\x5b\xb6\x78\x7d\xa0\xf3\xb7\x55\x64\x86\x42\x49\x35\xb5\xe5\x14\xe9\x2a\x4e\x61\x2e\x1c\xda\xf7\x33\xf4\xa7\x8f\x7f\x7f\x29\x65\x19\xbf\x0a\x2b\x13\x6c\xba\xef\x55\x45\xc0\xb4\x10\x46\x28\x87\x98\x50\xf0\xa4\x80\x51\xca\xa8\x05\x78\xb1\xbc\xa1\x3f\x12\x3b\x20\x04\x9c\x10\x82\x69\xa9\xe5\x6d\xdf\xc7\x9d\x2d\x02\xce\xc9\xb3\xd3\x0e\x0b\x5b\xb4\x6a\x69\x92\x0b\xe7\xd0\xa8\x21\xfc\xe3\xfd\xab\xc1\xdf\xc5\xe0\x97\x9b\xc3\xf2\xc7\xb3\xc1\x57\xff\xec\x0f\x6f\x9e\xd6\xfe\xde\x1c\xbd\xfc\xff\x7d\x43\x5b\x53\xea\xbf\x2c\x2b\xa6\x5a\x1e\x9f\x55\xae\x5b\x59\x43\xdf\x9f\xad\x7a\x02\xd7\xa6\xc0\x3e\x9c\x8b\xd4\x62\x1f\xbe\x57\xfe\xf0\x6b\x53\x14\xaa\xa2\xe5\xc6\x49\x37\x98\x1e\x89\x6a\x4e\x66\x7c\xb5\x1f\xa3\xbd\xbe\x1c\xfb\xbf\xba\x39\x6e\xa3\x10\x9f\xf1\xe9\x49\x3d\x9e\xd5\x5e\x80\xc0\xc7\x61\xca\x86\xa3\x32\xd3\x8e\x62\x9d\x1d\x2f\x5f\x88\x5a\x0d\x8f\xae\x03\x6f\x84\x9a\xc3\x32\xd8\x86\x7c\x78\xdd\x23\xac\xa3\x6c\x5a\xc4\x46\x5b\xbb\x78\x16\x6b\x77\xe6\x54\xde\x22\x2c\x92\xe9\x10\xda\xc7\x18\x0b\x7f\x87\x30\x63\xe9\x8c\x30\xf3\xda\xc5\x09\x62\xa1\xfc\x03\x97\xc5\x49\x91\xb6\x8a\x3d\xb4\x88\x10\x29\x9d\xe0\xe6\x19\x71\x14\x22\xbe\x18\xcb\x54\xba\x39\xc5\xf4\x04\x63\xad\x26\xa9\xf4\xd7\x9c\xf6\xc3\x22\xcb\xb5\x71\x42\xb9\xe0\xc6\x06\xa7\xf8\x00\xd2\x41\x46\x69\x2f\x5a\x3a\x38\x0e\x13\x65\x4f\x4e\x4e\x9f\x5f\x15\xe3\x44\x67\x42\xaa\xf3\xcc\x1d\x1f\xbd\x3c\xfc\xb9\x10\x29\x45\xcc\xe4\x3b\x91\xe1\x79\xe6\x8e\xb6\x48\x0e\x4e\xbe\x78\xd4\x0f\x0f\xdf\x07\x6f\xbb\x39\x7c\x3f\x28\x7f\x3d\xad\x3e\x1d\xbd\x3c\x1c\x45\x9d\xf5\x47\x4f\x69\x6a\x35\x1f\xbe\x79\x3f\x58\x3a\x70\x74\xf3\xf4\xe8\x65\xad\xee\x68\x4f\x77\x6e\x7e\x11\x08\x65\xd0\x90\x5e\x37\x36\x2b\x13\xb6\xc6\xba\x70\xb8\x34\x56\x85\xad\x6f\xac\xa2\x59\xb7\xbe\xa4\x35\x3e\x96\x75\x3f\xdf\x6c\x3e\xdd\x64\x22\x1f\xdc\xe2\xbc\x21\x8e\xb5\x8c\xde\xf6\xfa\x93\x89\xbc\xe9\xcd\xf0\xaa\x25\x4a\xae\x3e\x92\xb4\xbe\x8d\x94\x6e\xd1\xb2\xc8\xc6\xed\xec\x7a\x98\xeb\xea\x66\x10\x3f\xc6\x6b\x4a\xaa\xa7\x32\x16\xe9\x5f\x52\x1d\xdf\x5e\xc9\x5f\x1a\xe2\xe3\xfe\xb2\x33\x9d\x60\xfa\x5d\x91\x8d\xd1\xec\xb4\xd6\xee\x17\xc4\xd6\x37\x9e\x2d\x1e\x70\xb7\x35\xbb\x8e\x17\xc3\xae\xd7\xc2\x8e\x19\x50\x14\xa5\xb8\xb5\x53\xa7\x5c\x18\xe7\x7d\xfa\xbb\xa6\x43\xb5\x4b\xf5\xb9\x70\xb3\xdd\x86\x9a\xcd\xed\x47\x33\x04\xa3\xb5\xbb\xac\xd6\xb2\xd3\xb4\x2c\x1a\x29\xf6\xb1\x21\xa7\x73\x9d\xea\x69\x83\xaf\x7c\x6c\x3c\xc0\x69\x27\xd2\x0f\xef\xaa\x6d\x8f\xc2\xb4\xd3\x8f\x3f\x05\x6f\xf6\x1e\x2c\x80\xa3\xda\x27\xba\x12\x1c\xb4\x0a\x0a\x37\xc2\x21\x38\x53\x84\xc0\x6b\x9d\x36\x62\x8a\x43\x98\x50\xde\xb6\x02\x13\x8f\xd1\x31\x4a\xcc\x28\x71\x28\x8c\x12\x33\x4a\x5c\x16\x46\x89\x19\x25\x06\x46\x89\x37\x7b\xfe\x0e\x51\xe2\x38\x46\x6b\xaf\x65\x53\x66\xb7\x32\xfc\xab\x45\xc3\xc5\xa0\xa1\x2f\x38\x89\x66\xa7\xdb\x17\x23\xd3\x8c\x4c\x33\x32\xcd\xc8\x34\x23\xd3\x8c\x4c\x03\x23\xd3\x8c\x4c\x33\x32\xcd\xc8\x34\x23\xd3\xc0\xc8\x34\x23\xd3\x8c\x4c\x33\x32\xbd\x59\x18\x99\x06\x46\xa6\x19\x99\xae\xfa\x31\x32\xed\x0b\x23\xd3\x8c\x4c\x57\xe5\x33\x44\xa6\x4f\x43\x23\x46\xa6\x19\x99\x66\x64\x9a\x91\x69\x5f\x18\x99\x66\x64\x1a\x18\x99\xde\xec\xc9\xc8\x74\xeb\xf0\x8c\x4c\x33\x32\xcd\xc8\x34\x23\xd3\x8c\x4c\x33\x32\xcd\xc8\x74\xa7\x5a\x18\x99\x5e\x93\xcc\xc8\x34\x30\x32\xcd\xc8\x34\x23\xd3\x8c\x4c\x33\x32\xcd\xc8\x34\x23\xd3\x8c\x4c\x7f\x68\x64\x7a\xd1\xed\xfb\xef\x2f\x5e\xff\xf6\x41\x6d\xf1\x93\x36\x6d\x80\x64\x4d\xec\xf3\xd3\xdd\xc4\x4a\xf5\x51\xc4\x32\x04\xbf\x28\xff\x73\x08\xbe\xec\xb9\xb3\x5b\x30\x78\xcf\xe0\xfd\x27\x07\xef\x9f\x87\x46\x0c\xde\x33\x78\xcf\xe0\x3d\x83\xf7\xbe\x30\x78\xcf\xe0\x3d\x30\x78\xbf\xd9\x93\xc1\xfb\xd6\xe1\x19\xbc\x67\xf0\x9e\xc1\x7b\x06\xef\x19\xbc\x67\xf0\x9e\xc1\xfb\x4e\xb5\x30\x78\xbf\x26\x99\xc1\x7b\x60\xf0\x9e\xc1\x7b\x06\xef\x19\xbc\x67\xf0\x9e\xc1\x7b\x06\xef\x3f\x07\xf0\x3e\xdb\x19\x63\x4c\x76\x47\xce\x99\x22\x10\x6a\x7f\x57\x14\x01\x61\xdd\xae\x30\x7e\xb2\xb3\xc2\x99\x88\xf0\xeb\x25\x22\x5c\xd3\x39\x7c\xdd\x98\x6e\x6c\xd3\x73\x0f\x22\xc2\xa7\x20\x3f\x94\x3d\x9b\x8e\xa5\xae\xd7\xf6\x5f\x19\x6b\x02\x45\xf2\x56\xa5\x0d\x31\xab\x6b\x0d\x9f\x80\x6b\x61\xef\x45\xfe\xb6\x75\xac\xe6\x69\xfe\xe6\xf8\x19\x00\x05\xde\xa1\x72\xe7\x57\x3b\xdb\x6b\xe8\x78\xe5\xd5\xbf\x53\xc7\x3b\x54\x89\xde\x6d\xaf\xee\xa4\x71\x45\xfb\x30\xcd\x9b\x75\x7f\x2f\x5b\x3d\xa9\x61\x94\xcf\x90\xaa\x12\x76\x9a\xa9\x2a\x4c\x55\x61\xaa\x0a\x53\x55\x42\x61\xaa\x0a\x53\x55\x80\xa9\x2a\x9b\x3d\x99\xaa\xd2\x3a\x3c\x53\x55\x98\xaa\xc2\x54\x15\xa6\xaa\x30\x55\x85\xa9\x2a\x4c\x55\xe9\x54\x0b\x53\x55\xd6\x24\x33\x55\x05\x98\xaa\xc2\x54\x15\xa6\xaa\x30\x55\x85\xa9\x2a\x4c\x55\x61\xaa\x0a\x53\x55\xaa\x2e\x4c\x55\x01\xa6\xaa\xd4\xfb\x30\x55\x65\x4b\x89\x9f\x29\x55\xc5\x3f\x40\xbc\xa1\x01\xed\x85\x9a\xe8\x7d\x1f\x29\xdf\xae\x8a\x59\x5e\x24\xfc\x5a\xfc\x4b\xf0\x1e\x0f\x38\x8f\x03\x02\x8f\x40\x02\x5b\xa4\x59\x8f\xc3\x02\x7b\x00\x03\x8f\x43\x03\xdb\x62\x57\x5b\xad\xe2\x03\x1e\x61\x1d\x0b\x63\x56\xd3\x07\x67\x35\xc5\xf2\x92\x99\x50\x7b\x31\xa1\x6a\x1d\xbf\xd5\x6a\xca\x34\x2a\xa6\x51\x2d\xcb\x6f\x9f\x46\xf5\xa7\xd0\x88\x69\x54\x4c\xa3\x62\x1a\x15\xd3\xa8\x7c\x61\x1a\x15\xd3\xa8\x80\x69\x54\x9b\x3d\x99\x46\xd5\x3a\x3c\xd3\xa8\x98\x46\xc5\x34\x2a\xa6\x51\x31\x8d\x8a\x69\x54\x4c\xa3\xea\x54\x0b\xd3\xa8\xd6\x24\x33\x8d\x0a\x98\x46\xc5\x34\x2a\xa6\x51\x31\x8d\x8a\x69\x54\x4c\xa3\x62\x1a\x15\xd3\xa8\xaa\x2e\x4c\xa3\x02\xa6\x51\xd5\xfb\x30\x8d\x6a\x4b\x89\x4c\xa3\x62\x1a\xd5\xfa\xd4\x99\x46\xb5\xd6\x93\x69\x54\xab\x1d\x98\x46\xc5\x34\xaa\x5a\x25\xd3\xa8\x3e\x09\x8d\x2a\x74\xf8\x9b\x74\xb3\xaf\x1f\x1c\xaa\x4d\x6e\x50\x67\xef\x5f\x35\x09\xcb\x7f\xf9\x4f\x00\x00\x00\xff\xff\x9f\x7e\x96\x04\x84\xcb\x00\x00")

func config_crd_direct_csi_min_io_directcsidrives_yaml() ([]byte, error) {
	return bindata_read(
		_config_crd_direct_csi_min_io_directcsidrives_yaml,
		"config/crd/direct.csi.min.io_directcsidrives.yaml",
	)
}

var _config_crd_direct_csi_min_io_directcsivolumes_yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xdf\x6f\x1b\xb9\xf1\xc0\xdf\xfd\x57\x0c\xf4\xfd\x02\xb1\x53\xed\x3a\x76\xae\xd7\x8b\x80\x20\x08\x9c\xa6\x08\x72\x29\x82\xd8\xcd\x43\x2d\xb7\x37\xda\x1d\x49\x3c\xef\x92\x1b\x92\xab\x58\x57\xf4\x7f\x2f\x86\x5c\xad\x56\xd2\x4a\x51\x82\x5e\xf3\x32\xf3\x24\x91\xdc\x21\x39\x9c\x1f\xe4\x27\x49\x92\x13\xac\xd4\x47\xb2\x4e\x19\x3d\x02\xac\x14\x3d\x78\xd2\xfc\xcf\xa5\xf7\x3f\xb9\x54\x99\xf3\xc5\xc5\xc9\xbd\xd2\xf9\x08\xae\x6a\xe7\x4d\xf9\x81\x9c\xa9\x6d\x46\xaf\x68\xaa\xb4\xf2\xca\xe8\x93\x92\x3c\xe6\xe8\x71\x74\x02\x80\x5a\x1b\x8f\xdc\xec\xf8\x2f\x40\x66\xb4\xb7\xa6\x28\xc8\x26\x33\xd2\xe9\x7d\x3d\xa1\x49\xad\x8a\x9c\x6c\x50\xbe\x9a\x7a\xf1\x24\x7d\x96\x5e\x9e\x00\x64\x96\xc2\xe7\x37\xaa\x24\xe7\xb1\xac\x46\xa0\xeb\xa2\x38\x01\xd0\x58\xd2\x08\x72\x65\x29\xf3\x99\x53\x0b\x53\xd4\x25\xb9\x34\x36\xa4\x99\x53\x69\xa9\x74\xaa\xcc\x89\xab\x28\xe3\xc9\x67\xd6\xd4\xd5\xea\x8b\xee\x80\xa8\xab\x59\x60\xdc\xdc\xab\x30\xe8\xea\xfa\xcd\xc7\xa0\x36\xf4\x14\xca\xf9\xb7\x7d\xbd\x3f\x2b\xe7\xc3\x88\xaa\xa8\x2d\x16\xbb\x8b\x0a\x9d\x4e\xe9\x59\x5d\xa0\xdd\xe9\x3e\x01\x70\x99\xa9\x68\x04\x57\x45\xed\x3c\xd9\x13\x80\xc6\x10\x61\x4d\x49\xb3\xd5\xc5\x05\x16\xd5\x1c\x2f\xa2\xb6\x6c\x4e\x25\xc6\x25\x03\x98\x8a\xf4\xcb\xf7\x6f\x3e\x3e\xbd\xde\x68\x06\xc8\xc9\x65\x56\x55\x3e\x18\x75\x6b\xd9\x90\x93\x36\x9e\x1c\xc4\x65\xc0\xd5\x87\x57\x60\x26\xbf\xb2\x71\xda\xef\x2b\x6b\x2a\xb2\x5e\xad\xac\x13\xa5\xe3\x24\x9d\xd6\xad\xd9\x1e\xf1\x82\xe2\x28\xc8\xd9\x3b\xc8\x81\x9f\xd3\x6a\x6b\x94\x37\x7b\x00\x33\x05\x3f\x57\x0e\x2c\x55\x96\x1c\xe9\xe8\x2f\x1b\x8a\x81\x07\xa1\x5e\x2d\x0f\xae\xc9\xb2\x1a\x70\x73\x53\x17\x39\x3b\xd5\x82\xac\x07\x4b\x99\x99\x69\xf5\x5b\xab\xdb\x81\x37\x61\xd2\x02\x3d\x35\x87\xb4\x16\xa5\x3d\x59\x8d\x05\x2c\xb0\xa8\x69\x08\xa8\x73\x28\x71\x09\x96\x78\x16\xa8\x75\x47\x5f\x18\xe2\x52\x78\x67\x2c\x81\xd2\x53\x33\x82\xb9\xf7\x95\x1b\x9d\x9f\xcf\x94\x5f\x05\x47\x66\xca\xb2\xd6\xca\x2f\xcf\x83\x9f\xab\x49\xed\x8d\x75\xe7\x39\x2d\xa8\x38\x77\x6a\x96\xa0\xcd\xe6\xca\x53\xe6\x6b\x4b\xe7\x58\xa9\x24\x2c\x5d\x87\x00\x49\xcb\xfc\xff\x6c\x13\x4e\xee\xd1\xc6\x5a\xfd\x92\xdd\xc3\x79\xab\xf4\xac\xd3\x11\x7c\xf5\xc0\x09\xb0\xb7\x82\x72\x80\xcd\xa7\x71\x17\x6b\x43\x73\x13\x5b\xe7\xc3\x9f\xaf\x6f\x60\x35\x75\x38\x8c\x6d\xeb\x07\xbb\xaf\x3f\x74\xeb\x23\x60\x83\x29\x3d\x25\x1b\x0f\x71\x6a\x4d\x19\x74\x92\xce\x2b\xa3\xb4\x0f\x7f\xb2\x42\x91\xde\x36\xbf\xab\x27\xa5\xf2\x7c\xee\x9f\x6a\x72\x9e\xcf\x2a\x85\xab\x90\x31\x60\x42\x50\x57\x39\x7a\xca\x53\x78\xa3\xe1\x0a\x4b\x2a\xae\xd0\xd1\xef\x7e\x00\x6c\x69\x97\xb0\x61\x8f\x3b\x82\x6e\xb2\xdb\x1e\x1c\xad\xd6\xe9\x70\x1e\x7d\xed\x0e\x9c\xd8\x56\x84\x5e\x87\xf1\xdb\x71\xca\x9b\xb7\x65\x08\x92\x74\x43\x55\x7f\xb0\xb2\xe0\x02\x55\x81\x93\x82\xae\xb0\xc2\x4c\xf9\xe5\xf6\x00\x80\xa8\x73\xc4\x41\xf1\xe3\x0f\x3b\xbd\x71\x43\x1c\x30\xb3\x90\x9f\xba\x92\x19\x9d\xab\x4e\x8a\xef\x8a\xf2\x54\xf6\x34\x6f\x6d\x7b\x70\xb5\x52\x11\xea\x03\x2a\xcd\x9b\xf6\xa8\x0a\xc7\xeb\x02\xa3\x09\x90\xd3\xb8\x8f\xc9\x82\x20\xab\xad\xdd\xf5\xa8\xb5\x95\xa9\xcd\x2a\x2f\xdf\xbf\x81\x55\x91\x4a\x21\x49\x12\xb8\xe1\x66\xe7\x6d\x9d\x79\x0e\x0e\xde\x94\xce\x29\x0f\x33\xc5\xd4\xdc\xab\xb6\x76\xbc\x08\xce\x42\x68\x2d\x2e\x01\xa3\x6b\x4f\x15\x15\x39\x54\xe8\xe7\x90\xc6\xf3\x4d\xd7\x06\x49\x01\x5e\x1b\x0b\xf4\x80\x65\x55\xd0\xb0\x57\x2f\x9b\x16\x5e\x1b\xd3\x1c\x76\x5c\xd8\xbf\xe0\xfc\x1c\x3e\xb4\xe1\x16\x66\x32\x13\x47\x76\x11\x8b\x69\xc8\x87\x30\x35\xe6\xd1\x76\xa8\x36\x67\x12\xed\x13\x6d\x91\xb2\xb2\xb7\xda\x7c\xd6\x7d\x4b\x0c\xf3\xa3\xa5\x11\x8c\x07\x2f\x57\x7e\x32\x1e\xf4\x2f\x76\x3c\x78\x6f\xcd\xcc\x92\xe3\x6a\x36\x1e\xc4\x9c\x39\x1e\xbc\xa2\x99\xc5\x9c\xf2\xf1\x80\xa7\xfa\x43\x85\x3e\x9b\xbf\x23\x3b\xa3\xb7\xb4\x7c\x1e\x26\x68\x9b\xaf\xbd\x45\x4f\xb3\xe5\xf3\x92\xfb\x7b\x27\xe1\xb1\x5c\x6c\x6f\x96\x15\x3d\x2f\xb1\x6a\x1b\xde\x61\xd5\x2a\x6c\x5d\xc6\xc1\xed\x1d\x47\xe1\xe2\x22\x6d\xdb\x7a\xd5\xfe\xf2\xab\x33\x7a\x34\x1e\xac\xf7\x3e\x34\x25\x3b\x68\xe5\x97\xe3\x01\x6c\xac\x6e\x34\x1e\x84\xf5\xad\xda\x57\x9b\x19\x8d\x07\x3c\xfb\x78\xd0\x3b\x43\x65\x8d\x37\x93\x7a\x3a\x1a\x0f\x26\x4b\x4f\x6e\x78\x31\xb4\x54\x0d\xb9\x7c\x3f\x5f\xcf\x3a\x1e\xfc\x02\x63\xcd\x9b\x32\x7e\x4e\x36\x7a\x90\x83\x7f\xf7\xe9\xdc\x1f\xd2\x51\x0a\x74\xfe\xc6\xa2\x76\x6a\x75\x49\xea\x1f\xb7\x15\x70\xbb\x9f\x71\x24\xc4\x42\xe9\x3c\x78\x6e\x08\x61\x76\xd0\xa0\x2c\xbe\xd5\xc2\x11\xc4\xc9\x9f\x83\x35\x7a\x19\x17\x5f\xd4\x61\x93\x69\x13\x75\xb1\x5e\x4f\x08\x3e\xcf\xe9\x80\xd2\x39\x41\xad\x73\xb2\xc5\x92\x4b\x54\xb6\xce\x0e\x73\xd4\x33\xae\x09\xf0\x86\xc3\x1b\x43\x00\x73\xbd\xb8\x67\xef\x1e\xf2\x87\xfb\xb5\xd6\x6e\x55\xef\xc2\xfe\x78\x05\xe1\x1f\x67\x88\x18\xc5\x8d\xfa\x50\x32\xb3\x8c\x2a\xcf\xa1\x90\xee\x51\xb8\x4a\x98\x5c\xa5\x12\xd6\xb8\x67\xdc\x9e\xc2\xb1\x96\x92\x9c\xc3\xd9\x71\x07\xd7\x8c\x8d\x45\x7d\x5e\x97\xa8\xc1\x12\xe6\xbc\xce\x75\x9f\xce\x55\x86\x7e\xdf\x74\x51\x67\x4c\xae\x38\x31\x75\x4c\x63\xeb\x73\x6c\x8e\x8a\xeb\xfa\x84\x38\xdd\x85\x00\x69\x36\xb0\xcf\x18\x25\x3e\xfc\x4c\x7a\xe6\xe7\x23\x78\x7a\xf9\xa7\x1f\x7f\xfa\x56\x5b\xc4\x1c\x47\xf9\x5f\x48\x93\x0d\xa9\xee\x28\xb3\xec\x7e\xd6\xb9\xab\x84\xfd\xa5\xab\x42\x9d\xce\xda\x31\x07\xfc\xaf\x49\xee\x6b\xcf\xfb\x8c\x0e\x1c\x79\x98\xa0\xa3\x1c\xea\x8a\xed\xc4\xa9\x5d\x69\xe7\x51\x67\x34\x04\x35\xfd\xba\x49\x94\x5b\x65\xe9\x62\x09\x17\x97\x43\x98\x34\x47\xb1\x9b\xa3\x6f\x1f\xee\xd2\xdd\x2d\x1e\xd2\xfc\x6c\xb8\xb5\x7e\xe5\x80\x8f\xda\x4c\x83\xbf\xc2\x67\xe5\xe7\x7c\xe3\x0b\x35\xb5\xb9\x23\x1f\xaa\xa9\xb0\x59\x57\xa9\xdd\xf7\x97\xa2\xa3\xff\x3a\x11\xa5\x54\x5a\x95\x75\x39\x82\x27\x07\xdd\xa5\xff\xd6\x11\xc5\x12\xba\x23\x7d\x24\x0e\x5d\x5f\x30\x90\x93\xeb\xcc\x62\xc9\x57\xa9\x0c\x54\xce\xb7\xc0\xa9\x22\x7b\x4c\x00\xb1\x09\x1a\x85\x7c\x6d\xd8\xb0\xf5\x23\xd7\x64\xd1\x4e\x48\xbd\xb7\x26\xaf\x33\xb2\xfd\xd5\x1a\xe2\x03\x87\x4f\x43\x4d\x55\xd6\x39\xb6\x70\x1d\x0d\xb1\x18\x9f\x50\x40\x0f\x7c\x64\xed\x83\x84\xeb\xef\x5e\x95\x25\xa1\x56\x7a\xe6\x9a\x25\xf2\xed\x9c\xd3\x5c\x2c\xda\x9f\xe7\x14\xaa\x4f\x78\x92\x35\xba\x6c\xd8\x85\x53\x39\x59\xda\xaf\x16\x61\x56\xa3\x45\xed\x89\x72\x4e\x9e\x9c\x30\x1a\x1d\x9d\x04\x8f\xeb\x4b\xfb\x17\x72\x07\xc4\x84\x13\x53\x30\x6f\xb5\x79\x00\x84\xbc\x73\x44\xc2\xb9\x78\x72\x79\xc0\xc3\xda\x51\x7b\x86\x54\xe8\xf9\x15\x38\x82\x7f\xdc\xbe\x4c\xfe\x8e\xc9\x6f\x77\xa7\xcd\x8f\x27\xc9\xb3\x7f\x0e\x47\x77\x8f\x3b\x7f\xef\xce\x5e\xfc\xff\xb7\xa6\xb6\xbe\xcb\xff\x5a\x36\x5c\xb5\x29\x9f\xab\xbb\xee\xca\x1b\x86\xa1\xb6\x9a\x29\xdc\x58\x7e\xae\xbe\xc6\xc2\xd1\x10\xfe\xa6\x43\xf1\xdb\x67\x28\xd2\x75\xb9\x6f\xd2\x04\x06\xac\xaa\xff\x32\x13\xba\xc3\x1c\xfb\xfb\x9b\xb9\xbf\xd5\x24\x61\xc0\x31\x06\x09\x37\x3e\x33\xed\xe6\xb3\xce\xa3\x10\x42\x1e\xe6\xdb\x70\xda\xdc\xb4\xd3\xcc\x94\xe7\xeb\x47\xe3\x5e\xc7\xe3\xe7\xc0\x3b\xd4\x4b\x58\x27\xdb\x78\x1f\xde\x8e\x08\xe7\xf9\x36\x8d\x99\x35\xce\xb5\x2f\xe5\xfd\xc1\x5c\xa8\x7b\x82\xf6\x32\x1d\x53\xfb\x84\x32\x0c\x6f\x08\x3b\x51\xde\xa2\x5d\x76\x1e\x4e\x90\xa1\x0e\x6f\x5e\x47\xd3\xba\xd8\xab\xf6\xd4\x11\x41\xaa\x4d\x4e\xbb\x35\xe2\x2c\x66\x7c\x9c\xa8\x42\xf9\x25\xe7\xf4\x9c\x32\xa3\xa7\x85\x0a\xcf\x9c\xfd\xc5\xa2\xac\x8c\xf5\xa8\x7d\x0c\x63\x4b\x33\x7a\x00\xe5\xa1\xe4\x6b\x2f\x39\x2e\x1c\xa7\xb9\x76\x17\x17\x97\x4f\xaf\xeb\x49\x6e\x4a\x54\xfa\x75\xe9\xcf\xcf\x5e\x9c\x7e\xaa\xb1\xe0\x8c\x99\xff\x15\x4b\x7a\x5d\xfa\xb3\x23\x2e\x07\x17\x3f\x7e\x31\x0e\x4f\x6f\x63\xb4\xdd\x9d\xde\x26\xcd\xaf\xc7\xab\xa6\xb3\x17\xa7\xe3\xf4\x60\xff\xd9\x63\x5e\x5a\x27\x86\xef\x6e\x93\x75\x00\xa7\x77\x8f\xcf\x5e\x74\xfa\xce\xbe\x31\x9c\x2d\x7d\xaa\x95\xa5\xbc\xcf\x7b\x93\x9e\xeb\x75\xef\xb0\xe6\xc2\xd6\xdb\x17\x8b\x4b\x6f\x57\x3c\xfa\xde\x2e\x5e\x75\x4f\xc7\x1e\x1c\xd1\xed\x0c\x6f\xda\x9d\xbe\x87\xe4\xbe\x9e\x90\xd5\xe4\xc9\x25\xfc\x02\x4b\x4a\xac\x92\x7b\x5a\xf6\xe4\xb1\x3d\xb3\xef\xaa\x88\x13\x96\x58\xed\x72\x04\xae\xcc\x64\xdf\xa3\x9f\xef\xea\x3f\x70\x22\xb9\x55\x8b\x9e\x44\x72\xe0\x8b\xb9\x71\xfe\xab\xa7\xe1\xc0\x63\x57\xff\xaa\x8f\x9c\xc7\x99\xd2\xb3\xaf\x9e\xcc\x1b\x8f\xc5\xef\x81\x6b\x6a\x47\xf9\x7f\x5f\x6f\xaf\x8b\xed\x46\x49\xd2\x02\xb3\x93\xbd\x5f\xc6\x7b\xee\x08\xbc\xad\xa3\x3b\x39\x6f\x2c\x3f\x90\x60\xca\xd5\x68\x83\x88\x4f\xc8\x0b\x10\x17\x20\xbe\x12\x01\xe2\x02\xc4\x3b\x22\x40\xbc\xb5\xb2\x00\x71\x01\xe2\xdb\x22\x40\x5c\x80\xb8\x00\xf1\x46\xa7\x00\x71\x01\xe2\x02\xc4\x05\x88\x0b\x10\x17\x20\x2e\x40\x5c\x80\xb8\x00\xf1\x4d\x11\x20\x0e\x02\xc4\x05\x88\x1f\xad\xf7\x3b\x02\xf1\x4b\x01\xe2\x02\xc4\xa3\x08\x10\x17\x20\xde\x11\x01\xe2\xad\x95\x05\x88\x0b\x10\xdf\x16\x01\xe2\x02\xc4\x05\x88\x37\x3a\x05\x88\x0b\x10\x17\x20\x2e\x40\x5c\x80\xb8\x00\x71\x01\xe2\x02\xc4\x05\x88\x6f\x8a\x00\x71\x10\x20\x2e\x40\xfc\x68\xbd\xdf\x11\x88\x3f\x15\x20\x2e\x40\x3c\x8a\x00\x71\x01\xe2\x1d\x11\x20\xde\x5a\x59\x80\xb8\x00\xf1\x6d\x11\x20\x2e\x40\x5c\x80\x78\xa3\x53\x80\xb8\x00\x71\x01\xe2\x02\xc4\x05\x88\x0b\x10\x17\x20\x2e\x40\x5c\x80\xf8\xa6\x08\x10\x07\x01\xe2\x02\xc4\x8f\xd6\xfb\x1d\x81\xf8\x0f\x02\xc4\x05\x88\x47\x11\x20\x2e\x40\xbc\x23\x02\xc4\x5b\x2b\x0b\x10\x17\x20\xbe\x2d\x02\xc4\x05\x88\x0b\x10\x6f\x74\x0a\x10\x17\x20\x2e\x40\x5c\x80\xb8\x00\x71\x01\xe2\x02\xc4\x05\x88\x0b\x10\xdf\x14\x01\xe2\x20\x40\x5c\x80\xf8\xd1\x7a\xbf\x23\x10\xff\xa3\x00\x71\x01\xe2\x51\x04\x88\x0b\x10\xef\x88\x00\xf1\xd6\xca\x02\xc4\x05\x88\x6f\x8b\x00\x71\x01\xe2\x02\xc4\x1b\x9d\x02\xc4\x05\x88\x0b\x10\x17\x20\x2e\x40\x5c\x80\xb8\x00\x71\x01\xe2\x02\xc4\x37\x45\x80\x38\x08\x10\x17\x20\x7e\xb4\xde\xff\x09\x10\x0f\x2d\xff\x09\x00\x00\xff\xff\xf7\x39\x6a\xeb\xfb\x8b\x00\x00")

func config_crd_direct_csi_min_io_directcsivolumes_yaml() ([]byte, error) {
	return bindata_read(
		_config_crd_direct_csi_min_io_directcsivolumes_yaml,
		"config/crd/direct.csi.min.io_directcsivolumes.yaml",
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
	"config/crd/direct.csi.min.io_directcsidrives.yaml":  config_crd_direct_csi_min_io_directcsidrives_yaml,
	"config/crd/direct.csi.min.io_directcsivolumes.yaml": config_crd_direct_csi_min_io_directcsivolumes_yaml,
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
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"config": {nil, map[string]*_bintree_t{
		"crd": {nil, map[string]*_bintree_t{
			"direct.csi.min.io_directcsidrives.yaml":  {config_crd_direct_csi_min_io_directcsidrives_yaml, map[string]*_bintree_t{}},
			"direct.csi.min.io_directcsivolumes.yaml": {config_crd_direct_csi_min_io_directcsivolumes_yaml, map[string]*_bintree_t{}},
		}},
	}},
}}
