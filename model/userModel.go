package model

const (
	UserStatusOK      = 0
	UserStatusDisable = 1

	UserTypeNormal = 0 // 普通用户
	UserTypeGzh    = 1 // 公众号用户
)

type Model struct {
	ID int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
}

type User struct {
	Model
	Username    string `gorm:"size:32;unique" json:"username" form:"username"`
	Nickname    string `gorm:"size:32;unique" json:"nickname" form:"nickname"`
	Avatar      string `gorm:"type:text" json:"avatar" form:"avatar"`
	Email       string `gorm:"size:512" json:"email" form:"email"`
	Password    string `gorm:"size:512" json:"password" form:"password"`
	Status      int    `gorm:"index:idx_status;not null" json:"status" form:"status"`
	Roles       string `gorm:"type:text" json:"roles" form:"roles"`
	Type        int    `gorm:"not null" json:"type" form:"type"`
	Description string `gorm:"type:text" json:"description" form:"description"`
	CreateTime  int64  `json:"createTime" form:"createTime"`
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`
}

type GithubUser struct {
	Model
	UserID     int64  `json:"userId" form:"userId"`
	GithubID   int64  `gorm:"unique" json:"githubId" form:"githubId"`
	Login      string `gorm:"size:512" json:"login" form:"login"`
	NodeId     string `gorm:"size:512" json:"nodeId" form:"nodeId"`
	AvatarUrl  string `gorm:"size:1024" json:"avatarUrl" form:"avatarUrl"`
	Url        string `gorm:"size:1024" json:"url" form:"url"`
	HtmlUrl    string `gorm:"size:1024" json:"htmlUrl" form:"htmlUrl"`
	Email      string `gorm:"size:512" json:"email" form:"email"`
	Name       string `gorm:"size:32" json:"name" form:"name"`
	CreateTime int64  `json:"createTime" form:"createTime"`
	UpdateTime int64  `json:"updateTime" form:"updateTime"`
}
