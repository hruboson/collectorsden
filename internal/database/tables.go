package database

type Category struct {
	ID int		`storm:"id,unique,increment"`
	Name string
	Folder string
	DrivesMappings []string
}

type Content struct {
	ID int `storm:"id,increment"`
	Name string // name can be different from file name
	File string
	ParentCategory int
	FullPath string
}
