package types

type JWTCreator interface {
	CreateJWTtoken(sub string) (string, error)
}

type SubGetter interface {
	GetSubjectFromJWT(token string) (string, error)
}
