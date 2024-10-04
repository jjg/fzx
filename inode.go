package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// TODO: Revisit these types once basic compatibility testing works.
type Block struct {
	Block_hash string `json:"block_hash"`
	Last_seen  string `json:"last_seen"`
}

type Inode struct {
	Url               string  `json:"url"`
	Created           int     `json:"created"` // TODO: Figure out how to parse as datetime
	Version           int     `json:"version"`
	Private           bool    `json:"private"`
	Encrypted         bool    `json:"encrypted"`
	Fingerprint       string  `json:"fingerprint"`
	Access_key        string  `json:"access_key"`
	Content_type      string  `json:"content_type"`
	File_size         int     `json:"file_size"`
	Block_size        int     `json:"block_size"`
	Blocks_replicated int     `json:"blocks_replicated"`
	Inode_replicated  int     `json:"inode_replicated"`
	Blocks            []Block `json:"blocks"` // TODO: Consider using a pointer here
	Media_type        string  `json:"media_type"`
}

func stringToSha1(in string) string {
	h := sha1.New()
	h.Write([]byte(in))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (i *Inode) Load(storageLocation *string, fzxPath *string) error {
	var err error
	var inodeJson []byte
	var fingerprint = stringToSha1(*fzxPath)

	// TODO: Refactor to use storage driver
	inodeFilename := fmt.Sprintf("%v/%v.json", *storageLocation, fingerprint)
	if inodeJson, err = ioutil.ReadFile(inodeFilename); err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(inodeJson), i); err != nil {
		return err
	}

	return nil
}
