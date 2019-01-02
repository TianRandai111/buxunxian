package main

import (
	"encoding/json"
	"io/ioutil"
)

type student struct {
	Name string
	Sex  string
	Age  int
}

func (p *student) Save() (err error) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("/tmp/go_studen_test.dat", data, 0755)
	return
}

func (p *student) Load() (err error) {
	data, err := ioutil.ReadFile("/tmp/go_data.bat")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, p)
	return
}
