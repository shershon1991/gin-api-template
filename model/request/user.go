/**
 * @Description TODO
 **/
package request

// 登录参数
type LoginParam struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// 注册参数
type RegisterParam struct {
	NickName string `json:"nickName"`
	Birthday string `json:"birthday"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
