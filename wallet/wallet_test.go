package wallet

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"time"

	"github.com/labstack/echo/v4"
)

type StubWallet struct {
	wallet []Wallet
	err    error
}

func (s StubWallet) Wallets(wallet_type string) ([]Wallet, error) {
	return s.wallet, s.err
}

func TestWallet(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubError := StubWallet{err: echo.ErrInternalServerError}
		p := New(stubError)

		p.GetAllWalletsHandler(c)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}

	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		createdAt, _ := time.Parse(time.RFC3339, "2024-03-25T14:19:00.729237Z")
		stubError := StubWallet{wallet: []Wallet{{ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John's Savings", WalletType: "Savings", Balance: 1000.00, CreatedAt: createdAt}}}
		p := New(stubError)

		p.GetAllWalletsHandler(c)

		want := []Wallet{{ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John's Savings", WalletType: "Savings", Balance: 1000.00, CreatedAt: createdAt}}
		gotJSON := rec.Body.Bytes()
		var got []Wallet
		if err := json.Unmarshal(gotJSON, &got); err != nil {
			t.Errorf("unable to unmarshal json %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}
