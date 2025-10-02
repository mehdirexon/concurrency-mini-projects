package store

// ---------- User methods ----------

func (m *Models) GetAllUsers() ([]*User, error) {
	return m.User.GetAll()
}

func (m *Models) GetUserByID(id int) (*User, error) {
	return m.User.GetOne(id)
}

func (m *Models) GetUserByEmail(email string) (*User, error) {
	return m.User.GetByEmail(email)
}

func (m *Models) InsertUser(user User) (int, error) {
	return m.User.Insert(user)
}

func (m *Models) UpdateUser(u *User) error {
	return u.Update()
}

func (m *Models) DeleteUser(u *User) error {
	return u.Delete()
}

func (m *Models) DeleteUserByID(id int) error {
	return m.User.DeleteByID(id)
}

func (m *Models) ResetUserPassword(u *User, password string) error {
	return u.ResetPassword(password)
}

func (m *Models) UserPasswordMatches(u *User, plainText string) (bool, error) {
	return u.PasswordMatches(plainText)
}

// ---------- Plan methods ----------

func (m *Models) GetAllPlans() ([]*Plan, error) {
	return m.Plan.GetAll()
}

func (m *Models) GetPlanByID(id int) (*Plan, error) {
	return m.Plan.GetOne(id)
}

func (m *Models) SubscribeUserToPlan(user User, plan Plan) error {
	return m.Plan.SubscribeUserToPlan(user, plan)
}

func (m *Models) AmountForDisplay(plan Plan) string {
	return plan.AmountForDisplay()
}
