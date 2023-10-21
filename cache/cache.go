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
	"log"
	"signal-server/custerrors"
	"signal-server/queue"
	"strings"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/dgraph-io/ristretto/z"
)

var (
	Cache      = &myCache{Cache: &ristretto.Cache{}}
	keyHashMap = &sync.Map{}
)

const DefMaxCacheTTL = 30 * time.Second

func init() {
	var err error
	Cache.Cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // 用于跟踪频率的键数 （10M）。
		MaxCost:     1 << 30, // 缓存的最大成本 （1GB）。
		BufferItems: 64,      // 每个获取缓冲区的键数。
		OnEvict: func(item *ristretto.Item) {
			defer func() {
				if err := recover(); err != nil {
					log.Println("cache OnReject回调发生错误 ", err)
				}
			}()
			// 根据uint64 key conflict 获取字符串key
			if key, ok := keyHashMap.Load(fmt.Sprint(item.Key, item.Conflict)); ok && key != nil && strings.HasPrefix(key.(string), GetKeyConst()) {
				wg := &sync.WaitGroup{}
				wg.Add(4)
				go func() {
					defer func() {
						if err := recover(); err != nil {
							log.Println("cache OnReject回调发生错误 ", err)
						}
						wg.Done()
					}()
					queue.GetOfferSdpQueue(key.(string)).Close()
				}()
				go func() {
					defer func() {
						if err := recover(); err != nil {
							log.Println("cache OnReject回调发生错误 ", err)
						}
						wg.Done()
					}()
					queue.GetOfferCandidateQueue(key.(string)).Close()
				}()
				go func() {
					defer func() {
						if err := recover(); err != nil {
							log.Println("cache OnReject回调发生错误 ", err)
						}
						wg.Done()
					}()
					queue.GetAnswerSdpQueue(key.(string)).Close()
				}()
				go func() {
					defer func() {
						if err := recover(); err != nil {
							log.Println("cache OnReject回调发生错误 ", err)
						}
						wg.Done()
					}()
					queue.GetAnswerCandidateQueue(key.(string)).Close()
				}()
				wg.Wait()
				keyHashMap.Delete(fmt.Sprint(item.Key, item.Conflict))
			} else {
				log.Panicln(custerrors.CacheOnEvictNoGetKey)
			}
		},
		OnReject: func(item *ristretto.Item) {
			defer func() {
				if err := recover(); err != nil {
					log.Println("cache OnReject回调发生错误 ", err)
				}
			}()
			log.Panicln(custerrors.CacheOverrunMaxCost)
		},
		KeyToHash: func(key interface{}) (uint64, uint64) {
			v, vv := z.KeyToHash(key)
			if vvv, ok := key.(string); ok && strings.HasPrefix(vvv, GetKeyConst()) {
				keyHashMap.Store(fmt.Sprint(v, vv), key)
			}
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
	value = fmt.Sprint(int64(ttl))
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

func (c *myCache) GetTTL(key interface{}) (time.Duration, bool) {
	if keyStr, ok := key.(string); ok {
		key = fmt.Sprintf(keyConst, keyStr)
	}
	return c.Cache.GetTTL(key)
}
