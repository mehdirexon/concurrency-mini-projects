package store

import "final-project/internal/database"

type Store interface {
	// GetAllUsers Users
	GetAllUsers() ([]*database.User, error)
	GetUserByID(id int) (*database.User, error)
	GetUserByEmail(email string) (*database.User, error)
	InsertUser(user database.User) (int, error)
	UpdateUser(u *database.User) error
	DeleteUser(u *database.User) error
	DeleteUserByID(id int) error
	ResetUserPassword(u *database.User, password string) error
	UserPasswordMatches(u *database.User, plainText string) (bool, error)

	// GetAllPlans Plans
	GetAllPlans() ([]*database.Plan, error)
	GetPlanByID(id int) (*database.Plan, error)
	SubscribeUserToPlan(user database.User, plan database.Plan) error
	AmountForDisplay(plan database.Plan) string
}
