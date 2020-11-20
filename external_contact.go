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

// ListExternalContactCorpTags 获取企业标签库
func (c *WorkwxApp) ListExternalContactCorpTags(tagIDs ...string) ([]ExternalContactCorpTagGroup, error) {
	resp, err := c.execExternalContactListCorpTags(reqExternalContactListCorpTags{
		TagIDs: tagIDs,
	})
	if err != nil {
		return nil, err
	}
	return resp.TagGroup, nil
}

// AddExternalContactCorpTag 添加企业客户标签
func (c *WorkwxApp) AddExternalContactCorpTag(req ExternalContactCorpTagGroup) ([]ExternalContactCorpTagGroup, error) {
	resp, err := c.execExternalContactAddCorpTag(reqExternalContactAddCorpTag{
		ExternalContactCorpTagGroup: req,
	})
	if err != nil {
		return nil, err
	}
	return resp.TagGroup, nil
}

// EditExternalContactCorpTag 编辑企业客户标签
func (c *WorkwxApp) EditExternalContactCorpTag(id, name string, order uint32) error {
	_, err := c.execExternalContactEditCorpTag(reqExternalContactEditCorpTag{
		ID:    id,
		Name:  name,
		Order: order,
	})
	return err
}

// DelExternalContactCorpTag 删除企业客户标签
func (c *WorkwxApp) DelExternalContactCorpTag(tagID, groupID []string) error {
	_, err := c.execExternalContactDelCorpTag(reqExternalContactDelCorpTag{
		TagID:   tagID,
		GroupID: groupID,
	})
	return err
}

// MarkExternalContactTag 标记客户企业标签
func (c *WorkwxApp) MarkExternalContactTag(userID, externalUserID string, addTag, removeTag []string) error {
	_, err := c.execExternalContactMarkTag(reqExternalContactMarkTag{
		UserID:         userID,
		ExternalUserID: externalUserID,
		AddTag:         addTag,
		RemoveTag:      removeTag,
	})
	return err
}
