package account

import (
	"github.com/stretchr/testify/mock"
)

type useCaseMock struct {
	mock.Mock
}

var _ UseCaseInterface = (*useCaseMock)(nil)

func (m useCaseMock) CreateAccount(obj *Request) (data *Response, err error) {
	args := m.Called(obj)
	return args.Get(0).(*Response), args.Error(1)
}

func (m useCaseMock) GetAccount(id int) (data *Response, err error) {
	args := m.Called(id)
	return args.Get(0).(*Response), args.Error(1)
}

func (m useCaseMock) UpdateAccount(id int, obj *Request) (data *Response, err error) {
	args := m.Called(id, obj)
	return args.Get(0).(*Response), args.Error(1)
}

func (m useCaseMock) DeleteAccount(id int) (err error) {
	args := m.Called(id)
	return args.Error(0)
}
