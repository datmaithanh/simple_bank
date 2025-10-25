package api

import (
	"testing"

	mockdb "github.com/datmaithanh/simplebank/db/mock"
	db "github.com/datmaithanh/simplebank/db/sqlc"
	"github.com/datmaithanh/simplebank/util"
	"github.com/golang/mock/gomock"
)

func TestGetAccount(t *testing.T) {

	account := ramdomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.
		EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)


}

func ramdomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
