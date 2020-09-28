# viinet-log

log for ViiNet

## Table of Contents

- [Requirements](#requirements)
- [Usage](#usage)

## Requirements

requires the following to run:

- go ^1.12.6

## Usage

```
    config := &log.Config{
		Project:     "test",
		Environment: "dev",
	}
    log := log.NewLog("123123", config)
    client := &log.WecomClient{
		CorpId:     "123",
		CorpSecret: "123",
		AgentId:    123,
		ToUser:     "@all",
	}
    log.UseWecom(client)
    log.Fatalf("err:%s", "1231sadsad")
```
