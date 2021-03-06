// Code generated by goctl. DO NOT EDIT.
package types

type Result struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type LoginRequest struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type DetailRequest struct {
	Identity string `json:"identity"`
}

type DetailResponse struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type RefreshAuthorizationReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code,optional"`
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type FileUploadReply struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}
