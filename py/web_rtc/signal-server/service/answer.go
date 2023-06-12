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
