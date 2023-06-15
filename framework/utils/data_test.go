package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []string{"1", "2", "3"}
var testDataEncoded = "H4sIAAAAAAAA/4pWMlTSUTJS0lEyVooFBAAA///MXTCHDQAAAA=="

func TestDataExport(t *testing.T) {
	assert.Equal(t, testDataEncoded, DataExport(testData))
}

func TestDataImport(t *testing.T) {
	assert.Equal(t, testData, DataImport(testDataEncoded))
}
