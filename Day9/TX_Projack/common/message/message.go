package message

const (
	LoginMesType    = "LoginMes"
	LoginRseMesType = "LoginRseMes"
	RegisterMesType = "RegisterMes"
)

//定义两个消息
type Message struct {
	Type string //消息类型
	Data string //消息内容
}

type LoginMes struct {
	UserId   int    //用户ID
	UserPwd  string //用户密码
	UserName string //用户名
}
type LoginRseMes struct {
	Code  int    //状态码
	Error string //返回错误
}

type RegisterMes struct {
	//注册模块

}
