package home

import (
	"fmt"
)

type Index struct {
	Base
}

func (this *Index) GetIndex() {
	this.Ctx.ViewData("login_id", this.Login_id)
	this.GeneralView()
}

func (this *Index) GetTest() {
	fmt.Println("test")
	this.GeneralView()
}

func (this *Index) PostTest() {
	fmt.Println(this.Ctx.FormValues())
}
