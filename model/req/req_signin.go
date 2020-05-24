package req

type RepSignIn struct {
	Email    string `json:"email, omitempty" validate:"required"`
	Password string `json:"password, omitempty" validate:"required"`
}
