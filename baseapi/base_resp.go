package baseapi

import (
	"net/http"
)

// 通用响应信息
type CommonResp struct {
	Code      int         `json:"code"`           //响应编码: 200 成功 500 错误 403 无操作权限 401 鉴权失败  -1  失败
	Msg       string      `json:"msg"`            //消息内容
	Data      interface{} `json:"data,omitempty"` //数据内容
	Tag       bool        `json:"tag"`            //日志标识
	Type      int         `json:"type"`           //业务类型
	Name      string      `json:"name"`           //操作名称
	RequestId string      `json:"requestId"`      //请求编号
}

// 通用列表查询结果表单
type searchListResult struct {
	Total int64 `json:"total"`
	Rows  any   `json:"rows"`
}

// SuccessResp 返回一个成功的消息体
func (a *Api) SuccessResp() *Api {
	rid, _ := a.c.Get("requestId")
	resp := CommonResp{
		Code:      http.StatusOK,
		Msg:       "操作成功",
		RequestId: rid.(string),
	}

	a.r = &resp
	return a
}

// ErrorResp 返回一个错误的消息体
func (a *Api) ErrorResp() *Api {
	rid, _ := a.c.Get("requestId")
	resp := CommonResp{
		Code:      http.StatusInternalServerError,
		Msg:       "操作失败",
		RequestId: rid.(string),
	}

	a.r = &resp
	return a
}

// ForbiddenResp 返回一个拒绝访问的消息体
func (a *Api) ForbiddenResp() *Api {
	rid, _ := a.c.Get("requestId")
	resp := CommonResp{
		Code:      http.StatusForbidden,
		Msg:       "无操作权限",
		RequestId: rid.(string),
	}

	a.r = &resp
	return a
}

// UnauthorizedResp JWT认证失败
func (a *Api) UnauthorizedResp() *Api {
	rid, _ := a.c.Get("requestId")
	resp := CommonResp{
		Code:      http.StatusUnauthorized,
		Msg:       "鉴权失败",
		RequestId: rid.(string),
	}

	a.r = &resp
	return a
}

// SetMsg 设置消息体的内容
func (a *Api) SetMsg(msg string) *Api {
	a.r.Msg = msg
	return a
}

// SetCode 设置消息体的编码
func (a *Api) SetCode(code int) *Api {
	a.r.Code = code
	return a
}

// SetData 设置消息体的数据
func (a *Api) SetData(data interface{}) *Api {
	a.r.Data = data
	return a
}

// SetPageData 设置消息体的数据
func (a *Api) SetPageData(total int64, rows interface{}) *Api {
	a.r.Data = searchListResult{
		Total: total,
		Rows:  rows,
	}
	return a
}

// SetLogTag 设置日志标识
func (a *Api) SetLogTag(opType int, opName string) *Api {
	a.r.Tag = true
	a.r.Type = opType
	a.r.Name = opName
	return a
}

// WriteJsonExit 输出json到客户端
func (a *Api) WriteJsonExit() {
	a.c.Set("result", a.r)
	a.c.JSON(http.StatusOK, a.r)
	a.c.Abort()
}

// WriteErrJsonExit 输出json到客户端
func (a *Api) WriteErrJsonExit(errCode int) {
	a.c.Set("result", a.r)
	a.c.JSON(errCode, a.r)
	a.c.Abort()
}
