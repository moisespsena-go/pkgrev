package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {
	cookieJar, _ := cookiejar.New(nil)

	baseUrl := "http://localhost:5000/"

	c := &http.Client{
		Jar: cookieJar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	form := url.Values{}
	form.Add("user", "mps@ufv.br")
	form.Add("password", "123@456")

	r, err := http.NewRequest("POST", baseUrl+"auth/password/login", strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Do(r)
	if err != nil {
		panic(err)
	}
	println(resp.Status)
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

	var buf bytes.Buffer

	Url := baseUrl + path.Join("admin", "register", "git.moisespsena.com:moisespsena:almir-estagio", runtime.GOOS, runtime.GOARCH)
	println(Url)
	buf.WriteString(`{
  "Attributes": {
    "CommitID": "9eab9cb0e2bc44dc06ebb913c5af7b2ed47879f2",
    "CommitDate": "2018-12-20T17:57:54-02:00",
    "BuildDate": "2019-02-11T15:45:18-02:00"
  },
  "Assets": {
    "assets.bin": {
      "Attributes": {
        "BuildDate": "2019-02-11T15:45:18-02:00",
        "Hash": "SHA256=4a842cd53b4748896c287ab4bc7dd78d12471ca23b57a1cd966d0e2d9811c566"
      }
    }
  }
}`)
	r, err = http.NewRequest("PUT", Url, &buf)
	if err != nil {
		panic(err)
	}
	resp, err = c.Do(r)
	if err != nil {
		panic(err)
	}
	println(resp.Status)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
