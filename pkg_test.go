package file_rotater

import (
	"log"
	"os"
	"testing"
	"time"
)

var (
	filePath = "./app.log"
)

func TestDoRotate(t *testing.T) {
	fw, err := NewFileRotater(filePath)
	if err != nil {
		log.Printf("new rotater failed failed: %s", err)
		t.FailNow()
	}
	err = fw.doRotate()
	if err != nil {
		log.Printf("rotate failed: %s", err)
		t.FailNow()
	}
	time.Sleep(2 * time.Second)

}

func TestDoRotateFileNotFound(t *testing.T) {
	fw, err := NewFileRotater(filePath)
	if err != nil {
		log.Printf("new rotater failed: %s", err)
		t.FailNow()
	}
	os.Remove(filePath)
	err = fw.doRotate()
	if err != nil {
		log.Printf("rotate failed: %s", err)
		t.FailNow()
	}
	time.Sleep(2 * time.Second)

}

func TestDeleteOldFile(t *testing.T) {
	fw, err := NewFileRotater(filePath)
	if err != nil {
		log.Printf("new rotater failed: %s", err)
		t.FailNow()
	}
	fw.deleteOldFiles()

}
