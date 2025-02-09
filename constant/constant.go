package constant

type ResponseCode int

const (
	Fail      ResponseCode = 0 // 失败
	NeedLogin ResponseCode = 1 // 需要登录
	Success   ResponseCode = 100 // 成功
)
