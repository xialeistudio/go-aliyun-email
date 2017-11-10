package aliyun_email

import (
	"testing"
	"os"
)

func TestClientSingleRequest(t *testing.T) {
	client := NewClient(
		os.Getenv("ACCESS_KEY_ID"),
		os.Getenv("ACCESS_SECRET"),
		os.Getenv("ACCOUNT_NAME"),
		os.Getenv("FROM_ALIAS"),
		RegionCNHangZhou,
	)
	req := &SingleRequest{
		ReplyToAddress: true,
		AddressType:    1,
		ToAddress:      "1065890063@qq.com",
		Subject:        "天气不错",
		HtmlBody:       "<h1>天气不错</h1>",
	}
	resp, err := client.SingleRequest(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}

func TestClientBatchSendEmail(t *testing.T) {
	client := NewClient(
		os.Getenv("ACCESS_KEY_ID"),
		os.Getenv("ACCESS_SECRET"),
		os.Getenv("ACCOUNT_NAME"),
		os.Getenv("FROM_ALIAS"),
		RegionCNHangZhou,
	)
	req := &BatchRequest{
		AddressType:   1,
		TemplateName:  "xx",
		ReceiversName: "xx",
	}
	resp, err := client.BatchSendEmail(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
