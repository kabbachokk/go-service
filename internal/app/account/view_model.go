package account

import "fmt"

type accountView struct {
	Id        string `json:"id"`
	FullName  string `json:"full_name"`
	BirthDate string `json:"birth_date"`
}

type AccountView accountView

func newAccountView(user *Response) *AccountView {
	return &AccountView{
		Id:        user.Id,
		FullName:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		BirthDate: fmt.Sprintf("%04d/%02d/%02d", user.BirthYear, user.BirthMonth, user.BirthDay),
	}
}

type AccountsView struct {
	Users []accountView
}

func newAccountsView(users []*Response) *AccountsView {
	res := make([]accountView, 0, len(users))
	for _, user := range users {
		res = append(res, accountView(*newAccountView(user)))
	}
	return &AccountsView{Users: res}
}
