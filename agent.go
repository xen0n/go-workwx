package workwx

// AgentInfo 应用信息
type AgentInfo struct {
	// AgentID 企业应用id
	AgentID int64
	// Name 企业应用名称
	Name string
	// SquareLogoURL 企业应用方形头像
	SquareLogoURL string
	// Description 企业应用详情
	Description string
	// AllowUserInfos 企业应用可见范围（人员）
	AllowUserInfos AgentAllowUserInfos
	// AllowPartys 企业应用可见范围（部门）
	AllowPartys AgentAllowPartys
	// AllowTags 企业应用可见范围（标签）
	AllowTags AgentAllowTags
	// Close 企业应用是否被停用。0：未被停用；1：被停用
	Close int
	// RedirectDomain 企业应用可信域名
	RedirectDomain string
	// ReportLocationFlag 企业应用是否打开地理位置上报 0：不上报；1：进入会话上报
	ReportLocationFlag int
	// IsReportEnter 是否上报用户进入应用事件。0：不接收；1：接收
	IsReportEnter int
	// HomeURL 应用主页url
	HomeURL string
	// CustomizedPublishStatus 代开发自建应用返回该字段，表示代开发发布状态。
	// 0：待开发（企业已授权，服务商未创建应用）
	// 1：开发中（服务商已创建应用，未上线）
	// 2：已上线（服务商已上线应用且不存在未上线版本）
	// 3：存在未上线版本（服务商已上线应用但存在未上线版本）
	CustomizedPublishStatus int
}

// GetAgent 获取指定的应用详情
func (c *WorkwxApp) GetAgent(agentID int64) (*AgentInfo, error) {
	resp, err := c.execAgentGet(reqAgentGet{
		AgentID: agentID,
	})
	if err != nil {
		return nil, err
	}

	agentInfo := &AgentInfo{
		AgentID:                 resp.AgentID,
		Name:                    resp.Name,
		SquareLogoURL:           resp.SquareLogoURL,
		Description:             resp.Description,
		AllowUserInfos:          resp.AllowUserInfos,
		AllowPartys:             resp.AllowPartys,
		AllowTags:               resp.AllowTags,
		Close:                   resp.Close,
		RedirectDomain:          resp.RedirectDomain,
		ReportLocationFlag:      resp.ReportLocationFlag,
		IsReportEnter:           resp.IsReportEnter,
		HomeURL:                 resp.HomeURL,
		CustomizedPublishStatus: resp.CustomizedPublishStatus,
	}

	return agentInfo, nil
}
