package player

type Avatar struct {
	Url string
}

type Player struct {
	Name     string
	Age      int
	Avatar   Avatar
	password string
}

func New(name string) Player {
	return Player{
		Name:     name,
		password: "defaultPassword",
	}
}
