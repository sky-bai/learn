package main

import (
	"context"
	"fmt"
	"github.com/longportapp/openapi-go/config"
	"github.com/longportapp/openapi-go/quote"
	"log"
	"os"
)

func main() {

	// https://github.com/longportapp/openapi-go/blob/v0.9.2/config/config_test.go#L11
	conf, err := config.New(config.WithConfigKey(os.Getenv("LONGPORT_APP_KEY"), os.Getenv("LONGPORT_APP_SECRET"), os.Getenv("LONGPORT_ACCESS_TOKEN")))
	if err != nil {
		log.Fatal(err)
		return
	}

	// 期权发现工具

	qctx, err := quote.NewFromCfg(conf)
	if err != nil {
		log.Fatal(err)
		return //
	}

	optionQuotes, err := qctx.OptionQuote(context.Background(), []string{"BABA230120C160000.US"})
	if err != nil {
		log.Fatal(err)
		return // 我突然想到 也就是说为什么同能在一起了，
	}

	for _, v := range optionQuotes {
		fmt.Println(v)
	}

	// 写下来要做的事情

	// 交易量 获取隐含波动率
	// 1.入参如何确定
	//

	// roll cover call
	// 这边

}
