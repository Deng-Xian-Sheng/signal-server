//Copyright (c) [2023] [JinCanQi]
//[make_data_set_so-vits-svc] is licensed under Mulan PubL v2.
//You can use this software according to the terms and conditions of the Mulan PubL v2.
//You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PubL v2 for more details.

package service

import "github.com/gin-gonic/gin"

type answer struct{}

var Answer = &answer{}

func (o *answer) SdpGet(c *gin.Context) {

}

func (o *answer) SdpPost(c *gin.Context) {

}

func (o *answer) CandidateGet(c *gin.Context) {

}

func (o *answer) CandidatePost(c *gin.Context) {

}
