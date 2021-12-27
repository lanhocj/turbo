package controller

import "github.com/laamho/turbo/common/config"

type Menu struct {
	Name, Link string
	Role       int
	Menus      Menus
}

type Menus []*Menu

type GlobalPageData struct {
	Title, Description, Copyright string
	Menus                         Menus
	Body                          interface{}
}

func (g *GlobalPageData) AddMenu(name, link string) {
	g.Menus = append(g.Menus, &Menu{
		Name: name,
		Link: link,
	})
}

func (g *GlobalPageData) Init() {
	conf := config.Copy()
	g.Title = conf.App.Title
	g.Description = conf.App.Description
	g.Copyright = conf.App.Copyright

	g.AddMenu("配置文件", "/")
	g.AddMenu("节点", "/nodes")
	g.AddMenu("用户", "/users")
}

var GlobalData = &GlobalPageData{}
