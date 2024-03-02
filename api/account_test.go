package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/mahdikarami0111/simplebank/db/sqlc"
	"github.com/mahdikarami0111/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()
	recorder := httptest.NewRecorder()
	acc, err := server.store.Queries.CreateAccount(context.Background(), db.CreateAccountParams{
		Owner:    account.Owner,
		Balance:  account.Balance,
		Currency: account.Currency,
	})
	if err != nil {
		return
	}

	uri := fmt.Sprintf("/accounts/%d", acc.ID)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)
	require.Equal(t, http.StatusOK, recorder.Code)

}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomString(6),
		Balance:  util.RandomInt(0, 100),
		Currency: util.RandomCurrency(),
	}
}
