package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

func DataExport(data []string) string {
	var err error

	je, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err = gz.Write(je)
	if err != nil {
		panic(err)
	}
	if err = gz.Close(); err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func DataImport(v string) []string {
	var err error

	be, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	_, err = buf.Write(be)
	if err != nil {
		panic(err)
	}

	gz, err := gzip.NewReader(&buf)
	if err != nil {
		panic(err)
	}

	je, err := ioutil.ReadAll(gz)
	if err != nil {
		panic(err)
	}

	var data []string
	err = json.Unmarshal(je, &data)
	if err != nil {
		panic(err)
	}

	return data
}
