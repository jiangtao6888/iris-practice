package Code

const (
	SuccessCode = 200
	ErrorCode   = iota + 1
)

var Message = map[int64]string{
	SuccessCode: "success",
	ErrorCode:   "The server seems to have an error",
}
