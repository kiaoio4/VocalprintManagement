package enums

type JsonResultCode int

const (
	JRCodeSucc         JsonResultCode = 200
	JRCodeFailed                      = 600
	JRCodeError                       = 500
	JRCodeWarn                        = 501
	JRCodeRequestError                = 400
	JRCode302                         = 302 //跳转至地址
	JRCode401                         = 401 //未授权访问
)

const (
	Deleted  = 1
	Disabled = 0
	Enabled  = 1
)
