package store

type Store interface {
	// GetAllUsers Users
	GetAllUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	InsertUser(user User) (int, error)
	UpdateUser(u *User) error
	DeleteUser(u *User) error
	DeleteUserByID(id int) error
	ResetUserPassword(u *User, password string) error
	UserPasswordMatches(u *User, plainText string) (bool, error)

	// GetAllPlans Plans
	GetAllPlans() ([]*Plan, error)
	GetPlanByID(id int) (*Plan, error)
	SubscribeUserToPlan(user User, plan Plan) error
	AmountForDisplay(plan Plan) string
}
