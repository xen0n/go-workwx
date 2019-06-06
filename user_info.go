package workwx

// UserGender 用户性别
type UserGender int

const (
	// UserGenderUnspecified 性别未定义
	UserGenderUnspecified UserGender = 0
	// UserGenderMale 男性
	UserGenderMale UserGender = 1
	// UserGenderFemale 女性
	UserGenderFemale UserGender = 2
)

// UserStatus 用户激活信息
//
// 已激活代表已激活企业微信或已关注微工作台（原企业号）。
// 未激活代表既未激活企业微信又未关注微工作台（原企业号）。
type UserStatus int

const (
	// UserStatusActivated 已激活
	UserStatusActivated UserStatus = 1
	// UserStatusDeactivated 已禁用
	UserStatusDeactivated UserStatus = 2
	// UserStatusUnactivated 未激活
	UserStatusUnactivated UserStatus = 4
)

//
//
//

const userGetEndpoint = "/cgi-bin/user/get"

// GetUser 读取成员
func (c *WorkwxApp) GetUser(userid string) (*UserInfo, error) {
	req := reqUserGet{
		UserID: userid,
	}

	var resp respUserGet
	err := c.executeQyapiGet(userGetEndpoint, req, &resp, true)
	if err != nil {
		// TODO: error_chain
		return nil, err
	}

	// TODO: return bare T instead of &T?
	obj := resp.intoUserInfo()
	return &obj, nil
}
