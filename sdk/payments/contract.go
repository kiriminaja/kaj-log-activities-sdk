package payments

type PaymentsContract interface {
	Search(search string) (*ResponseSearch, error)
}
