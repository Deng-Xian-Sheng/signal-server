//Copyright (c) [2023] [JinCanQi]
//[signal-server] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package custerrors

const (
	KeyIsEmpty                       = "key为空"
	KeyIsTooLong                     = "key太长，长度不能超过32位"
	KeyNotFound                      = "key未找到"
	CtxIsEmpty                       = "ctx为空"
	CacheOverrunMaxCost              = "缓存超出MaxCost，服务已达到最大负载"
	CacheOnEvictNoGetKey             = "缓存OnEvict回调，无法获取原始key"
	QueueNewFailBecauseAlreadyExists = "队列创建失败，因为已经存在"
	QueueGetFailBecauseNotExists     = "队列获取失败，因为不存在"
	DataIsEmpty                      = "data为空"
	KeyIsExpired                     = "key已过期"

	// offer 或 answer 的 sdp
	SdpNoValues     = "sdp没有值" // 对方（offer或answer）未放入，或已经被取出
	SdpAlreadyExist = "sdp已存在" // 对方（offer或answer）已经放入过一次，且只能放入一次

	// offer 或 answer 的 candidate
	CandidateNoValues = "candidate没有值" // 对方（offer或answer）从未放入，或已全部取出

	BodyIsEmpty = "body为空"
)
