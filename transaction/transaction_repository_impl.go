package transaction

type RepositoryImpl struct{}

func NewTransactionRepositoryImpl() *RepositoryImpl {
	return &RepositoryImpl{}

}
