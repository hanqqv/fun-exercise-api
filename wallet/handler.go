package wallet

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets(wallet_type string) ([]Wallet, error)
	WalletByUserID(id int) ([]Wallet, error)
	CreateWallet(wallet Wallet) error
	UpdateWallet(id int, wallet Wallet) error
	DeleteWallet(id int) error
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// GetAllWalletsHandler
//
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Failure		500	{object}	Err
//	@Router /api/v1/wallets [get]
func (h *Handler) GetAllWalletsHandler(c echo.Context) error {
	walletType := c.QueryParam("wallet_type")
	wallets, err := h.store.Wallets(walletType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}

// GetWalletByIDHandler
//
//	@Summary		Get wallet by user id
//	@Description	Get wallet by user id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/users/:id/wallets [get]
//	@Failure		500	{object}	Err
//	@Router /api/v1/users/:id/wallets [get]
func (h *Handler) GetWalletByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	wallets, err := h.store.WalletByUserID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}

// CreateWalletHandler
//
//	@Summary		Create wallet
//	@Description	Create wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	Wallet
//	@Router			/api/v1/wallets [post]
//	@Failure		500	{object}	Err
//	@Router /api/v1/wallets [post]
func (h *Handler) CreateWalletHandler(c echo.Context) error {
	var wallet Wallet
	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	err := h.store.CreateWallet(wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, wallet)
}

// UpdateWalletHandler
//
//	@Summary		Update wallet by id
//	@Description	Update wallet by id
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets/:id [put]
//	@Failure		500	{object}	Err
//	@Router /api/v1/wallets/:id [put]
func (h *Handler) UpdateWalletHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	var wallet Wallet
	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	err = h.store.UpdateWallet(id, wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallet)
}

// DeleteWalletByIDHandler
//
//	@Summary		Delete wallet by user_id
//	@Description	Delete wallet by user_id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		204	{object}	Wallet
//	@Router			/api/v1/users/:id/wallets [delete]
//	@Failure		500	{object}	Err
//	@Router /api/v1/users/:id/wallets [delete]
func (h *Handler) DeleteWalletByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	err = h.store.DeleteWallet(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}
