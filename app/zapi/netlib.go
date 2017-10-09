package zapi

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	ConnectStreamTimeoutSec  = 10
	ConnectTimeoutSec        = 10
	CommunicationTimeoutSec  = 10
	ResponseSizeMax          = 1024 * 1024 * 10
	ResponseStreamTimeoutSec = 20
	StreamBufferSize         = 64 * 1024
)

type ErrorCode struct {
	Code     int64  `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"moreInfo"`
}
type HeartBeat struct {
	Time CustomTime `json:"time"`
}

var zAPIClient *http.Client
var zAPIStreamClient *http.Client

func init() {
	zAPIClient = &http.Client{
		Transport: &http.Transport{
			Dial:                  dialTimeout,
			DisableKeepAlives:     false,
			DisableCompression:    true, // 圧縮解凍は全てこっちで指示する
			ResponseHeaderTimeout: ConnectTimeoutSec * time.Second,
			MaxIdleConnsPerHost:   10, // 同じドメインへの接続数を10へ拡張
		},
		CheckRedirect: redirectPolicy,
	}
	zAPIStreamClient = &http.Client{
		Transport: &http.Transport{
			Dial:                  dialTimeoutStream,
			DisableKeepAlives:     false,
			DisableCompression:    true, // 圧縮解凍は全てこっちで指示する
			ResponseHeaderTimeout: ConnectTimeoutSec * time.Second,
		},
		CheckRedirect: redirectPolicy,
	}
}

func dialTimeout(network, addr string) (net.Conn, error) {
	con, err := net.DialTimeout(network, addr, ConnectTimeoutSec*time.Second)
	if err == nil {
		con.SetDeadline(time.Now().Add(CommunicationTimeoutSec * time.Second))
	}
	return con, err
}

func dialTimeoutStream(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, ConnectStreamTimeoutSec*time.Second)
}

func responseRead(resp *http.Response, v interface{}) (err error) {
	var r io.Reader
	var gz io.ReadCloser

	ce := resp.Header.Get("Content-Encoding")
	if ce == "gzip" {
		// 解凍する
		gz, _ = gzip.NewReader(resp.Body)
		r = gz
	} else {
		// 圧縮されていない場合
		r = resp.Body
	}
	err = json.NewDecoder(io.LimitReader(r, ResponseSizeMax)).Decode(v)
	if gz != nil {
		gz.Close()
	}
	return
}

func apiError(resp *http.Response) error {
	ec := ErrorCode{}
	err := responseRead(resp, &ec)
	if err == nil {
		msg := ec.Message
		if msg == "" {
			msg = "response error"
		}
		err = errors.New(msg)
	}
	return err
}

type RedirectError struct {
	Host string
	Path string
	Msg  string
}

func (e *RedirectError) Error() string {
	return e.Msg
}

func redirectPolicy(r *http.Request, _ []*http.Request) error {
	return &RedirectError{r.URL.Host, r.URL.Path, "redirect error"}
}

func getRedirectError(err error) *RedirectError {
	uerr, uok := err.(*url.Error)
	if !uok {
		return nil
	}
	rerr, rok := uerr.Err.(*RedirectError)
	if !rok {
		return nil
	}
	return rerr
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	ct.Time = time.Unix(ts, 0)
	return nil
}

func (ct *CustomTime) String() string {
	return strconv.FormatInt(ct.Time.Unix(), 10)
}
