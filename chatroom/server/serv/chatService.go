package serv

type (
	User struct {
		Id   int
		Name string
	}

	UserMsg struct {
		User User
		Msg  string
		Time string
	}
)
