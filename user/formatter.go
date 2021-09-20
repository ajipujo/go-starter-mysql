package user

type userFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	RoleID int    `json:"role_id"`
}

func FormatterUser(user User, token string) userFormatter {
	var userFormat userFormatter

	userFormat.ID = user.ID
	userFormat.Name = user.Name
	userFormat.Email = user.Email
	userFormat.Token = token
	userFormat.RoleID = user.RoleID

	return userFormat
}
