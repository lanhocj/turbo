package internal

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
	CurrentPage                   string
}

func (g *GlobalPageData) AddMenu(name, link string, role int) {
	g.Menus = append(g.Menus, &Menu{
		Name: name,
		Link: link,
		Role: role,
	})
}

func (g *GlobalPageData) SetCurrentPath(p string) {
	g.CurrentPage = p
}

func (g *GlobalPageData) SetCurrentTitle(p string) {
	g.Title = p
}

func (g *GlobalPageData) Init() *GlobalPageData {
	conf := config.Copy()
	g.Title = conf.App.Title
	g.Description = conf.App.Description
	g.Copyright = conf.App.Copyright

	g.AddMenu("配置文件", "/", -1)
	g.AddMenu("节点", "/nodes", 1)
	g.AddMenu("用户", "/users", 1)
	return g
}

var Glob = &GlobalPageData{}
