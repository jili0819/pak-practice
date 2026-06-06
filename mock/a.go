//go:generate mockgen -destination=a_mock.go -package=mock . UserRepository
package mock

type User struct {
	ID   string
	Name string
}

type UserRepository interface {
	GetByID(id int64) (*User, error)
	Save(user *User) error
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUserName(id int64) (string, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
