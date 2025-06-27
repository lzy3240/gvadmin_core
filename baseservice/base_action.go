package baseservice

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gvadmin_v3/core/baseservice/search"
	"gvadmin_v3/core/config"
	"reflect"
)

// SetCondition 设置查询条件
func (s *Service) SetCondition(q interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		condition := &search.GormCondition{
			GormPublic: search.GormPublic{},
			Join:       make([]*search.GormJoin, 0),
		}
		search.ResolveSearchQuery(q, condition)
		for _, join := range condition.Join {
			if join == nil {
				continue
			}
			db = db.Joins(join.JoinOn)
			for k, v := range join.Where {
				db = db.Where(k, v...)
			}
			for k, v := range join.Or {
				db = db.Or(k, v...)
			}
			for _, o := range join.Order {
				db = db.Order(o)
			}
		}
		for k, v := range condition.Where {
			db = db.Where(k, v...)
		}
		for k, v := range condition.Or {
			db = db.Or(k, v...)
		}
		for _, o := range condition.Order {
			db = db.Order(o)
		}
		return db
	}
}

// SetOrder 设置排序条件
func (s *Service) SetOrder(orderByColumn string, isAsc string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if orderByColumn == "" {
			orderByColumn = "id"
		}

		var desc bool
		switch isAsc {
		case "ascending":
			desc = false
		case "descending":
			desc = true
		default:
			desc = true
		}

		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: orderKey[orderByColumn]}, Desc: desc})
	}
}

// SetPaginate 设置分页条件
func (s *Service) SetPaginate(pageSize, pageNum int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 设置默认值
		if pageSize == 0 {
			pageSize = 10
		}

		if pageNum == 0 {
			pageNum = 1
		}

		offset := (pageNum - 1) * pageSize
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(pageSize)
	}
}

// SetDataPerm 设置数据权限
func (s *Service) SetDataPerm(tableName string, p *DataPermission) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if config.Instance().App.EnableDP == 0 {
			return db
		}
		switch p.DataScope {
		case "2":
			return db.Where(tableName+".create_by in (select sys_user.user_id from sys_role_dept left join sys_user on sys_user.dept_id=sys_role_dept.dept_id where sys_role_dept.role_id = ?)", p.RoleId)
		case "3":
			return db.Where(tableName+".create_by in (SELECT user_id from sys_user where dept_id = ? )", p.DeptId)
		case "4":
			return db.Where(tableName+".create_by in (SELECT user_id from sys_user where sys_user.dept_id in(select dept_id from sys_dept where dept_id = ? union select dept_id from sys_dept where parent_id = ?))", p.DeptId, p.DeptId)
		case "5":
			return db.Where(tableName+".create_by = ?", p.UserId)
		default:
			return db
		}
	}
}

// StructToMapByKey gorm更新方法updates忽略零值, 故对象转换为map更新, 前提是请求参数标记key
func (s *Service) StructToMapByKey(obj interface{}, key string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get(key) == "" {
			continue
		}
		data[t.Field(i).Tag.Get(key)] = v.Field(i).Interface()
	}
	return data
}
