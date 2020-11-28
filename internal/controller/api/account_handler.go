package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kskumgk63/containized-firestore/internal/application/user"
)

type accountCRUDHandler struct {
	userUseCase user.UseCase
}

func newAccountCRUDHandler(
	userUseCase user.UseCase,
) CRUDHandler {
	return accountCRUDHandler{
		userUseCase: userUseCase,
	}
}

type account struct {
	UserID string `json:"user_id"`
	ID     string `json:"account_id"`
	Mail   string `json:"mail"`
}

func (h accountCRUDHandler) extractAccount(r *http.Request) (*account, error) {
	body, err := parseBody(r)
	if err != nil {
		return nil, err
	}
	var account *account
	if id, ok := body["account_id"].(string); ok {
		account.ID = id
	}
	if mail, ok := body["mail"].(string); ok {
		account.Mail = mail
	}
	if account == nil {
		return nil, errors.New("account attributes must be set")
	}
	return account, nil
}

func (h accountCRUDHandler) Create(w http.ResponseWriter, r *http.Request) {
	account, err := h.extractAccount(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	out, err := h.userUseCase.Create(r.Context(), &user.CreateInput{
		AccountID: account.ID,
		Mail:      account.Mail,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	account.UserID = out.UserID
	raw, err := json.Marshal(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}

func (h accountCRUDHandler) Read(w http.ResponseWriter, r *http.Request) {
	id := parseIdentifier(r, "/account")
	if id == "" {
		h.loadAccounts(r.Context(), w)
		return
	}
	h.getAccount(r.Context(), w, id)
}

func (h accountCRUDHandler) loadAccounts(ctx context.Context, w http.ResponseWriter) {
	out, err := h.userUseCase.Load(ctx, &user.LoadInput{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if out == nil || len(out.Users) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	accounts := make([]account, len(out.Users))
	for i, user := range out.Users {
		accounts[i] = account{
			UserID: user.ID,
			ID:     user.AccountID,
			Mail:   user.Mail,
		}
	}
	raw, err := json.Marshal(accounts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}

func (h accountCRUDHandler) getAccount(ctx context.Context, w http.ResponseWriter, id string) {
	out, err := h.userUseCase.Get(ctx, &user.GetInput{
		UserID: id,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if out == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	account := account{
		UserID: out.UserID,
		ID:     out.AccountID,
		Mail:   out.Mail,
	}
	raw, err := json.Marshal(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}

func (h accountCRUDHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := parseIdentifier(r, "/account")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errors.New("Please specify the resource").Error()))
		return
	}
	account, err := h.extractAccount(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = h.userUseCase.Update(r.Context(), &user.UpdateInput{
		UserID:    id,
		AccountID: account.ID,
		Mail:      account.Mail,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h accountCRUDHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := parseIdentifier(r, "/account")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errors.New("Please specify the resource").Error()))
		return
	}
	err := h.userUseCase.Delete(r.Context(), &user.DeleteInput{
		UserID: id,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
