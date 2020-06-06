package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
}

func GetStatus(code int) string {
	status, ok := MsgFlags[code]
	if ok {
		return status
	}
	return MsgFlags[ERROR]
}
