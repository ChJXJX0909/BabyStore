// Code generated by goctl. DO NOT EDIT.
package types

type UserOptReq struct {
	Mobile string `json:"mobile"`
	Passwd string `json:"passwd"`
	Code   string `json:"code"`
}

type UserOptResp struct {
	Id    uint   `json:"id"`
	Token string `json:"token"`
}

type VerifyReq struct {
	Ticket string `json:"ticket"`
}

type VerifyResp struct {
	Data string `json:"data"`
}
