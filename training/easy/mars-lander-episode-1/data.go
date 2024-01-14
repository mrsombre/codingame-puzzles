package main

// Export/Import of data in the form of string arrays.
// This can be used to unload the conditions of a problem (input)
// in a compressed form into the debug console and unpack it in the IDE.

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
)

// DataExport serializes and compresses a slice of strings,
// returning a base64 encoded string.
func DataExport(data []string) string {
	var err error

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	var gzBuf bytes.Buffer
	gz := gzip.NewWriter(&gzBuf)
	if _, err = gz.Write(jsonData); err != nil {
		panic(err)
	}
	if err = gz.Close(); err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(gzBuf.Bytes())
}

// DataImport decodes a base64 string, decompresses it,
// and deserializes the JSON data into a slice of strings.
func DataImport(encodedData string) []string {
	var err error

	gzData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		panic(err)
	}

	gz, err := gzip.NewReader(bytes.NewBuffer(gzData))
	if err != nil {
		panic(err)
	}
	defer func(gz *gzip.Reader) {
		err = gz.Close()
		if err != nil {
			panic(err)
		}
	}(gz)

	var jsonData bytes.Buffer
	if _, err = jsonData.ReadFrom(gz); err != nil {
		panic(err)
	}

	var data []string
	if err = json.Unmarshal(jsonData.Bytes(), &data); err != nil {
		panic(err)
	}

	return data
}
