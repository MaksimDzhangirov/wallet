package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	db "github.com/MaksimDzhangirov/wallet/db/sqlc"
)

type transactionRequest struct {
	AccountID   int64  `json:"user_id" binding:"required,min=1"`
	Amount      int64  `json:"amount" binding:"required,gt=0"`
	Currency    string `json:"currency" binding:"required,currency"`
	Description string `json:"description" binding:"required,max=100"`
}

func (server *Server) createDepositTransaction(ctx *gin.Context) {
	server.createTransaction(ctx, "deposit")
}

func (server *Server) createWithdrawalTransaction(ctx *gin.Context) {
	server.createTransaction(ctx, "withdrawal")
}

func (server *Server) createTransaction(ctx *gin.Context, mode string) {
	var req transactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TransferTxParams{
		AccountID:   req.AccountID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	if mode == "withdrawal" {
		account, valid := server.validAccount(ctx, req.AccountID, req.Currency)
		if !valid {
			return
		}
		if account.Balance < req.Amount {
			err := errors.New("not enough funds")
			ctx.JSON(http.StatusUnprocessableEntity, errorResponse(err))
			return
		}

		todayUserWithdrawal, err := server.store.TodayUserWithdrawal(ctx, req.AccountID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		if len(todayUserWithdrawal)+1 > server.config.WithdrawalLimit {
			err := errors.New(fmt.Sprintf("Max withdrawal limit exceeded. Only %d requests per day allowed", server.config.WithdrawalLimit))
			ctx.JSON(http.StatusTooManyRequests, errorResponse(err))
			return
		}
		arg.Amount = -arg.Amount
	} else {
		_, valid := server.validAccount(ctx, req.AccountID, req.Currency)
		if !valid {
			return
		}
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true
}

type listTransactionRequest struct {
	StartDate time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
	StopDate  time.Time `form:"stop_date" binding:"required,gtfield=StartDate" time_format:"2006-01-02"`
	PageID    int32     `form:"page_id" binding:"required,min=1"`
	PageSize  int32     `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTransactions(ctx *gin.Context) {
	var req listTransactionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTransactionsParams{
		CreatedAt:   req.StartDate,
		CreatedAt_2: req.StopDate,
		Limit:       req.PageSize,
		Offset:      (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListTransactions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
