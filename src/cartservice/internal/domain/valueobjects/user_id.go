package valueobjects

type UserID string

func ParseUserID(id string) UserID {
	return UserID(id)
}

func (u UserID) String() string {
	return u.String()
}
