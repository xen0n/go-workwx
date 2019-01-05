package workwx

// UserInfo 用户信息
type UserInfo struct {
	// UserID 成员UserID
	// 对应管理端的账号，企业内必须唯一。不区分大小写，长度为1~64个字节
	UserID string
	// Name 成员名称
	Name string
	// Position 职务信息；第三方仅通讯录应用可获取
	Position string
	// Departments 成员所属部门信息
	Departments []UserDeptInfo
	// Mobile 手机号码；第三方仅通讯录应用可获取
	Mobile string
	// Gender 性别
	Gender UserGender
	// 邮箱；第三方仅通讯录应用可获取
	Email string
	// AvatarURL 头像 URL；第三方仅通讯录应用可获取
	//
	// NOTE：如果要获取小图将url最后的”/0”改成”/100”即可。
	AvatarURL string
	// Telephone 座机；第三方仅通讯录应用可获取
	Telephone string
	// IsEnabled 成员的启用状态
	IsEnabled bool
	// Alias 别名；第三方仅通讯录应用可获取
	Alias string
	// Status 成员激活状态
	Status UserStatus
	// QRCodeURL 员工个人二维码；第三方仅通讯录应用可获取
	//
	// 扫描可添加为外部联系人
	QRCodeURL string

	// TODO: extattr external_profile external_position
}

// UserDeptInfo 用户部门信息
type UserDeptInfo struct {
	DeptID int64
	// Order 部门内的排序值，默认为0，数值越大排序越前面
	Order uint32
	// IsLeader 在所在的部门内是否为上级
	IsLeader bool
}

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
