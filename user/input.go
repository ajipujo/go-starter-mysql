package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Alamat   string `json:"alamat"`
	NoNik    string `json:"no_nik"`
	NoTelp   string `json:"no_telp"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Name   string `json:"name" binding:"required"`
	Alamat string `json:"alamat"`
	NoNik  string `json:"no_nik"`
	NoTelp string `json:"no_telp"`
	Email  string `json:"email" binding:"required,email"`
}

type CheckEmailAvailabilityInput struct {
	Email string `json:"email" binding:"required,email"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
