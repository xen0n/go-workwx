package workwx

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"testing"
)

var (
	wxclient *WorkwxApp
	// 企业ID
	corpID string = "xxxxxxx"
	// 应用ID
	agentID int64 = 007
	//  应用秘钥
	agentSecret string = "xxxxxxxxxx"
	// 测试接收消息的用户ID
	userID string = "userid"
)

func init() {
	var wx = New(corpID)
	wxclient = wx.WithApp(agentSecret, agentID)
	wxclient.SpawnAccessTokenRefresher() // 自动刷新token
}

func TestSendTextMessage(t *testing.T) {
	recipient := Recipient{
		UserIDs:  []string{userID},
		PartyIDs: []string{},
		TagIDs:   []string{},
		ChatID:   "",
	}
	_ = wxclient.SendTextMessage(&recipient, "这是一条普通文本消息", false)
}

func TestSendMarkdownMessage(t *testing.T) {
	recipient := Recipient{
		UserIDs:  []string{userID},
		PartyIDs: []string{},
		TagIDs:   []string{},
		ChatID:   "",
	}
	_ = wxclient.SendMarkdownMessage(&recipient, "您的会议室已经预定，稍后会同步到`邮箱`  \n>**事项详情**  \n>事　项：<font color=\"info\">开会</font>  \n>组织者：@miglioguan  \n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang  \n>  \n>会议室：<font color=\"info\">广州TIT 1楼 301</font>  \n>日　期：<font color=\"warning\">2018年5月18日</font>  \n>时　间：<font color=\"comment\">上午9:00-11:00</font>  \n>  \n>请准时参加会议。  \n>  \n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)", false)
}

func TestSendImageMessage(t *testing.T) {
	recipient := Recipient{
		UserIDs:  []string{userID},
		PartyIDs: []string{},
		TagIDs:   []string{},
		ChatID:   "",
	}
	url := "https://wwcdn.weixin.qq.com/node/wework/images/202201062104.366e5ee28e.png"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP GET请求失败:", err)
		return
	}
	defer resp.Body.Close()
	reader := resp.Body

	rst, err := wxclient.UploadTempImageMedia(&Media{
		filename: "test.jpg",
		filesize: 0,
		stream:   reader,
	})
	if err != nil {
		fmt.Printf("upload temp image failed: %v\n", err)
	}
	_ = wxclient.SendImageMessage(&recipient, rst.MediaID, false)
}

func TestSendFileMessage(t *testing.T) {
	recipient := Recipient{
		UserIDs:  []string{userID},
		PartyIDs: []string{},
		TagIDs:   []string{},
		ChatID:   "",
	}
	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	rst, err := wxclient.UploadTempFileMedia(&Media{
		filename: "go.mod",
		filesize: 0,
		stream:   reader,
	})
	if err != nil {
		fmt.Printf("upload temp file failed: %v\n", err)
	}
	_ = wxclient.SendFileMessage(&recipient, rst.MediaID, false)
}

func TestSendTextCardMessage(t *testing.T) {
	recipient := Recipient{
		UserIDs:  []string{userID},
		PartyIDs: []string{},
		TagIDs:   []string{},
		ChatID:   "",
	}
	_ = wxclient.SendTextCardMessage(
		&recipient, "领奖通知", "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>", "https://wiki.eryajf.net", "更多", false)
}

func TestSendNewsMessage(t *testing.T) {
	recipient := Recipient{
		UserIDs:  []string{userID},
		PartyIDs: []string{},
		TagIDs:   []string{},
		ChatID:   "",
	}

	msgObj := []Article{
		{Title: "中秋节礼品领取",
			Description: "今年中秋节公司有豪礼相送",
			Url:         "https://wiki.eryajf.net",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
			AppId:       "",
			PagePath:    ""},
		{Title: "中秋节礼品领取2",
			Description: "今年中秋节公司有豪礼相送2",
			Url:         "https://wiki.eryajf.net",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
			AppId:       "",
			PagePath:    ""},
	}

	_ = wxclient.SendNewsMessage(&recipient, msgObj, false)
}
func TestSendMPNewsMessage(t *testing.T) {
	recipient := Recipient{
		UserIDs:  []string{userID},
		PartyIDs: []string{},
		TagIDs:   []string{},
		ChatID:   "",
	}

	url := "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP GET请求失败:", err)
		return
	}
	defer resp.Body.Close()
	reader := resp.Body

	rst, err := wxclient.UploadTempImageMedia(&Media{
		filename: "test.jpg",
		filesize: 0,
		stream:   reader,
	})
	if err != nil {
		fmt.Printf("upload temp image failed: %v\n", err)
	}

	msgObj := []MPArticle{
		{Title: "中秋节礼品领取",
			ThumbMediaID:     rst.MediaID,
			Author:           "eryajf",
			ContentSourceUrl: "https://wiki.eryajf.net",
			Content:          "这是正文里边的内容。",
			Digest:           "这里是图文消息的描述"},
	}

	_ = wxclient.SendMPNewsMessage(&recipient, msgObj, false)
}

func TestSendTaskCardMessage(t *testing.T) {
	recipient := Recipient{
		UserIDs:  []string{userID},
		PartyIDs: []string{},
		TagIDs:   []string{},
		ChatID:   "",
	}
	btn := []TaskCardBtn{
		{
			Key:         "yes",
			Name:        "通过",
			ReplaceName: "已通过",
			Color:       "blue",
			IsBold:      false,
		}, {
			Key:         "no",
			Name:        "拒绝",
			ReplaceName: "已拒绝",
			Color:       "red",
			IsBold:      false,
		}}
	_ = wxclient.SendTaskCardMessage(&recipient, "请审核该条信息", "这是说明信息", "https://wiki.eryajf.net", "aaab", btn, false)
}

func TestSendTemplateCardMessage(t *testing.T) {
	recipient := Recipient{
		UserIDs:  []string{userID},
		PartyIDs: []string{},
		TagIDs:   []string{},
		ChatID:   "",
	}
	msgObj := TemplateCard{
		CardType: CardTypeTextNotice,
		Source: Source{
			IconURL:   "https://t.eryajf.net/imgs/2023/02/712e2287455b9a0c.png",
			Desc:      "二丫讲梵的公众号",
			DescColor: 0,
		},
		ActionMenu: &ActionMenu{
			Desc: "卡片副交互辅助文本说明",
			ActionList: []ActionList{
				{Text: "接受推送", Key: "A"},
				{Text: "不再推送", Key: "B"},
			},
		},
		TaskID: "aaadaa",
		MainTitle: MainTitle{
			Title: "欢迎使用企业微信",
			Desc:  "你的朋友也都在用。",
		},
		QuoteArea: QuoteArea{
			Type:      0,
			URL:       "baidu.com",
			Title:     "百度",
			QuoteText: "去往百度",
		},
		EmphasisContent: &EmphasisContent{
			Title: "100",
			Desc:  "核心数据",
		},
		SubTitleText: "下载企业微信还能抢红包！",
		CardAction: CardAction{
			Type:     1,
			URL:      "qq.com",
			Appid:    "aaaaaaa",
			Pagepath: "/index.html",
		},
	}
	err := wxclient.SendTemplateCardMessage(&recipient, msgObj, false)
	if err != nil {
		fmt.Printf("get err: %v\n", err)
	}
}
