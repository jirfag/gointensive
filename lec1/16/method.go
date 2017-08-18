package main

type User struct {
	ID   int
	name string
}

type UserStorage struct {
	// ???
}

func (us *UserStorage) RegisterUser(name string) error         {}
func (us UserStorage) GetUserByName(name string) (User, error) {}

func getOrCreateUser(us *UserStorage, name string) (User, error) {}

func main() {

}
