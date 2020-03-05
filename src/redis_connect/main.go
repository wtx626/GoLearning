package main

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

type cfg struct {
	Cluster0 []RedisAddr // 集群1列表
	Cluster1 []RedisAddr // 集群2列表
	Password string
	MaxConns int
}

type RedisAddr struct {
	Host string
	Port int
}

func (o RedisAddr) String() string {
	return fmt.Sprintf("%v:%v", o.Host, o.Port)
}

type AppCfg struct {
	RedisConf cfg
}

func (c AppCfg) RedisOptions() [2]*redis.ClusterOptions {
	var opts [2]*redis.ClusterOptions
	rc := c.RedisConf
	convert := func(r []RedisAddr) *redis.ClusterOptions {
		opt := redis.ClusterOptions{
			Password:     rc.Password,
			ReadTimeout:  DefaultRedisReadTimeout,
			WriteTimeout: DefaultRedisWriteTimeout,
			DialTimeout:  DefaultRedisDialTimeout,
			PoolSize:     DefaultRedisPoolSize,
		}
		addrs := make([]string, len(r))
		for i, addr := range r {
			addrs[i] = addr.String()
		}
		opt.Addrs = addrs
		return &opt
	}
	opts[0] = convert(c.RedisConf.Cluster0)
	opts[1] = convert(c.RedisConf.Cluster1)
	return opts
}

var (
	DefaultRedisDialTimeout  = time.Millisecond * 10
	DefaultRedisReadTimeout  = time.Second * 10
	DefaultRedisWriteTimeout = time.Second * 10
	DefaultRedisPoolSize     = 1000
	appCfg                   AppCfg
	ErrNotExists             = errors.New("not exists")
	ErrEmptyTable            = errors.New("empty table")
	redisClusters            [2]*redis.ClusterClient
)

func InitRedis(rcs [2]*redis.ClusterOptions) error {
	for i, rc := range rcs {
		wrRc := *rc
		wrRc.ReadOnly = false
		r := redis.NewClusterClient(&wrRc)

		// if err := checkRedis(r); err != nil {
		// 	devcomm.Log.Errorf("check rediscluster %v:wr failed, err=%v", i, err)
		// 	return err
		// }

		redisClusters[i] = r
	}
	return nil
}

func GetRedis(idx uint32) (*redis.ClusterClient, error) {
	if idx < 0 || idx > 1 {
		return nil, ErrNotExists
	}

	return redisClusters[idx], nil
}

func set(key string, val string, expTime int32, client *redis.ClusterClient) {
	client.Set(key, val, time.Duration(expTime)*time.Second)
	client.Save()
}

func insertSet(client *redis.ClusterClient, key string, items []*redis.Z) {
	for _, item := range items {
		client.ZAdd(key, item)
	}
}

var a int

func init() {
	a = 3
}

func main() {
	// parse config
	cfgFile := "./conf/redis.toml"
	if _, err := toml.DecodeFile(cfgFile, &appCfg); err != nil {
		fmt.Printf("Config file %s decode err: %v\n", cfgFile, err)
		return
	}
	err := InitRedis(appCfg.RedisOptions())
	if err != nil {
		fmt.Printf("Config file err: %v\n", err)
		return
	}
	//client, err := GetRedis(0)
	//if err != nil {
	//	fmt.Printf("Config redis client err: %v\n", err)
	//	return
	//}

	//set("37_alter", "1", 0, client)
	//set("37_alter_1", "123", 0, client)
	//set("37_alter_1_2", "1234", 0, client)
	//set("37_", "45", 0, client)
	//insert(client)
	//推荐商品列表
	//items := map[string][]*redis.Z{
	//	"38_alter": {
	//		{100, "a1"}, {60, "a2"}, {40, "bq"}, {70, "s1"},
	//	},
	//	"38_alter_1": {
	//		{100, "b1"}, {30, "d2"},
	//	},
	//	"38_alter_1_2": {
	//		{50, "cf"}, {70, "df"},
	//	},
	//	"38_": {
	//		{40, "ryt"}, {6, "qq"}, {20, "gta"},
	//	},
	//}
	//for k, v := range items {
	//	insertSet(client, k, v)
	//}
	//client.ForEachMaster(func(cl *redis.Client) error {
	//	keys := cl.Scan(0, "38_alter*", 1).Iterator()
	//	for keys.Next() {
	//		//val := cl.Get(keys.Val())
	//		val1, err := cl.ZRevRangeWithScores(keys.Val(), 0, -1).Result()
	//		if err != nil {
	//			fmt.Println("获取出错", err)
	//		}
	//		fmt.Println(keys.Val(), val1)
	//	}
	//	return nil
	//})
	//fmt.Println(client.Get("37_alter").Result())
	//fmt.Println(client.Get("37_alter_1").Result())
	//fmt.Println(client.Get("37_alter_1_2").Result())
	//fmt.Println(client.Get("37_").Result())
	//r, err := client.Scan(0,"37_",1).Iterator()
	//fmt.Println(r, err)

	//测试地理位置
	//client.GeoAdd("4044_lbs", &redis.GeoLocation{"test", 113.980799, 22.541701, 0, 0})
	//fmt.Println(client.GeoRadius("4044_lbs", 113.980789, 22.541691, &redis.GeoRadiusQuery{
	//	Radius:      50,
	//	Unit:        "m",
	//	WithCoord:   false,
	//	WithDist:    true,
	//	WithGeoHash: false,
	//	Count:       0,
	//	Sort:        "ASC",
	//	Store:       "",
	//	StoreDist:   "",
	//}).Result())

	//var testMap sync.Map
	//testMap.Store("dfs", 12)
	//value, ok := testMap.Load("dfs")
	//if !ok {
	//
	//}
	//fmt.Println(reflect.TypeOf(value))
	//v, ok := value.(int)
	//if !ok {
	//
	//}
	//fmt.Println(uint32(v))

	//testarr := []int{1, 2, 3}
	//fmt.Println(testarr[:3])
	//a = 5
	//fmt.Println(a)
	var test sync.Map
	test.Store(1,"df")
	test.Store(2,"hel")
	test.Range(func(k, v interface{}) bool {
		test.Delete(k)
		return true
	})



	test.Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		return true
	})
}
