package user

type InputRegister struct {
	Fullname     string `json:"nama_lengkap" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	NoTlpn       string `json:"no_tlpn" binding:"required"`
	BusinessName string `json:"nama_bisnis" binding:"required"`
	Password     string `json:"kata_sandi" binding:"required"`
}

type InputLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"kata_sandi" binding:"required"`
}

type InputCheckEmail struct {
	Email string `json:"email" binding:"required,email"`
}

type InputUpdate struct {
	Fullname     string `json:"nama_lengkap" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	NoTlpn       string `json:"no_tlpn" binding:"required"`
	BusinessName string `json:"nama_bisnis" binding:"required"`
}
