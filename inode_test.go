package main

import (
	"fmt"
	"testing"
)

func TestLoadExistingInode(t *testing.T) {

	// Taken from files in testdata dir
	fzxPath := "/localhost/utils/jsfsmoke"
	storageLocation := "./testdata/blocks"

	inode := &Inode{}
	if err := inode.Load(&storageLocation, &fzxPath); err != nil {
		t.Fatal(err)
	}

	// Compare loaded to original values.
	if inode.Url != "/localhost/utils/jsfsmoke" {
		fmt.Printf("wanted: '%s', got: '%s'\n", "/localhost/utils/jsfsmoke", inode.Url)
		t.Fatal("url mismatch")
	}
}
