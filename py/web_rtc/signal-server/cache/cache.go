//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package cache

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
	"github.com/dgraph-io/ristretto/z"
	"log"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/custerrors"
	"make_data_set_so-vits-svc/py/web_rtc/signal-server/queue"
	"strings"
	"sync"
	"time"
)

var (
	Cache      = &myCache{Cache: &ristretto.Cache{}}
	KeyHashMap = &sync.Map{}
)

const DefMaxCacheTTL = 30 * time.Second

func init() {
	var err error
	Cache.Cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // 用于跟踪频率的键数 （10M）。
		MaxCost:     1 << 30, // 缓存的最大成本 （1GB）。
		BufferItems: 64,      // 每个获取缓冲区的键数。
		OnEvict: func(item *ristretto.Item) {
			// 根据uint64 key conflict 获取字符串key
			if key, ok := KeyHashMap.Load(fmt.Sprint(item.Key, item.Conflict)); ok && strings.HasPrefix(key.(string), GetKeyConst()) {
				queue.GetOfferSdpQueue(key.(string)).Close()
				queue.GetOfferCandidateQueue(key.(string)).Close()
				queue.GetAnswerSdpQueue(key.(string)).Close()
				queue.GetAnswerCandidateQueue(key.(string)).Close()
			} else {
				log.Panicln(custerrors.CacheOnEvictNoGetKey)
			}
		},
		OnReject: func(item *ristretto.Item) {
			log.Panicln(custerrors.CacheOverrunMaxCost)
		},
		KeyToHash: func(key interface{}) (uint64, uint64) {
			v, vv := z.KeyToHash(key)
			KeyHashMap.Store(fmt.Sprint(v, vv), key)
			return v, vv
		},
	})
	if err != nil {
		log.Panicln(err)
	}
}

const (
	keyConst = "key_%s"
)

func GetKeyConst(key ...string) string {
	if len(key) == 0 {
		return fmt.Sprintf(keyConst, "")
	}
	return fmt.Sprintf(keyConst, key)
}

type myCache struct {
	*ristretto.Cache
}

// SetWithTTL works like Set but adds a key-value pair to the cache that will expire
// after the specified TTL (time to live) has passed. A zero value means the value never
// expires, which is identical to calling Set. A negative value is a no-op and the value
// is discarded.
func (c *myCache) SetWithTTL(key interface{}, value interface{}, cost int64, ttl time.Duration) bool {
	if ttl <= 0 || ttl > DefMaxCacheTTL {
		ttl = DefMaxCacheTTL
	}
	cost = 1
	if keyStr, ok := key.(string); ok {
		key = fmt.Sprintf(keyConst, keyStr)
	}
	if ok := c.Cache.SetWithTTL(key, value, cost, ttl); !ok {
		log.Panicln(custerrors.CacheOverrunMaxCost)
		return false
	}
	c.Cache.Wait()
	return true
}

func (c *myCache) Get(key interface{}) (interface{}, bool) {
	if keyStr, ok := key.(string); ok {
		key = fmt.Sprintf(keyConst, keyStr)
	}
	return c.Cache.Get(key)
}
