package design

type User struct {
	//必要字段：ID和名字
	Id   int
	Name string

	//可选字段：地址
	Address string
}

type Option func(user *User)

func NewUser(id int, name string, opt ...Option) *User {
	u := &User{
		Id:   id,
		Name: name,
	}
	for _, o := range opt {
		o(u)
	}
	return u
}

func WithAddress(address string) Option {
	return func(u *User) {
		u.Address = address
	}
}
