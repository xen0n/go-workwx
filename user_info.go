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
