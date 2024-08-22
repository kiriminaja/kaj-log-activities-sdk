package withdrawals

type WithdrawalContract interface {
	Search(search string) (*ResponseSearch, error)
}
