package wecom

import (
	ww "github.com/dfang/wechat-work-go"
	"github.com/dfang/wechat-work-go/message"
)

func NewWecom(corpId, corpSecret string, agentId int64) *message.Message {
	corp := ww.New(corpId)
	logApp := corp.NewApp(corpSecret, agentId)
	return message.WithApp(logApp)
}
