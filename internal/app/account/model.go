package account

import "fmt"

type Input struct {
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}

func (i *Input) FormatDate() string {
	return fmt.Sprintf("%04d/%02d/%02d", i.BirthYear, i.BirthMonth, i.BirthDay)
}

type Output struct {
	Id         string
	FirstName  string
	LastName   string
	BirthYear  int
	BirthMonth int
	BirthDay   int
}
