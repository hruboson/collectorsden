package database

type Category struct {
	ID int		`storm:"id,increment"`
	Name string
	Folder string
	DrivesMappings []string
}
