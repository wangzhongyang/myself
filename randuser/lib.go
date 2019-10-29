package randuser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	RedisPool       *redis.Pool
	setRandUserSync sync.Once
	redisKey        = "test:integration:rand_users"
)

func init() {
	SetRedisConn()
}

type RandUserInfo struct {
	Age      int `json:"age"`
	Birthday struct {
		Dmy string `json:"dmy"`
		Mdy string `json:"mdy"`
		Raw int    `json:"raw"`
	} `json:"birthday"`
	CreditCard struct {
		Expiration string `json:"expiration"`
		Number     string `json:"number"`
		Pin        int    `json:"pin"`
		Security   int    `json:"security"`
	} `json:"credit_card"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
	Region   string `json:"region"`
	Surname  string `json:"surname"`
	Sex      string `json:"sex"`
	Title    string `json:"title"`
}

// RandUserInfo 获取用户随机信息
// 项目原地址：https://uinames.com/
// api限制一分钟七次请求
// param:
// 	true:一分钟超过3500条
// 	false:一分钟不超过3500条
func GetRandUserInfo(isAll bool) (RandUserInfo, error) {
	var (
		itemString string
		res        interface{}
		err        error
	)

	conn := RedisPool.Get()
	defer conn.Close()
	getUser := func() ([]RandUserInfo, error) {
		url := "https://uinames.com/api/?amount=500&region=china&ext"
		resp, err := http.Get(url)
		if err != nil {
			return nil, errors.New("http get error:" + err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("get resp body error:" + err.Error())
		}
		users := make([]RandUserInfo, 500)
		if err := json.Unmarshal(body, &users); err != nil {
			return nil, errors.New(fmt.Sprintf("body json.Unmarshal error:%s, body:%s", err.Error(), body))
		}
		return users, nil
	}
	if isAll {
		// 每分钟超过3500，初始化请求七次API，随机取数据
		setRandUserSync.Do(func() {
			_, _ = conn.Do("DEL", redisKey)
			var wg sync.WaitGroup
			for i := 0; i < 7; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					users, err := getUser()
					if err != nil {
						panic(fmt.Sprintf("get rand user api failed,i is %d,error:%s", i, err.Error()))
					}
					if err := SetUserToRedis(users); err != nil {
						panic(err)
					}
				}()
			}
			wg.Wait()
		})
		res, _ = conn.Do("LINDEX", redisKey, rand.Intn(3500))

	} else {
		// 一分钟不超过3500次借口调用，每次pop
		if res, err = conn.Do("LPOP", redisKey); err != nil || res == nil {
			users, err := getUser()
			if err != nil {
				return RandUserInfo{}, err
			}
			if err := SetUserToRedis(users); err != nil {
				return RandUserInfo{}, err
			}
			res, _ = conn.Do("LPOP", redisKey)
		}
	}

	itemString = string(res.([]uint8))
	var user RandUserInfo
	_ = json.Unmarshal([]byte(itemString), &user)
	user.UserName = user.Name + user.Surname
	switch user.Gender {
	case "male":
		user.Sex = "男"
	case "female":
		user.Sex = "女"
	}
	return user, nil
}

func SetRedisConn() {
	RedisPool = &redis.Pool{
		Dial:            func() (redis.Conn, error) { return redis.Dial("tcp", ":6379") },
		MaxIdle:         3,
		IdleTimeout:     240 * time.Second,
		MaxConnLifetime: 60 * time.Second,
	}
}

func SetUserToRedis(users []RandUserInfo) error {
	conn := RedisPool.Get()
	defer conn.Close()
	for _, v := range users {
		vByte, _ := json.Marshal(v)
		if _, err := conn.Do("LPUSH", redisKey, string(vByte)); err != nil {
			return errors.New("redis LPUSH error:" + err.Error())
		}
	}
	return nil
}
