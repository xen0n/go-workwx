package workwx

// DeptInfo 部门信息
type DeptInfo struct {
	// ID 部门 ID
	ID int64
	// Name 成员名称
	Name string
	// ParentID 父亲部门id。根部门为1
	ParentID int64
	// Order 在父部门中的次序值。order值大的排序靠前。值范围是[0, 2^32)
	Order uint32
}

//
//
//

const deptListEndpoint = "/cgi-bin/department/list"

// ListAllDepts 获取全量组织架构。
func (c *WorkwxApp) ListAllDepts() ([]*DeptInfo, error) {
	return c.listDepts(false, 0)
}

// ListDepts 获取指定部门及其下的子部门。
func (c *WorkwxApp) ListDepts(id int64) ([]*DeptInfo, error) {
	return c.listDepts(true, id)
}

func (c *WorkwxApp) listDepts(haveID bool, id int64) ([]*DeptInfo, error) {
	req := reqDeptList{
		HaveID: haveID,
		ID:     id,
	}

	var resp respDeptList
	err := c.executeQyapiGet(deptListEndpoint, req, &resp, true)
	if err != nil {
		// TODO: error_chain
		return nil, err
	}

	obj := resp.Department
	return obj, nil
}
