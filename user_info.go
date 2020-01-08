package workwx

// GetUser 读取成员
func (c *WorkwxApp) GetUser(userid string) (*UserInfo, error) {
	resp, err := c.execUserGet(reqUserGet{
		UserID: userid,
	})
	if err != nil {
		return nil, err
	}

	// TODO: return bare T instead of &T?
	obj := resp.intoUserInfo()
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
		userInfo := user.intoUserInfo()
		users[index] = &userInfo
	}
	return users, nil
}
