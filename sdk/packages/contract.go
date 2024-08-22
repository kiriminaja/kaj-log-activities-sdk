package packages

type PackageContract interface {
	Search(search string) (*ResponseSearchPackage, error)
}
