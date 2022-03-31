package API

import (
	"io/ioutil"
	"net/http"
)

type input string

func (s input) ReadURL() ([]byte, error) {
	resp, _ := (http.Get(string(s)))
	return ioutil.ReadAll(resp.Body)
}
