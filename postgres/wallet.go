package postgres

import (
	"database/sql"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) Wallets(walletType string) ([]wallet.Wallet, error) {
	var rows *sql.Rows
	var err error
	if walletType != "" {
		rows, err = p.Db.Query("SELECT * FROM user_wallet WHERE wallet_type = $1", walletType)
	} else {
		rows, err = p.Db.Query("SELECT * FROM user_wallet")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

// func WalletByType(wallet_type string) (string, []interface{}) {
// 	var query string
// 	var args []interface{}
// 	if wallet_type != "" {
// 		query += " WHERE wallet_type = $1"
// 		args = append(args, wallet_type)
// 	}
// 	return "SELECT * FROM user_wallet " + query, args
// }

func (p *Postgres) WalletByUserID(id int) ([]wallet.Wallet, error) {
	query, args := WalletByUserIDQuery(id)
	rows, err := p.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func WalletByUserIDQuery(id int) (string, []interface{}) {
	return "SELECT * FROM user_wallet WHERE user_id = $1", []interface{}{id}
}
