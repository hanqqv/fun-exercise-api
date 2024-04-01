package wallet

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

var db *sql.DB

func GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, user_id, user_name, wallet_name, wallet_type, balance, created_at FROM wallets WHERE user_id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	row := stmt.QueryRow(id)
	u := Wallet{}
	err = row.Scan(&u.ID, &u.UserID, &u.UserName, &u.WalletName, &u.WalletType, &u.Balance, &u.CreatedAt)
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, u)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user" + err.Error()})
	}
}
