package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gvadmin_v3/core/baseapi"
	"gvadmin_v3/core/basedto"
	"gvadmin_v3/core/cache"
	"gvadmin_v3/core/global/E"
	"gvadmin_v3/core/log"
	"gvadmin_v3/core/util"
	"net/http"
	"regexp"
	"strings"
)

func Perm() func(c *gin.Context) {
	return func(c *gin.Context) {
		roleIds, _ := c.Get("roleId")
		rids := util.AnyToString(roleIds)

		if rids == "1" {
			//admin, 直接通过
			c.Set("dataScope", "1")
			c.Next()
		} else {
			dataScope, ok := roleAuthCheck(rids, c.Request.URL.Path)
			if !ok {
				c.JSON(http.StatusUnauthorized, baseapi.CommonResp{
					Code: http.StatusUnauthorized,
					Msg:  "权限不足",
				})
				c.Abort()
				return
			}
			c.Set("dataScope", dataScope)
			c.Next()
		}
	}
}

func roleAuthCheck(rids string, path string) (string, bool) {
	var tag = false
	var dataScope = "99"
	for _, rid := range strings.Split(rids, ",") {
		//验证api权限: 遍历roleKey, 在cache中检验角色是否具备接口权限
		roleAuthView, err := cache.Instance().Get(E.SystemRole, rid)
		if err != nil {
			log.Instance().Error("Get RoleAuthCache Failed..." + err.Error())
		}
		var roleAuth basedto.SysRoleAuthCacheView
		err = json.Unmarshal([]byte(roleAuthView), &roleAuth)
		if err != nil {
			log.Instance().Error("Unmarshal RoleAuth Failed..." + err.Error())
		}

		tmpP := strings.Split(path, "/")
		tmpPath := strings.Join(tmpP[:len(tmpP)-1], "/") //路由参数精确匹配
		if util.IsContain(roleAuth.AuthPath, path) || util.IsContain(roleAuth.AuthPath, tmpPath) {
			tag = true
			if util.AnyToInt(roleAuth.DataScope) < util.AnyToInt(dataScope) {
				dataScope = roleAuth.DataScope
			}
		}
	}

	return dataScope, tag
}

func checkAuthUrl(roleAuthUrl []string, path string) bool {
	for _, v := range roleAuthUrl {
		if strings.Contains(v, ":") {
			v = strings.Split(v, ":")[0]
		}

		b, err := regexp.MatchString(v, path)
		if err != nil {
			log.Instance().Error("Compare url failed: " + err.Error())
			break
		}
		if b {
			return true
		}
	}
	return false
}
