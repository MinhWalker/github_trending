package req

type ReqUpdate struct {
	FullName string `json:"fullName,omitempty" validate:"required"` // tags
	Email    string `json:"email,omitempty" validate:"required"`
}