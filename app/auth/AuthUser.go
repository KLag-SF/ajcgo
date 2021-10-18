package auth

type AuthUser struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
