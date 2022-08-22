package workwx

// GetUser 读取成员
func (c *WorkwxApp) GetUser(userid string) (*UserInfo, error) {
	resp, err := c.execUserGet(reqUserGet{
		UserID: userid,
	})
	if err != nil {
		return nil, err
	}

	obj, err := resp.intoUserInfo()
	if err != nil {
		return nil, err
	}

	// TODO: return bare T instead of &T?
	return &obj, nil
}

// ListUsersByDeptID 获取部门成员详情
func (c *WorkwxApp) ListUsersByDeptID(deptID int64, fetchChild bool) ([]*UserInfo, error) {
	resp, err := c.execUserList(reqUserList{
		DeptID:     deptID,
		FetchChild: fetchChild,
	})
	if err != nil {
		return nil, err
	}
	users := make([]*UserInfo, len(resp.Users))
	for index, user := range resp.Users {
		userInfo, err := user.intoUserInfo()
		if err != nil {
			return nil, err
		}
		users[index] = &userInfo
	}
	return users, nil
}

// GetUserIDByMobile 通过手机号获取 userid
func (c *WorkwxApp) GetUserIDByMobile(mobile string) (string, error) {
	resp, err := c.execUserIDByMobile(reqUserIDByMobile{
		Mobile: mobile,
	})
	if err != nil {
		return "", err
	}
	return resp.UserID, nil
}

// GetUserInfoByCode 获取访问用户身份，根据code获取成员信息
func (c *WorkwxApp) GetUserInfoByCode(code string) (*UserIdentityInfo, error) {
	resp, err := c.execUserInfoGet(reqUserInfoGet{
		Code: code,
	})
	if err != nil {
		return nil, err
	}
	return &resp.UserIdentityInfo, nil
}
