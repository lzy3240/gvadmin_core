package E

const (
	// MinGoVersion 最小 Go 版本
	MinGoVersion = "1.16"

	// ProjectVersion 项目版本
	ProjectVersion = "v1.0.0"

	// ProjectName 项目名称
	ProjectName = "gvadmin_v3"

	// ProjectDomain 项目域名
	ProjectDomain = "http://127.0.0.1"

	// HeaderSignToken 签名验证 Authorization，Header 中传递的参数
	HeaderSignToken = "Authorization"

	// HeaderSignTokenStr 签名验证 Authorization，Header 中传递的参数
	HeaderSignTokenStr = "Bearer "

	// 文件上传路径
	LocalFilePath = "./static/files/"
	ProfilePath   = "./static/images/"

	// 查看文件路径
	ShowFilePrefix  = "/storage/"
	ShowProfilePath = "/profile/"
	DefUploadSize   = 2 * 1024 * 1024 // 默认最大上传
	DefaultSaltLen  = 10

	//AllowAuth = "/system,/system/index,/system/main,/system/base,/system/menu" // 不需要验证的地址放在这里

	// 密码加盐
	Salt = "1pbxZyh" //用户密码加盐
	// $2a$10$5mK8zOSlHCyOBHvbyLOXbOMMr/vqtFmimshKbZBMDTFDLltVNFCxa
	// $2a$10$fJ4btCJMZrUKCg9.yiirXOlf1F7xwnpVj125rGL8uYGxLb8kDt3CO

	// 日期格式化
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"

	// 响应编码
	SUCCESS      = 200 // 成功
	ERROR        = 500 //错误
	UNAUTHORIZED = 401 //鉴权失败
	FORBIDDEN    = 403 //无操作权限
	FAIL         = -1  //失败

	// 操作类型 OperType
	OperTypeBase = 0 //其他
	OperTypeAdd  = 1 //新增
	OperTypeEdit = 2 //编辑
	OperTypeDel  = 3 //删除
	OperTypeView = 4 //查询

	// 操作名称 OperName

	// 消息主题
	TopicOperLog = "oper_log"

	// 缓存分区名称
	SystemBase   = "system_base"   //系统基础
	SystemConfig = "system_config" //系统配置
	SystemDict   = "system_dict"   //系统字典
	SystemMenu   = "system_menu"   //系统菜单
	SystemRole   = "system_role"   //系统角色权限
	SystemDept   = "system_dept"   //系统部门
	SystemUser   = "system_user"   //用户Token

	// 缓存公共KEY(SystemBase下的公共Key)
	MonitorKey = "monitor_key" //系统监控

	// 缓存码表公共KEY(SystemDice下的公共KEY)
	SystemArea        = "system_area"        //系统地区码表
	SystemIndustry    = "system_industry"    //系统行业码表
	SystemMajor       = "system_major"       //系统专业码表
	SystemTitle       = "system_title"       //系统职称码表
	SystemCredentials = "system_credentials" //系统职业资格码表

	// 缓存时间
	UserErrTimes  = 5     //密码错误5次
	UserLockTime  = 5     //锁定时间5分钟
	MonitorTime   = 60    //系统监控刷新时间
	MenuCacheTime = 86400 //系统菜单刷新时间
	RoleCacheTime = 86400 //系统角色权限刷新时间
)

var CacheNames = []string{
	SystemBase,
	SystemDict,
	SystemConfig,
	SystemMenu,
	SystemRole,
	SystemUser,
}
