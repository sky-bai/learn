package main

import (
	"fmt"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/menu"
	"log"
	"net/http"
)

func serveWechat(rw http.ResponseWriter, req *http.Request) {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     "wx812ceea1f2d8a505",
		AppSecret: "0ad5ee2b74902ea86f6d9686906ae660",
		Token:     "amwukyni123chrishy",
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	at, _ := officialAccount.GetAccessToken()
	log.Println(at)
	m := officialAccount.GetMenu()
	getMenu, err := m.GetMenu()
	if err != nil {
		log.Println(getMenu.Menu.MenuID)
	}

	buttons := make([]*menu.Button, 0)
	buttons = append(buttons, &menu.Button{
		Type:       "view",
		Name:       "订单查询",
		Key:        "V1001_GOOD",
		URL:        "https://www.restapp.cn/h5/#/",
		MediaID:    "",
		AppID:      "",
		PagePath:   "",
		SubButtons: nil,
	})
	buttons = append(buttons, &menu.Button{
		Type:       "miniprogram",
		Name:       "最新活动",
		Key:        "V1001_GOOD",
		URL:        "https://www.restapp.cn/h5/#/",
		MediaID:    "",
		AppID:      "wx2ea9729d329745dd",
		PagePath:   "pages/tabbar/sound/sound",
		SubButtons: nil,
	})
	buttons = append(buttons, &menu.Button{
		Type:       "click",
		Name:       "联系客服",
		Key:        "GETPICTURE",
		URL:        "",
		MediaID:    "",
		AppID:      "",
		PagePath:   "",
		SubButtons: nil,
	})
	err = m.SetMenu(buttons)
	log.Println(err)
	if err != nil {
		log.Println("设置成功")
		return
	}

}
func main() {
	http.HandleFunc("/", serveWechat)
	fmt.Println("wechat server listener at", ":8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}
}
