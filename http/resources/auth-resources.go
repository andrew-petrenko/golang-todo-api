package resources

type AuthenticatedUser struct {
	Id    uint   `json:"id"`
	Token string `json:"token"`
}
