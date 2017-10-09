package main

import (
	"fmt"
	"github.com/tanaton/peperoncino/app/zapi"
)

func main() {
	t, err := zapi.GetPublicTrades("pepecash_jpy")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, it := range t.Trades {
		var tt string
		if it.TradeType == "bid" {
			tt = "買"
		} else {
			tt = "売"
		}
		fmt.Printf("%s => %.4f\n", tt, it.Amount)
	}
}
