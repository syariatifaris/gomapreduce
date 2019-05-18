package main

import (
	"fmt"
	"time"

	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/syariatifaris/gomapreduce"
)

var UserIdentities []*gomapreduce.UserIdentity

func init() {
	UserIdentities = gomapreduce.UserIdentitiesFunc()
}

func main() {
	//case 1: waiting to write
	//write all to db (redis)
	start := time.Now()
	clt := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	for _, u := range UserIdentities {
		bytes, _ := json.Marshal(u)
		clt.Set(fmt.Sprintf("dummy_%s", u.NIS), string(bytes), time.Second)
	}
	for _, s := range gomapreduce.Segments {
		for _, u := range UserIdentities {
			if gomapreduce.Match(s.Attributes, u.Attributes) {
				u.SegmentResults = append(u.SegmentResults, s.Name)
			}
		}
	}
	fmt.Println("wait process", time.Since(start).Seconds())
	//case 2: using map-reduce
	//non-blocking
	start = time.Now()
	mr := gomapreduce.New()
	for _, u := range UserIdentities {
		mr.Map(u)
	}
	mr.Process()
	fmt.Println("map reduce:", time.Since(start).Seconds())
}
