package mdy

import (
	"fmt"
	"sync"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

var (
	mdyPool *sync.Pool
	once    sync.Once
)

// Client 明道云的client
type Client struct {
	client    sync.Pool // resty.Client
	appKey    string    // 明道云 appKey
	secretKey string    // 明道云 secretKey
	sign      string    // 明道云 sign
}

func InitPool(appKey, sign, secretKey string) *sync.Pool {
	once.Do(func() {
		mdyPool = &sync.Pool{
			New: func() any {
				return New(appKey, sign, secretKey)
			},
		}
	})
	return mdyPool
}

func Get() *Client {
	return mdyPool.Get().(*Client)
}

func Put(mdyObj *Client) {
	mdyPool.Put(mdyObj)
}

// New creates a new Client
// secretKey和sign必须有一个不能为空
func New(appKey, sign, secretKey string) *Client {
	if appKey == "" {
		panic("appKey cannot be empty")
	}

	if sign == "" && secretKey == "" {
		panic("sign or secretKey cannot be empty")
	}

	m := &Client{
		appKey:    appKey,
		secretKey: secretKey,
		sign:      sign,
	}

	m.client.New = func() any {
		return buildClient(false)
	}

	return m
}

// buildClient build a *resty.Client
func buildClient(debug bool) *resty.Client {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	client := resty.New()
	client.SetDebug(debug)
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal
	client.OnAfterResponse(WithMdyAfterResponse)
	return client
}

// WithMdyAfterResponse 状态码错误处理
func WithMdyAfterResponse(_ *resty.Client, res *resty.Response) error {
	if res.IsError() {
		return fmt.Errorf("请求'%v'发生了错误STATUS: %v", res.Request.URL, res.Status())
	}

	if mdyRes, ok := res.Result().(ResponseOk); ok {

		// 明道云请求返回的success 是否成功
		if !mdyRes.Ok() {
			return fmt.Errorf("请求'%v'发生了错误:%v-%v", res.Request.URL, mdyRes.Code(), mdyRes.Msg())
		}
	}

	return nil
}

// EnabledDebug 开启debug日志
func (m *Client) EnabledDebug() *Client {
	m.client.New = func() any {
		return buildClient(true)
	}
	return m
}

// SetClient set a *resty.Client
func (m *Client) SetClient(new func() *resty.Client) *Client {
	if new != nil {
		panic("new func must not be nil")
	}
	m.client = sync.Pool{
		New: func() any {
			return new()
		},
	}
	return m
}

// SetAppKey set an app key
func (m *Client) SetAppKey(appKey string) *Client {
	m.appKey = appKey
	return m
}

// SetSign set a sign
func (m *Client) SetSign(sign string) *Client {
	m.sign = sign
	return m
}

// SetSecretKey set a secret key
func (m *Client) SetSecretKey(secretKey string) *Client {
	m.secretKey = secretKey
	return m
}

// GetClient return a resty.Client
func (m *Client) getClient() *resty.Client {
	return m.client.Get().(*resty.Client)
}

// GetClient return a resty.Client
func (m *Client) freeClient(c *resty.Client) {
	m.client.Put(c)
}

func (m *Client) WorkSheetReq() *WorkSheetRequest {
	client := m.client.Get().(*resty.Client)
	defer m.freeClient(client)
	return &WorkSheetRequest{
		mdy: m,
		Req: client.R(),
	}
}

// AppReq return an AppRequest
func (m *Client) AppReq() *AppRequest {
	client := m.client.Get().(*resty.Client)
	defer m.freeClient(client)
	return &AppRequest{
		mdy: m,
		Req: client.R(),
	}
}

// GetSign return a sign
func (m *Client) GetSign() string {
	return m.sign
}

// GetAppKey return a appKey
func (m *Client) GetAppKey() string {
	return m.appKey
}

// Free free *Client
func (m *Client) Free() {
	Put(m)
}
