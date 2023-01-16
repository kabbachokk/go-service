package account

type UseCaseInterface interface {
	CreateAccount(obj *Request) (data *Response, err error)
	GetAccount(id int) (data *Response, err error)
	UpdateAccount(id int, obj *Request) (data *Response, err error)
	DeleteAccount(id int) (err error)
}

type UseCase struct {
	repository *RepositoryInterface
}

type Request struct {
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}

type Response struct {
	Id         string
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}

func NewUseCase(r *RepositoryInterface) *UseCase {
	return &UseCase{
		repository: r,
	}
}

var _ UseCaseInterface = (*UseCase)(nil)

func (uc *UseCase) CreateAccount(obj *Request) (data *Response, err error) {
	in := Input(*obj)
	out, err := (*uc.repository).CreateAccount(&in)
	if err != nil {
		return
	}
	res := Response(*out)
	data = &res
	return
}

func (uc *UseCase) GetAccount(id int) (data *Response, err error) {
	out, err := (*uc.repository).FindAccountById(id)
	if err != nil {
		return
	}
	res := Response(*out)
	data = &res
	return
}

func (uc *UseCase) UpdateAccount(id int, obj *Request) (data *Response, err error) {
	in := Input(*obj)
	out, err := (*uc.repository).UpdateAccountById(id, &in)
	if err != nil {
		return
	}
	res := Response(*out)
	data = &res
	return
}

func (uc *UseCase) DeleteAccount(id int) (err error) {
	err = (*uc.repository).DeleteAccountById(id)
	return
}
