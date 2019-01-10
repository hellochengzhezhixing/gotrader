package  main

import(
	"fmt"
	"log"
	// exchange "github.com/thrasher-/gocryptotrader/exchanges"
	"github.com/thrasher-/gocryptotrader/exchanges/huobi"
	"github.com/thrasher-/gocryptotrader/config"
)

//OrderHuobi is ?
func OrderHuobi(){

	defaultPath, err := config.GetFilePath("")
	if err != nil {
		log.Fatal(err)
	}

	//初始化一个config对象
	var huobiconf = config.Cfg

	//载入配置文件信息
	err = huobiconf.LoadConfig(defaultPath)
	if err != nil {
		fmt.Println("failed to get the config file content")
	}

	//获得交易所配置中的信息
	var huobicfg config.ExchangeConfig
	for _,info := range huobiconf.Exchanges{
		if info.Name == "Huobi" {
			huobicfg = info
		}
	}
	

	// huo := new(exchange.IBotExchange)
	// fmt.Println("i am bot ??",huo)

	var huobiex huobi.HUOBI
	huobiex.SetDefaults()
	huobiex.Setup(huobicfg)

	hello,err := huobiex.GetAccounts()
	if err != nil {
		fmt.Println("failed to get the account:",err)
	}
	for _,account  := range hello {
		fmt.Println(".....",account)
	}
}
