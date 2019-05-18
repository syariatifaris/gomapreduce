package gomapreduce

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"encoding/json"

	"github.com/go-redis/redis"
)

func New() MapReduce {
	return MapReduce{
		clt: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
		identityGroups: make(map[string]*IdentityGroup, 0),
	}
}

type MapReduce struct {
	sync.Mutex
	wg             sync.WaitGroup
	identityGroups map[string]*IdentityGroup
	clt            *redis.Client
}

func (m *MapReduce) Map(identity *UserIdentity) {
	m.Lock()
	defer m.Unlock()
	key := createKey(Alphabetic(identity.Attributes))
	if _, ok := m.identityGroups[key]; !ok {
		m.identityGroups[key] = &IdentityGroup{
			Attributes: identity.Attributes,
			NIS:        []string{},
		}
		m.wg.Add(1)
		go m.Reduce(key)
	}
	m.identityGroups[key].NIS = append(m.identityGroups[key].NIS, identity.NIS)
	go func(r *redis.Client, id *UserIdentity) {
		bytes, _ := json.Marshal(id)
		r.Set(fmt.Sprintf("dummy_%s", id.NIS), string(bytes), time.Second)
	}(m.clt, identity)
}

func (m *MapReduce) Reduce(key string) {
	m.Lock()
	defer m.Unlock()
	cpgroup := m.identityGroups[key]
	for _, segment := range Segments {
		if Match(segment.Attributes, cpgroup.Attributes) {
			m.identityGroups[key].Segments = append(m.identityGroups[key].Segments, segment.Name)
		}
	}
	m.wg.Done()
}

func (m *MapReduce) Process() {
	m.wg.Wait()
}

func Match(segmentAttrs []string, identityGroupAttrs []string) bool {
	for _, segmentAttr := range segmentAttrs {
		if !stringInSlice(segmentAttr, identityGroupAttrs) {
			return false
		}
	}
	return true
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func createKey(attrs []string) string {
	return strings.Join(attrs, "-")
}
