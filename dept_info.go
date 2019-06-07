package workwx

// ListAllDepts 获取全量组织架构。
func (c *WorkwxApp) ListAllDepts() ([]*DeptInfo, error) {
	resp, err := c.execDeptList(reqDeptList{
		HaveID: false,
		ID:     0,
	})
	if err != nil {
		return nil, err
	}

	return resp.Department, nil
}

// ListDepts 获取指定部门及其下的子部门。
func (c *WorkwxApp) ListDepts(id int64) ([]*DeptInfo, error) {
	resp, err := c.execDeptList(reqDeptList{
		HaveID: true,
		ID:     id,
	})
	if err != nil {
		return nil, err
	}

	return resp.Department, nil
}
