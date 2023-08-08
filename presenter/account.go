package presenter

import "github.com/viniciusarambul/transaction/entity"

type accountPresenter struct{}

func NewAccountPresenter() entity.AccountPresenter {
	return &accountPresenter{}
}

func (accountPresenter accountPresenter) Output(account entity.Account) entity.AccountOutput {
	doc := entity.RemoveLGPDFromResponse(account.Document)

	return entity.AccountOutput{
		ID:       account.ID,
		Document: doc,
		LimitMax: account.LimitMax,
	}
}
