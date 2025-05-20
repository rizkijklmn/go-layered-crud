package user

// Service berisi logika bisnis
func CreateNewUser(name, email string) (*User, error) {
	user := &User{Name: name, Email: email}
	err := CreateUser(user)
	return user, err
}

func GetUsersService() ([]User, error) {
	return GetAllUsers()
}

func GetUserService(id uint) (User, error) {
	return GetUserByID(id)
}

func UpdateUserService(id uint, name, email string) (User, error) {
	user, err := GetUserByID(id)
	if err != nil {
		return user, err
	}

	user.Name = name
	user.Email = email

	err = UpdateUser(&user)
	return user, err
}

func DeleteUserService(id uint) error {
	return DeleteUser(id)
}
