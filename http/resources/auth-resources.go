package resources

type AuthenticatedUserResource struct {
	Id    uint   `json:"id"`
	Token string `json:"token"`
}
