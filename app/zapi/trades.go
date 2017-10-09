package zapi

import (
	"errors"
	"net/http"
)

type PublicTradeData struct {
	Date         CustomTime `json:"date"`          // 取引日時
	Price        float64    `json:"price"`         // 取引価格
	Amount       float64    `json:"amount"`        // 取引量
	Tid          uint64     `json:"tid"`           // 取引ID
	CurrencyPair string     `json:"currency_pair"` // 通貨ペア
	TradeType    string     `json:"trade_type"`    // 取引種別
}
type PublicTrades struct {
	Trades []PublicTradeData `json:"trades"`
}

func GetPublicTrades(cp string) (t *PublicTrades, err error) {
	if cp == "" {
		return nil, errors.New("cp invalid error")
	}
	req, nrerr := http.NewRequest("GET", "https://api.zaif.jp/api/1/trades/"+cp, nil)
	if nrerr != nil {
		return nil, nrerr
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	resp, doerr := zAPIClient.Do(req)
	if doerr != nil {
		// エラーがnilじゃない場合、すでにBodyは閉じられている
		if rerr := getRedirectError(doerr); rerr != nil {
			err = rerr
		} else {
			err = doerr
		}
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		t = &PublicTrades{}
		err = responseRead(resp, &t.Trades)
	default:
		err = apiError(resp)
	}
	return
}
