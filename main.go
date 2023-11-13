package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/kavindarochana/ussdapp/utils"
)

type ReqFrmt struct {
	SourceAddress string
	Message       string
	RequestId     int
	ApplicationId string
	Encoding      string
	Version       string
	SessionId     string
	UssdOperation string
}

type Out struct {
	ApplicationId      string
	Password           string
	Version            string
	Message            string
	SessionId          string
	UssdOperation      string
	DestinationAddress string
	Encoding           string
}

var err error

var redisClient *redis.Client

func runUssd(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var rq ReqFrmt
	err := decoder.Decode(&rq)
	if err != nil {
		panic(err)
	}
	utils.Debug("request - ", rq)

	var out Out
	out.ApplicationId = rq.ApplicationId
	out.DestinationAddress = rq.SourceAddress
	out.Encoding = rq.Encoding
	out.Message = "Welcome\n 1. Name \n2. Age \n3. Exit"
	out.Password = ""
	out.SessionId = rq.SessionId
	out.UssdOperation = ""
	out.Version = rq.UssdOperation

	utils.Debug("USSD Operation - ", rq.UssdOperation)

	if "mo-init" == rq.UssdOperation {
		utils.Debug("USSD Operation - ", "mo-init")

		setSession(rq.SessionId, "main")
	} else {
		utils.Debug("USSD Operation - ", rq.UssdOperation)

		uMenu := getSession(rq.SessionId)
		utils.Debug("USSD Operation - ", uMenu)

		switch uMenu {
		case "main":
			switch rq.Message {

			case "1":
				utils.Debug("USSD Operation - main->1")

				out.Message = "Enter your name"

				setSession(rq.SessionId, "name")
			}

		case "name":
			out.Message = "Hi " + rq.Message
		}

	}

	sender(out)
}

func main() {

	http.HandleFunc("/ussd/v1/app", runUssd)
	fmt.Println(http.ListenAndServe(":9092", nil))
}

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
}

func setSession(sId, menu string) {
	err = redisClient.Set(sId, menu, 0).Err()
	if err != nil {
		panic(err)
	}
}

func getSession(sId string) string {
	val, err := redisClient.Get(sId).Result()
	if err != nil {
		panic(err)
	}
	return val
}
