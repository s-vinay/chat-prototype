package utils

var (
	WS_USERNAME = "username"
	WS_TOKEN    = "token"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
