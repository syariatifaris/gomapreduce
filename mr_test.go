package gomapreduce

import (
	"fmt"
	"testing"
	"time"

	"encoding/json"

	"github.com/go-redis/redis"
)

var UserIdentities []*UserIdentity

func init() {
	UserIdentities = UserIdentitiesFunc()
	fmt.Println("start benchmarking..")
}

func BenchmarkMapReduce(b *testing.B) {
	mr := New()
	for _, u := range UserIdentities {
		mr.Map(u)
	}
	mr.Process()
}

func BenchmarkBlocking(b *testing.B) {
	clt := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	for _, u := range UserIdentities {
		bytes, _ := json.Marshal(u)
		clt.Set(fmt.Sprintf("dummy_%s", u.NIS), string(bytes), time.Second)
	}
	for _, s := range Segments {
		for _, u := range UserIdentities {
			if Match(s.Attributes, u.Attributes) {
				u.SegmentResults = append(u.SegmentResults, s.Name)
			}
		}
	}
}
