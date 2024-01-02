package workwx

// CreateKfAccount 创建客服账号
func (c *WorkwxApp) CreateKfAccount(name, mediaID string) (openKfID string, err error) {
	resp, err := c.execKfAccountCreate(reqKfAccountCreate{
		Name:    name,
		MediaID: mediaID,
	})
	if err != nil {
		return "", err
	}
	return resp.OpenKfID, nil
}

// DeleteKfAccount 删除客服账号
func (c *WorkwxApp) DeleteKfAccount(openKfID string) (err error) {
	_, err = c.execKfAccountDelete(reqKfAccountDelete{
		OpenKfID: openKfID,
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateKfAccount 修改客服账号
func (c *WorkwxApp) UpdateKfAccount(openKfID, name, mediaID string) (err error) {
	_, err = c.execKfAccountUpdate(reqKfAccountUpdate{
		OpenKfID: openKfID,
		Name:     name,
		MediaID:  mediaID,
	})
	if err != nil {
		return err
	}
	return nil
}

// ListKfAccount 获取客服账号列表
func (c *WorkwxApp) ListKfAccount(offset, limit int64) ([]*KfAccount, error) {
	resp, err := c.execKfAccountList(reqKfAccountList{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	return resp.AccountList, nil
}

// AddKfContact 获取客服账号链接
func (c *WorkwxApp) AddKfContact(openKfID, scene string) (url string, err error) {
	resp, err := c.execAddKfContact(reqAddKfContact{
		OpenKfID: openKfID,
		Scene:    scene,
	})
	if err != nil {
		return "", err
	}
	return resp.Url, nil
}
