package basedto

// 树形结构
type SysCommonTree struct {
	Id       int             `json:"id"`
	Label    string          `json:"label"`              /** 节点名称 */
	Children []SysCommonTree `json:"children,omitempty"` /** 子节点 */
}

// 分页参数
type PageParams struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// 时间参数, 暂未使用
type TimeParams struct {
	BeginTime string `json:"beginTime" form:"beginTime" search:"type:gte"`
	EndTime   string `json:"endTime" form:"beginTime" search:"type:lte"`
}

// 排序参数
type OrderParams struct {
	OrderByColumn string `json:"orderByColumn" form:"orderByColumn"`
	IsAsc         string `json:"isAsc" form:"isAsc"`
}

// 控制参数, 不使用, 使用base_model
//type ControlBy struct {
//	CreateBy string    `json:"createBy"`
//	CreateAt time.Time `json:"createAt"`
//	UpdateBy string    `json:"updateBy"`
//	UpdateAt time.Time `json:"updateAt"`
//}
//
//func (m *ControlBy) SetCreate(userName string) {
//	m.CreateBy = userName
//	m.CreateAt = time.Now()
//}
//
//func (m *ControlBy) SetUpdate(userName string) {
//	m.UpdateBy = userName
//	m.UpdateAt = time.Now()
//}

type SysRoleAuthCacheView struct {
	RoleKey   string   `json:"roleKey"`
	DataScope string   `json:"dataScope"`
	AuthPath  []string `json:"authPath"`
}
