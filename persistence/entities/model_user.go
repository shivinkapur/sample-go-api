package entities

type User struct {
	Id string `json:"id,omitempty"`

	Username string `json:"username,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`

	Phone string `json:"phone,omitempty"`

	// User Status
	UserStatus int32 `json:"user_status,omitempty"`

	Deleted bool `json:"deleted"`

	CreatedAt int64 `json:"created_at,omitempty"`

	ModifiedAt int64 `json:"modified_at,omitempty"`
}
