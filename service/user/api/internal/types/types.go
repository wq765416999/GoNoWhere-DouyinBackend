// Code generated by goctl. DO NOT EDIT.
package types

type Douyin_user_login_request struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type Douyin_user_login_response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int    `json:"user_id"`
	Token      string `json:"token"`
}

type Douyin_user_register_request struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type Douyin_user_register_response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int    `json:"user_id"`
	Token      string `json:"token"`
}

type Douyin_user_request struct {
	UserID int    `form:"user_id"`
	Token  string `form:"token"`
}

type Douyin_user_response struct {
	StatusCode int              `json:"status_code"`
	StatusMsg  string           `json:"status_msg"`
	User       Douyin_user_info `json:"user"`
}

type Douyin_user_info struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
