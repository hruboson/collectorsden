package database

type Category struct {
	ID int		`storm:"id,unique,increment"`
	Name string
	FullPath string //TODO later will be deleted and all paths for this category will be stored in DrivesMappings
	DrivesMappings []string
}

type Content struct {
	ID int `storm:"id,increment"`
	Name string // name can be different from file name
	File string
	ParentCategory int
	FullPath string
}
