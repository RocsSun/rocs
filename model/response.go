package model

// UserInfo UserInfo的response返回结构体
type UserInfo struct {
	ID          int64    `json:"id"`
	Username    string   `json:"username"`
	Nickname    string   `json:"nickname"`
	Avatar      string   `json:"avatar"`
	Email       string   `json:"email"`
	Type        int      `json:"type"`
	Roles       []string `json:"roles"`
	Description string   `json:"description"`
	CreateTime  int64    `json:"createTime"`
}

// HasRole 判断是否包含有某个角色
func (u *UserInfo) HasRole(role string) bool {
	if len(u.Roles) == 0 {
		return false
	}
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}