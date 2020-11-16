// Code generated by "esc -o bindata.go -pkg schema -ignore .*\.go -private -modtime=1518458244 data"; DO NOT EDIT.

package schema

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/data/compose-spec.json": {
		name:    "compose-spec.json",
		local:   "data/compose-spec.json",
		size:    25664,
		modtime: 1518458244,
		compressed: `
H4sIAAAAAAAC/+w8S4/juNF3/wpBu7ft13z4sMDMbZFTgCwSYCcBkoZXoKWyzGmK5JCUu72D/u8BRUnW
gxIpWXZ7Nn0YTFsqklWlYr3Jb6sgCH+U8Q4yFH4Kwp1S/NP9/RfJ6K15esdEep8ItFX3//fw4ePtw8d7
8+KH8EYPxokeF7OMMwmR5BDf6dHmpTpw0K/Z5gvEqnyGFSke/sWMCX7jEOMtjpHC1bgEZCwwLx58CsLP
Owgq6C0mEGAZoODfv/z6N/MzgS2mmKYBCrKcKHwbM6oQpiBksEESkgBxTsoV7sKblV6DC8ZBKAwy/BRo
PgRBuAchzZrmQYMEqQSmaYFe8byD4r/MyIBtA9XAVjZpC3IJyV3wmTEiA8pUgDNOIAOqNO4CvuZYQBKU
SAS//vO3z4EAzblizpjRLU5zYebShN+FBTavBUFBEEoQexw3CKq/zw/3R3Lva7CbLpGN71Q850gpEPQf
fVYVr39/RLd//HL7n4fbj3fR7fqnH1uvtWQJ2JrlzSfSmNfrhzXka/nXa70wSpICGJHW2ltEJLRppqCe
mXhy0VyDvRHN5foWmtvk7BnJM+cXrKDeiBiz/DLfT0IsQLlF1kC9mcTq5Zch2GxjF8EV1BsRbJY/jeBV
RbQdx/D3l1v9/2sx5+h8ZpYGfgURLZ1nY6dN5wzzs2boACcT4IQdCsztPDMAWp2HNZuCINzkmCRdrjMK
f9dTPDYeBsG3rrVpzFO8b/0aFor6/QAt9XttJ+FFFUSNL21YwOInENry+I5Awkj6AMsIlipiIkpwrKzj
CdoAOWmGGMU7iLaCZc5ZtpGhRFonqjS4J+UKiRS8OSt3WSTxHy2+PoaYKkhBhDf12LVtMLwogaIdk+ok
TmHJCCr9mT7SHfjeBI4N3BXMcY3QAm/+Wq8sCIQb8oRZVOqszk4b2SJj2yNMQOuLSABKog237Z96aiQE
OvQ3H1aQjX0QgzXBGVYd9naY28IFs+tA5llgBdfCGYPMm7PmGXC6a+/5agvbQSOD/9mQLhEa3k8ttNw2
vbvxYsQjlCQtikuEmyj2FGAQ5hR/zeGvJYgSOXTnTQTjy0+cCpbziCOhzfS4ctZBbYboUrZ7Ch1uldf3
Ih1Sc1ytLYh2agIPs2XxJxz+iNsj0aaQ5SL2dTCmGtogCHOc+AOnU4AzlrTxpnm20Xu/B2x1lSbYz4kW
tL3nmyLVfNMRL5NAiSjKwLlReB7FLKd21XcThBmmOMu1B/bQHcdBxOA1Uv9CL+WvDw+9meQOCZBtD6r8
AnYHqhj1NWcKTR3EQWCWTB0l1PyBIqcKZzBx5FRuSHArRQEJUIURKZJ9S7lbR+fNsddCz/AjFJBiqcTB
5c36m8A2MZ77r8m7BDjQREatDOOoTZkVrkwOE0933fv6yZV48LcZk3Wjj5kxImdmHQQZTv/2YzBaqKXH
Kh8RSYWEgqTYbOWjHSCidodwbZ3k1fLUbmeqTHGx3pGK/rTjVsDtZJTedek0iZyAnBdElzPJxf25hI5h
ZKbRQa/GLewMjBhXNoROxCeSgES8m4kWyxCmPjYXqBIHzrAxnVfnnwLdR7WmnswGoHssGM0qx8Avo9EY
/8KZhNOd49pyVtu/9unWXavERIY0stXagxamL0B2Br4obSsulTicmjlspF198luDRtuZVmopu2rV9SQT
NlHrac4LikhEMH1aXmWdlq0LjS4ug+6zSfep0mtMXbyD+GmEyCZUazSTykcH4gylbiCK22p+wxgBRNtA
PHbOM5oebQDOzlmH5xE4wtJUQ7o8de/MqMB7ED4uOOPHas3UyoWnl3lnHMsRWS7+IiRc+3tDFwwJMhTr
zSxASpdcZZCVWcgJQZ0eJEDrzZ7sNnlVRdy9sfIZcY5pFz1LWlODa+jpOJaVlaiXP7EwwVLuv4IAaumY
aCgpN5R48w+lfKMjI/kEIznX5+/Xl/j+/z1l3Tb257ljtVaNCIsRiTBfihguMBNYHfzyfSPR3Kz4dkYG
cBSFJgmEWAnwiC7P6ZQxlkVPmJAowRJtOs6kzaLrATJmAiKUfHFnF28/PDz0MoytFCPHybClKexLG1hO
V4ScIKUDCZcS5Eyoi6T+j+ge4xuzeL8a0EuV+Aw6TwnBw5CMFxAGqnYVAvmGYLmDZMoYwRSLGfEJgq6r
PjDH/ecC7zGBtMMi2z7lgungcG5iieeERJwRHFsTvTfHzFzLpSPP6CCLrQt7I+B4G1GmIq6dJdpo9Wpt
0KICzyg5OMkSUCQAXTu5kdZftcraXECMVME/zeWb1XBO0voFZIzIYAKkElX7SIhzbdmG0mMnxSL2Ppdx
rdjvEHwvdb6XOhcrdcqDjNW8eF2qBNOIcaBOfSAV41EqUAyWGqDVZCZlp3d/GolTiohLtaiMb2cmx5Vy
K7i88G7kpPqfM0oyEZI9MBoJiryM8Ozik08taYfEBG+g2PrbAZdjXhmomO+mRGS9UBFtYoDRfbIeTrBY
d2IunWmlAoZKr1RBv5X++7Abre9agK9nWZdyJU+Nfm5b5O06tSu0EksF1O7c2QdtcK9la2qKxC9BUkCh
1LtldjQCn5ECmJEAsHOsPM5xEZ5RFjM+IAPfC79qy3p+dlW+8nAiqpfVGAZt9FINZaiu8BP86YLj0i7V
ZxKmGqdRNeiLxDMTT9r9S7AY012v4+eM2md4Jh6E6tQIx07vNEHdJ6KGDvB4Zg71ThJ7u4Pt9tAFKIE7
nS5VcNR0sUFeZ+OGwhmwXM0NT5C1ydExw6TDZb4SuWpMGTbOZjlErQHZlbTHWtSqlK9T5nz8VKBJ0cfj
5dQKKA4yu6tx86vQghGyQfHTwgdqOBKIECBYZl4nJBIg6DBLDE3HCcIkFxCh2KNmX34rihUT85fM0EtU
LVuAOJSA2fQiAeGfujxus9stFlKZpBnj5a+2pXqjmnbOE6TgXXzexWeW+AgwsahcSnSsmaplzs9Oa7Fv
tkowd2f6pc9Z9hro67aR74V5FugUKAgcRy2pGjCJfVjrKehGL/PQSXADcSWHZi+7d40LWFfkFjoccuzU
d+nhExW/1sKa8Iwr6XeSE9OEPU/3di/8ZThBMXQc4FM/ilQCYaomtyx2WcgFbEEAjeGk87tnKtJxASj5
Lur2NlmuAgYdlUW0G2HYqg7nFsqzRnurYaU/FvX1B/TSDG3ps0jdsLQNS1mYYBkLUFCvbDtV55LkcSkO
n8qMuNNchntEco+6blfO/EV8gnjb1ZjfUp7LdKSmYdjHMwQDt2edR0JixNEGE1xTMaN5MeydznU1/x5P
f+Fk/rJVq7gjZ1n3ifvlKi4uGI2rX8YEowJbIG/kc/rBqw+/hIoYX75W7+69X7s3NOYoW+zOFu+TCdbc
xDU4G/mGgjv5asAizOueaHulCvNIIJpOKOumSMEzmlBuRflLhQScXBRbrqLUkU1rAvxqS0wLnqNZIsD0
4+UbxTXVab0BHfJYl3huaiatvRVKpYZXvRTASH/maI/mG3KqqGt1G8dsBTCgaEMgwnz/sxMWKYXinVdd
bWJF4gJhQq/xwWrWS6h3qz7Bqr/vSt9deX27orzs1HmhZgE1uxrvsxc8rnq5jKSN3h+zxGf90ykABRkn
SEE0Qs4FZLnn51tluYT635DlM2rNhXfClchQp1G+IUv9hp2xz+udQlk1+3NqNLpgljvqhxJQg0gNdYt1
Fi2/zTjlC6qgu59GQp2xI/5nuudygTNQ9m/avFrU8j199Q9HaudsHELK43bhWfeaN0g53oB6TlrG7lk9
7Wb2ZkVrVZ8/E71LWofMSDW+d9+4RoQeeu1939rd/uau8HYetgNi7jxpWOC1VxbZdgu55Sqx4jbwgTNT
7dzoSv97/W8AAAD///7BNlpAZAAA
`,
	},

	"/data": {
		name:  "data",
		local: `data`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"data": {
		_escData["/data/compose-spec.json"],
	},
}
