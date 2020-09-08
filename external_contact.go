package workwx

// ListExternalContact 获取客户列表
func (c *WorkwxApp) ListExternalContact(userID string) ([]string, error) {
	resp, err := c.execExternalContactList(reqExternalContactList{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExternalUserID, nil
}

// GetExternalContact 获取客户详情
func (c *WorkwxApp) GetExternalContact(externalUserid string) (*ExternalContactInfo, error) {
	resp, err := c.execExternalContactGet(reqExternalContactGet{
		ExternalUserID: externalUserid,
	})
	if err != nil {
		return nil, err
	}
	return &resp.ExternalContactInfo, nil
}

// RemarkExternalContact 修改客户备注信息
func (c *WorkwxApp) RemarkExternalContact(req *ExternalContactRemark) error {
	_, err := c.execExternalContactRemark(reqExternalContactRemark{
		Remark: req,
	})
	return err
}
