package service

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"operation-backend/tool"
	"os/exec"
)
type DeployReq struct {
	Image  string `form:"image" json:"image" binding:"required"`
	Secret string `form:"secret" json:"secret" binding:"required"`
}


func Deploy(c *gin.Context) {
	// 1. 绑定参数
	req := DeployReq{}
	err := c.Bind(&req)
	if err != nil {
		tool.RenderErrJson(c, err)
		return
	}
	command := "./test.sh "+req.Image+" "+req.Secret//脚本的路径
	cmd := exec.Command("/bin/bash", "-c",command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}
