package cache

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
	"log"
)

var Cache *ristretto.Cache

const DefCacheTTL = 30 // seconds

func init() {
	var err error
	Cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // 用于跟踪频率的键数 （10M）。
		MaxCost:     1 << 30, // 缓存的最大成本 （1GB）。
		BufferItems: 64,      // 每个获取缓冲区的键数。
	})
	if err != nil {
		log.Panicln(err)
	}
}

const (
	key                = "key_%s"
	offerSdpKey        = "offer_sdp_%s"
	offerCandidateKey  = "offer_candidate_%s"
	answerSdpKey       = "answer_sdp_%s"
	answerCandidateKey = "answer_candidate_%s"
)

func Key(key string) string {
	return fmt.Sprintf(key, key)
}

func OfferSdpKey(key string) string {
	return fmt.Sprintf(offerSdpKey, key)
}

func OfferCandidateKey(key string) string {
	return fmt.Sprintf(offerCandidateKey, key)
}

func AnswerSdpKey(key string) string {
	return fmt.Sprintf(answerSdpKey, key)
}

func AnswerCandidateKey(key string) string {
	return fmt.Sprintf(answerCandidateKey, key)
}
