package presenter

import "github.com/viniciusarambul/transaction/src/entity"

type accountPresenter struct{}

func NewAccountPresenter() entity.AccountPresenter {
	return &accountPresenter{}
}

func (accountPresenter accountPresenter) Output(account entity.Account) entity.AccountOutput {
	return entity.AccountOutput{
		Document: account.Document,
		LimitMax: account.LimitMax,
	}
}
