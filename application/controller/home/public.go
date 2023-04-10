package home

type Public struct {
	Base
}

func (this *Public) GetLogin() {
	this.GeneralView()
}

func (this *Public) PostLogin() {
	this.Session.Set("login_id", 1)
}
