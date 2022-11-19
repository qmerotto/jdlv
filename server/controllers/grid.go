package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type input [][2]int
type GridController struct {
	web.Controller
}

func (gc *GridController) Get() {

}

func (gc *GridController) Post() {
	body, err := gc.RenderBytes()
	if err != nil {
		fmt.Printf(err.Error())
		gc.Finish()
	}

	var inputBody input
	if err := json.Unmarshal(body, &inputBody); err != nil {
		fmt.Printf(err.Error())
		gc.Finish()
	}

	fmt.Printf(fmt.Sprintf("%v", inputBody))
}
