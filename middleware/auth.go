package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gvadmin_core/baseapi"
	"gvadmin_core/config"
	"gvadmin_core/global/E"
	"net/http"
	"strings"
	"time"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 根据实际情况取TOKEN, 这里从request header取
		header := c.Request.Header
		tokenStr := header.Get(E.HeaderSignToken)
		if len(tokenStr) < 1 {
			c.JSON(http.StatusInternalServerError, baseapi.CommonResp{
				Code: http.StatusInternalServerError,
				Msg:  "参数错误",
			})
			c.Abort()
			return
		}

		token, err := VerifyToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, baseapi.CommonResp{
				Code: http.StatusUnauthorized,
				Msg:  "认证失败",
			})
			c.Abort()
			return
		}
		userId := token.Claims.(jwt.MapClaims)["user_id"]
		deptId := token.Claims.(jwt.MapClaims)["dept_id"]
		roleId := token.Claims.(jwt.MapClaims)["role_id"]

		// 此处已经通过了, 可以把Claims中的有效信息拿出来放入上下文使用
		c.Set("userId", userId)
		c.Set("deptId", deptId)
		c.Set("roleId", roleId)
		c.Next()
	}
}

func CreateToken(UserId int, DeptId int, RoleIds string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": UserId,
		"dept_id": DeptId,
		"role_id": RoleIds,
		"exp":     time.Now().Unix() + int64(config.Instance().Jwt.Ttl),
		"iss":     "gvadmin_v3",
	})

	mySigningKey := []byte(config.Instance().Jwt.Secret)
	//token加密
	//TODO
	return token.SignedString(mySigningKey)
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	//token解密
	//TODO
	mySigningKey := []byte(config.Instance().Jwt.Secret)
	tokenStr = strings.ReplaceAll(tokenStr, E.HeaderSignTokenStr, "")
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
}
