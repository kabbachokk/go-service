package account

import (
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

var _ RepositoryInterface = (*repositoryMock)(nil)

func (m *repositoryMock) CreateAccount(in *Input) (user *Output, err error) {
	args := m.Called(in)
	return args.Get(0).(*Output), args.Error(1)
}

func (m *repositoryMock) FindAccountById(pk int) (user *Output, err error) {
	args := m.Called(pk)
	return args.Get(0).(*Output), args.Error(1)
}

func (m *repositoryMock) UpdateAccountById(pk int, in *Input) (user *Output, err error) {
	args := m.Called(pk, in)
	return args.Get(0).(*Output), args.Error(1)
}

func (m *repositoryMock) DeleteAccountById(pk int) (err error) {
	args := m.Called(pk)
	return args.Error(0)
}
