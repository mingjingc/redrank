package main

import (
	"fmt"
	"github.com/mingjingc/redrank/rank"
)

func main() {

	//r := redis.NewClient(&redis.Options{
	//	Addr: ":6379",
	//})
	//
	////list, _ := r.Do("zrevrange", "ar", 0, -1, "withscores").Result()
	////fmt.Println(list)
	////fmt.Println(reflect.TypeOf(list.([]interface{})[1]))
	////
	//score, _:= r.ZScore("ar","a").Result()
	//fmt.Println(score)

	accountBalanceRank := rank.NewRedRank("acccount_balance_rank", rank.RedisSettings{Addr:":6379"})
	fmt.Println(accountBalanceRank)


	//res,err := r.Do("zrevrange","ar",0,0,"WITHSCORES").Result()
	//fmt.Println(err)
	//fmt.Println(res)
	//fmt.Println(reflect.TypeOf(res.([]interface{})[1]))

	//client,err := rpc.Dial("http://127.0.0.1:8545/")
	//if err!=nil {
	//	log.Fatal(err)
	//}
	//var balance string
	//err = client.Call(&balance, "eth_getBalance", "0x104208d818afa0aaa6eb63c3ff0a120c5336a81a","latest")
	//if err!=nil {
	//	log.Fatal(err)
	//}
	//b,_ := big.NewInt(0).SetString(balance[2:], 16)
	//fmt.Println(b.String())
}
