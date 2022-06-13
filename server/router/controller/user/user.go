package user

type Data struct {
	Name     string `json: "name"`
	Lastname string `json: "lastname"`
	Email    string `json: "email"`
	Password string `json: "password'`
	Confirm  string `json: "confirm"`
	Accept   bool   `json: "accept"`
}

var group struct {
	users []Data
}

func NewUser(user Data) {

	group.users = append(group.users, user)

}

func SearchUser(email string, password string) bool {

	for _, user := range group.users {

		if user.Email == email {
			if user.Password == password {
				return true
			}
		}
	}

	return false
}

func GetUsers() []Data {
	return group.users
}
