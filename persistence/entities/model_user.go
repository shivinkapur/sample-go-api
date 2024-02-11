package entities

type User struct {
	Id string `json:"id,omitempty"`

	Username string `json:"username,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`

	Phone string `json:"phone,omitempty"`

	// User Status
	UserStatus int32 `json:"userStatus,omitempty"`

	Deleted bool `json:"deleted"`

	CreatedAt int64 `json:"createdAt,omitempty"`

	ModifiedAt int64 `json:"modifiedAt,omitempty"`
}
