package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GetParams struct {
	Name string `json:"name" form:"name"`
}

type CreateParams struct {
	Name      string `json:"name" form:"name"`
	Namespace string `json:"namespace" form:"namespace"`
}

func AutoBindGet(c *gin.Context, params *GetParams) {
	fmt.Printf("name:%s", params.Name)
	fmt.Println("do something!!")
	return
}

func AutoBindCreate(c *gin.Context, params *CreateParams) {
	fmt.Printf("name:%s,namespace:%s", params.Name, params.Namespace)
	fmt.Println("do something!!")
	return
}
