package tests

import (
	"fmt"
	"io"
	"os"
	"testing"
)

const testFilesDir = "../../data/testFiles"

func loadDataFile(t *testing.T, fileName string) []byte {
	f, err := os.Open(fmt.Sprintf("%s/%s", testFilesDir, fileName))
	if err != nil {
		t.Errorf(err.Error())
	}

	fileData, err := io.ReadAll(f)
	if err != nil {
		t.Errorf(err.Error())
	}
	return fileData
}
