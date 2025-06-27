package baseservice

type Service struct {
	RequestID string
	DP        DataPermission
}

type DataPermission struct {
	DataScope string
	UserId    int
	DeptId    int
	RoleId    int
}

var orderKey = map[string]string{
	"id":        "id",
	"operTime":  "oper_time",
	"userName":  "user_name",
	"loginTime": "login_time",
	"operName":  "oper_name",
}
