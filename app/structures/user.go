package structures

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser() *User {
	return &User{}
}

func (m *User) User() {
	// Model1 için metodlar burada yer alır
}
func (p *User) TableName() string {
	return "calendar.users"
}
