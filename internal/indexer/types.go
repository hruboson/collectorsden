package indexer

type NodeType int64

const (
	FOLDER NodeType = iota
	FILE
	SYMLINK
)

type FileType int64

const (
	PNG FileType = iota
	JPG
	//TODO
)
