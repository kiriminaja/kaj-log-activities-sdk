package packages

type PackageContract interface {
	Search(search string) (*ResponseSearch, error)
}
