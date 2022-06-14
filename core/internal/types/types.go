// Code generated by goctl. DO NOT EDIT.
package types

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
