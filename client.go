package aliyun_email

import (
	"time"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
)

const (
	RegionCNHangZhou   = "cn-hangzhou"
	RegionAPSouthEast1 = "ap-southeast-1"
	RegionAPSouthEast2 = "ap-southeast-2"
)

var (
	regionMap = map[string]string{
		RegionCNHangZhou:   "https://dm.aliyuncs.com",
		RegionAPSouthEast1: "https://dm.ap-southeast-1.aliyuncs.com",
		RegionAPSouthEast2: "https://dm.ap-southeast-2.aliyuncs.com",
	}
	versionMap = map[string]string{
		RegionCNHangZhou:   "2015-11-23",
		RegionAPSouthEast1: "2017-06-22",
		RegionAPSouthEast2: "2017-06-22",
	}
)

type client struct {
	accessKeyId     string
	accessKeySecret string
	accountName     string
	fromAlias       string
	regionId        string
}

type SingleRequest struct {
	ReplyToAddress bool
	AddressType    int `json:",int"`
	ToAddress      string
	Subject        string
	HtmlBody       string
	TextBody       string
	ClickTrace     string
}

func NewClient(accessKeyId, accessKeySecret, accountName, fromAlias, regionId string) *client {
	return &client{
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		accountName:     accountName,
		fromAlias:       fromAlias,
		regionId:        regionId,
	}
}

func (c *client) NewRequest() *Params {
	params := &Params{}
	params.PutString("Format", "JSON")
	params.PutString("Version", versionMap[c.regionId])
	params.PutString("AccessKeyId", c.accessKeyId)
	params.PutString("SignatureMethod", "HMAC-SHA1")
	params.PutString("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	params.PutString("SignatureVersion", "1.0")
	params.PutString("SignatureNonce", uuid.NewV4().String())
	params.PutString("RegionId", c.regionId)
	return params
}

func (c *client) request(method, link string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, link, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *client) SingleRequest(regionId string, req *SingleRequest) (map[string]interface{}, error) {
	params := c.NewRequest()
	params.PutString("Action", "SingleSendMail")
	params.PutString("AccountName", c.accountName)
	params.PutBoolean("ReplyToAddress", req.ReplyToAddress)
	params.PutInt("AddressType", req.AddressType)
	params.PutString("ToAddress", req.ToAddress)
	params.PutString("FromAlias", c.fromAlias)
	if req.Subject != "" {
		params.PutString("Subject", req.Subject)
	}
	if req.HtmlBody != "" {
		params.PutString("HtmlBody", req.HtmlBody)
	}
	if req.TextBody != "" {
		params.PutString("TextBody", req.TextBody)
	}
	if req.ClickTrace != "" {
		params.PutString("ClickTrace", req.ClickTrace)
	}
	params.Sign("POST", c.accessKeySecret+"&")

	buf, err := c.request("POST", regionMap[c.regionId], strings.NewReader(params.ToUrlValues().Encode()))
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(buf, &result)
	return result, err
}
