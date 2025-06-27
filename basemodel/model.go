package basemodel

import (
	"gvadmin_core/util"
	"time"
)

type Model struct {
	CreateBy string    `json:"createBy" gorm:"create_by;size:64;comment:创建人"`
	CreateAt time.Time `json:"createAt" gorm:"column:create_at;type:datetime;default:CURRENT_TIMESTAMP;"`
	UpdateBy string    `json:"updateBy" gorm:"update_by;size:64;comment:操作人"`
	UpdateAt time.Time `json:"updateAt" gorm:"column:update_at;type:datetime;default:CURRENT_TIMESTAMP on update current_timestamp;"`
}

func (m *Model) SetCreate(userId int) {
	m.CreateBy = util.AnyToString(userId)
	m.CreateAt = time.Now()
}

func (m *Model) SetUpdate(userId int) {
	m.UpdateBy = util.AnyToString(userId)
	m.UpdateAt = time.Now()
}
