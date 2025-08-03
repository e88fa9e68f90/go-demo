package remote

import (
	"bytes"
	"github.com/gin-gonic/gin"
	loggerop "github.com/kiwi633/go-demo/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/http/httputil"
)

func GetPolicyList(c *gin.Context) {
	client := &http.Client{}
	url := "https://pensionlife.95522.cn/gtw/policy-web/policy/v2/get-policy-list.json?clue-id=892cfa48408172ac0080962f4fa4666e"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`["1"]`)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJkNTIzODYzZC0zODg4LTQzNmQtYjJjZC0zM2Y1MzI5YTBlYmYiLCJ1c2VySWQiOiIxNzUyOTIwODEzNTMzNzYxNTM3IiwibmFtZSI6IuWtmemAmiIsInZlcnNpb24iOjAsInVzZXJUeXBlIjoxLCJleHBpcmUiOjc3NzYwMDAsInBsYXlsb2FkIjoie1wiY3VzdG9tVHlwZVwiOlwiMVwiLFwiYWNjb3VudE5hbWVcIjpcImY0NjI4NTRiLTgyMGQtNDljNy1hOGQyLWNjYjBhMzNiZWFlNlwiLFwidXNlck5hbWVcIjpcImY0NjI4NTRiLTgyMGQtNDljNy1hOGQyLWNjYjBhMzNiZWFlNlwiLFwidG9rZW5cIjpcImY0NjI4NTRiLTgyMGQtNDljNy1hOGQyLWNjYjBhMzNiZWFlNlwifSIsImV4cCI6MTc2MTg5MzYyOH0.Li_OD5xD7waHHg-_JV5aES2mmQhnbsueQiLFR1gicr5bJ2X4rd2cJ3kesJPQ3iAvLP1_2rvDjvf4a2W1KBBYIJTTrPA9g9PkldeSZn1gwpvdno8Pw8OtDAbrZldTuB2BusYBRksv60XNOzxei128Gvu44VyP_8QfssTmvX8dmCc")
	req.Header.Set("Content-Type", "application/json")
	// 打印请求数据
	bytessss, err := httputil.DumpRequest(req, true)
	if err != nil {
		panic(err)
	}
	loggerop.Info(string(bytessss))
	resp, err := client.Do(req)
	respionse, _ := httputil.DumpResponse(resp, true)
	loggerop.Info(string(respionse))

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	loggerop.Info("获取保单列表打印返回参数：", zap.String("body", string(body)))
	c.JSON(http.StatusOK, gin.H{"message": "OK!"})
}
