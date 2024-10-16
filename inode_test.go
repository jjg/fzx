package main

import (
	"fmt"
	"testing"
	"time"
)

// This test was created to ensure basic compatibility with existing JSFS pools.
// Once that's established it will probably get mutated into a more generic
// inode loading test.
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

func TestStoreInode(t *testing.T) {

	// Create a new node
	inode := &Inode{Url: "/localhost/utils/jsfsmoke2"}

	// Store it
	storageLocation := "./"
	if err := inode.Store(&storageLocation); err != nil {
		t.Fatal(err)
	}
}

func TestModifyInode(t *testing.T) {

	// Load an existing JSFS inode
	// Taken from files in testdata dir
	fzxPath := "/localhost/utils/jsfsmoke"
	storageLocation := "./testdata/blocks"

	inode := &Inode{}
	if err := inode.Load(&storageLocation, &fzxPath); err != nil {
		t.Fatal(err)
	}

	// Modify version
	inode.Version = inode.Version + 1

	// Modify Created
	now := time.Now()
	inode.Created = now.UnixMilli()

	// Store the modified inode
	if err := inode.Store(&storageLocation); err != nil {
		t.Fatal(err)
	}

	// NOTE: At this point JSFS should be used to test the modified inode
}

func msToTime(ms int64) time.Time {
	return time.UnixMilli(ms)
}

func timeToMs(t time.Time) int64 {
	return t.UnixMilli()
}
