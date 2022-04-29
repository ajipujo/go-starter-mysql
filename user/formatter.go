package user

type userFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
	NoNik  string `json:"no_nik"`
	NoTelp string `json:"no_telp"`
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
