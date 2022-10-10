package workwx

// CreateAppchat 创建群聊会话
func (c *WorkwxApp) CreateAppchat(chatInfo *ChatInfo) (chatid string, err error) {
	resp, err := c.execAppchatCreate(reqAppchatCreate{
		ChatInfo: chatInfo,
	})
	if err != nil {
		return "", err
	}
	return resp.ChatID, nil
}

// GetAppchat 获取群聊会话
func (c *WorkwxApp) GetAppchat(chatid string) (*ChatInfo, error) {
	resp, err := c.execAppchatGet(reqAppchatGet{
		ChatID: chatid,
	})
	if err != nil {
		return nil, err
	}

	// TODO: return bare T instead of &T?
	obj := resp.ChatInfo
	return obj, nil
}

// GetAppChatList 获取群聊列表
func (c *WorkwxApp) GetAppChatList(req ReqChatList) (*RespAppchatList, error) {
	resp, err := c.execAppchatListGet(reqAppchatList{
		ReqChatList: req,
	})
	if err != nil {
		return nil, err
	}
	return resp.RespAppchatList, nil
}

func (c *WorkwxApp) GetAppChatInfo(chatId string) (*RespAppChatInfo, error) {
	resp, err := c.execAppchatInfoGet(reqAppchatInfo{
		ChatId:   chatId,
		NeedName: ChatNeedName,
	})
	if err != nil {
		return nil, err
	}
	return resp.GroupChat, nil
}
