package log

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dfang/wechat-work-go/message"
	"github.com/joker8023/log/wecom"
)

const (
	red    uint8 = 91
	yellow uint8 = 93
	blue   uint8 = 94
)

func Setcolor(s string, color uint8) string {
	return fmt.Sprintf("\x1b[%dm[%s]\x1b[0m ", color, s)
}

type Config struct {
	Project     string
	Environment string
}

type WecomClient struct {
	CorpId     string
	CorpSecret string
	AgentId    int64
	ToUser     string
}
type Log struct {
	wecomClient *WecomClient
	config      *Config
	prefix      string
}

type LogPrefix struct {
	Log
}

func NewLog(prefix string, config *Config) (l *Log) {
	l = &Log{prefix: prefix, config: config}
	return l
}

func (l *Log) UseWecom(wecomClient *WecomClient) {
	l.wecomClient = wecomClient
}

func (l *Log) SendMessage(level, info string) {
	if l.wecomClient != nil {
		digest := fmt.Sprintf("项目%s的%s环境有一个%s级别的日志", l.config.Project, l.config.Environment, level)
		content := fmt.Sprintf(`
		<div style="color:red;font-size:16px;">%s:%s</div>
		<div style="color:red;font-size:16px;padding:10px 0">	lever:%s	</div>
		<div style="color:red;font-size:14px;">	[%s]	</div>
		<pre style="color:red;max-width: 500px;
		white-space: pre-wrap;
		word-wrap: break-word;">%s</pre>`, l.config.Project, l.config.Environment, level, time.Now(), info)
		a := message.AppMPNewsMessage{}
		a.AgentID = l.wecomClient.AgentId
		a.MsgType = "mpnews"
		a.ToUser = l.wecomClient.ToUser
		a.MPNews = struct {
			Articles []struct {
				Title            string `json:"title"`
				ThumbMediaID     string `json:"thumb_media_id"`
				Author           string `json:"author"`
				ContentSourceURL string `json:"content_source_url"`
				Content          string `json:"content"`
				Digest           string `json:"digest"`
			} `json:"articles"`
		}{
			Articles: []struct {
				Title            string `json:"title"`
				ThumbMediaID     string `json:"thumb_media_id"`
				Author           string `json:"author"`
				ContentSourceURL string `json:"content_source_url"`
				Content          string `json:"content"`
				Digest           string `json:"digest"`
			}{
				{
					Title:        "日志报警提醒",
					ThumbMediaID: "2vGmYtMGZH8dj5luO0W9vzMtjVYIEU39sl5Y-WnphMV1OVYNivT99SU2UJs-zLJuV",
					Author:       "蒋凯",
					Content:      content,
					Digest:       digest,
				},
			},
		}
		message := wecom.NewWecom(l.wecomClient.CorpId, l.wecomClient.CorpSecret, l.wecomClient.AgentId)
		msg, err := message.SendAppMessage(a)
		log.Println("msg", msg, err)
	}
}

func (l *Log) Print(v ...interface{}) {
	info := fmt.Sprint(v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, blue) + info
	}
	_ = log.Output(2, info)
}

func (l *Log) Printf(format string, v ...interface{}) {
	info := fmt.Sprintf(format, v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, blue) + info
	}
	_ = log.Output(2, info)

}

func (l *Log) Println(v ...interface{}) {
	info := fmt.Sprintln(v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, blue) + info
	}
	_ = log.Output(2, info)

}

func (l *Log) Fatal(v ...interface{}) {
	info := fmt.Sprint(v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, red) + info
	}
	_ = log.Output(2, info)
	l.SendMessage("Fatal", info)
	os.Exit(1)
}

func (l *Log) Fatalf(format string, v ...interface{}) {
	info := fmt.Sprintf(format, v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, red) + info
	}
	_ = log.Output(2, info)
	l.SendMessage("Fatal", info)
	os.Exit(1)
}

func (l *Log) Fatalln(v ...interface{}) {
	info := fmt.Sprintln(v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, red) + info
	}
	_ = log.Output(2, info)
	l.SendMessage("Fatal", info)
	os.Exit(1)
}

func (l *Log) Panic(v ...interface{}) {
	info := fmt.Sprint(v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, red) + info
	}
	_ = log.Output(2, info)
	panic(info)
}

func (l *Log) Panicf(format string, v ...interface{}) {
	info := fmt.Sprintf(format, v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, red) + info
	}
	_ = log.Output(2, info)
	l.SendMessage("Panic", info)
	panic(info)
}

func (l *Log) Panicln(v ...interface{}) {
	info := fmt.Sprintln(v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, red) + info
	}
	_ = log.Output(2, info)
	l.SendMessage("Panic", info)
	panic(info)
}

func (l *Log) Warning(v ...interface{}) {
	info := fmt.Sprint(v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, yellow) + info
	}
	_ = log.Output(2, info)
	l.SendMessage("Warning", info)
}

func (l *Log) Warningf(format string, v ...interface{}) {
	info := fmt.Sprintf(format, v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, yellow) + info
	}
	_ = log.Output(2, info)
	l.SendMessage("Warning", info)
}

func (l *Log) Warningln(v ...interface{}) {
	info := fmt.Sprintln(v...)
	if l.prefix != "" {
		info = Setcolor(l.prefix, yellow) + info
	}
	_ = log.Output(2, info)
	l.SendMessage("Warning", info)
}
