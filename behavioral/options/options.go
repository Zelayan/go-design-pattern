package options

type User struct {
	Name string
	Address string

	Age int
	Email string
	Qq string
}


func NewUser(name string, opt ...Option) *User{

	user := &User{Name: name}

	for _, ops := range opt {
		ops(user)
	}

	return user
}

type Option func(*User)

func WithAddress (address string) Option {
	return func(u *User) {
		u.Address = address
	}
}

func WithName(name string) Option {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int) Option {
	return func(u *User) {
		u.Age =age
	}
}

func WithEmail(email string) Option {
	return func(u *User) {
		u.Email = email
	}
}

func WithQq(qq string) Option {
	return func(u *User) {
		u.Qq = qq
	}
}