package file

import (
	"log"
	"testing"
)

func TestFileExist(t *testing.T) {
	fe := FileExist("testfile.txt")
	if !fe {
		t.Errorf("File reading failed")
	} else {
		log.Println(fe)
	}
}

func TestFileWrite(t *testing.T) {
	err := WriteFile("This is a test\r\n", "testfile.txt")
	if err != nil {
		t.Errorf("File writing failed")
	}
	err = WriteFile("This is a second test", "testfile.txt")
	if err != nil {
		t.Errorf("File writing 2nd failed")
	}
	fe := FileExist("testfile.txt")
	if !fe {
		t.Errorf("File reading failed")
	} else {
		log.Println(fe)
	}
}
