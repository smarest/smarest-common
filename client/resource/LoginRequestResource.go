package resource

type LoginRequestResource struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Cookie   string `json:"cookie"`
}
