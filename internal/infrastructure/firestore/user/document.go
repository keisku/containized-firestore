package user

type accountData struct {
	AccountID string `firestore:"account_id"`
	Mail      string `firestore:"mail"`
}
