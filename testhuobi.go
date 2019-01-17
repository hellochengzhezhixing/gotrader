package  main

import(
	"fmt"
	"log"
	// exchange "github.com/thrasher-/gocryptotrader/exchanges"
	"github.com/thrasher-/gocryptotrader/exchanges/huobi"
	"github.com/thrasher-/gocryptotrader/config"
	"strconv"

	"github.com/gorilla/websocket"
	"time"
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

		//现货
		if account.Type == "spot"{
			//挂单尝试
			arg := huobi.SpotNewOrderRequestParams{
				Symbol:    "eoseth",
				AccountID: int(account.ID),
				Amount:    0.1,
				Price:     0.001,
				Type:      huobi.SpotNewOrderRequestTypeBuyLimit,
			}
			buyorderid,err := huobiex.SpotNewOrder(arg)
			if err != nil {
				fmt.Println("failed to place an order....",err)
			}
			fmt.Println("succeed in placing a new buy order......",buyorderid)
			cancelorder,err := huobiex.CancelExistingOrder(buyorderid)
			if err != nil {
				fmt.Println("failed to cancel the existing order")
			}
			fmt.Println("cancelorder successful...",cancelorder)
			
			accountb,err := huobiex.GetAccountBalance(strconv.FormatInt(account.ID,10))
			if err != nil {
				fmt.Println("hello balance get errr...",err)
			}
			for _,accountbalance := range accountb{
				if accountbalance.Currency == "eth"{
					fmt.Println("....current account id : ",account.ID," balance : ",accountbalance)
					
				}
			}
		}

	}


	// wsclient := NewWebsocketHub()
	Strategy()

	// huobiex.

}

//strategy function for update...
func Strategy(){
	if !wsHubStarted {
		StartWebsocketHandler()
	}

	connectionLimit := bot.config.Webserver.WebsocketConnectionLimit
	numClients := len(wsHub.Clients)

	if numClients >= connectionLimit {
		log.Printf("websocket: client rejected due to websocket client limit reached. Number of clients %d. Limit %d.",
			numClients, connectionLimit)
		return
	}

	// upgrader := websocket.Upgrader{
	// 	WriteBufferSize: 1024,
	// 	ReadBufferSize:  1024,
	// }

	// conn, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	
	conn := websocket.Conn{}
	client := &WebsocketClient{Hub: wsHub, Conn: &conn, Send: make(chan []byte, 1024)}
	client.Hub.Register <- client
	log.Printf("websocket: client connected. Connected clients: %d. Limit %d.",
		numClients+1, connectionLimit)

	for {
		client.read()
		time.Sleep(1 * time.Second)
	}
}
