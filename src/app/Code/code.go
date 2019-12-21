package Code

const (
	SuccessCode = 200
	ErrorCode   = iota + 1
)

const (
	AuthenticatedCode = 401
)

const (
	UserNotExist = 404
)

var Message = map[int64]string{
	SuccessCode:       "success",
	ErrorCode:         "The server seems to have an error",
	AuthenticatedCode: "You have no permission to operate",
	UserNotExist:      "The user name does not exist",
}
