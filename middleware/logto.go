package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gvadmin_core/baseapi"
	"gvadmin_core/global/E"
	"gvadmin_core/log"
	"gvadmin_core/queue"
	"gvadmin_core/util"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func LogTo() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 获取请求参数
		var param string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			bf := bytes.NewBuffer(nil)
			wt := bufio.NewWriter(bf)
			_, err := io.Copy(wt, c.Request.Body)
			if err != nil {
				log.Instance().Error("copy request body Failed." + err.Error())
				err = nil
			}
			rb, _ := ioutil.ReadAll(bf)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
			param = string(rb)
		}

		// 继续执行
		c.Next()

		// 结束时间
		endTime := time.Now()

		// OPTION不记录
		if c.Request.Method == http.MethodOptions {
			return
		}

		// 接收返回结果
		outBody, _ := c.Get("result")
		respBody := outBody.(*baseapi.CommonResp)

		// Tag为false不记录
		if respBody.Tag == false {
			return
		}

		// 封装日志信息
		userId, _ := c.Get("userId")
		uid := util.AnyToInt(userId)
		deptId, _ := c.Get("deptId")
		did := util.AnyToInt(deptId)

		// 构造日志map
		operLog := make(map[string]interface{})
		operLog["user_id"] = uid
		operLog["request_method"] = c.Request.Method
		operLog["operator_type"] = util.GetDevice(c.Request) //访问渠道 0-爬虫 1-Mobil 2-PC
		operLog["dept_id"] = did
		operLog["oper_url"] = c.Request.URL.Path
		operLog["oper_ip"] = util.GetClientIp(c.Request)
		operLog["oper_location"] = ""
		operLog["oper_param"] = param
		operLog["oper_time"] = time.Now()
		operLog["latency_time"] = strconv.FormatInt(endTime.Sub(startTime).Milliseconds(), 10)
		operLog["user_agent"] = c.Request.Header.Get("User-Agent")

		// 处理返回结果, 暂记录256
		respData := util.AnyToString(respBody.Data)
		if len(respData) > 255 {
			operLog["json_result"] = respData[0:255]
		} else {
			operLog["json_result"] = respData
		}

		if respBody.Code == 500 {
			operLog["status"] = "1"
		} else {
			operLog["status"] = "0"
		}

		operLog["request_id"] = respBody.RequestId
		operLog["business_type"] = util.AnyToString(respBody.Type)
		operLog["business_name"] = respBody.Name
		operLog["msg"] = respBody.Msg

		//写入异步队列
		msg, err := json.Marshal(operLog)
		if err != nil {
			log.Instance().Error("Marshal OperLog Failed..." + err.Error())
		}

		err = queue.Instance().Publish(E.TopicOperLog, string(msg))
		if err != nil {
			log.Instance().Error("Push Msg Failed..." + err.Error())
		}
	}
}
