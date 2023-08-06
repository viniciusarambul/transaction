package usecase

import (
	"github.com/viniciusarambul/transaction/src/entity"
)

type AccountUseCase struct {
	accountRepository entity.AccountRepository
	accountPresenter  entity.AccountPresenter
}

func NewAccountUseCase(repository entity.AccountRepository, presenter entity.AccountPresenter) entity.AccountUseCase {
	return &AccountUseCase{
		accountRepository: repository,
		accountPresenter:  presenter,
	}
}

func (accountUseCase *AccountUseCase) Create(accountInput entity.AccountInput) (entity.Account, error) {
	account := entity.Account{

		Document: accountInput.Document,
	}

	err := accountUseCase.accountRepository.Create(&account)

	if err != nil {
		return entity.Account{}, err
	}

	return account, err
}

func (accountUseCase *AccountUseCase) Find(id int) (entity.AccountOutput, error) {

	account, err := accountUseCase.accountRepository.Find(id)
	if err != nil {
		return entity.AccountOutput{}, err
	}

	accountOutput := entity.AccountOutput{
		ID:       account.ID,
		Document: account.Document,
	}

	return accountOutput, nil
}
