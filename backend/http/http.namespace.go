package http

var (
	JwtSecret = []byte("your-jwt-secret")
)

type User struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type Schedule struct {
	Day   string `bson:"day"`
	Start string `bson:"start"`
	End   string `bson:"end"`
}
