package adminlte

import (
	adminTemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/components"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/resource"
	"html/template"
)

const (
	ColorschemeSkinBlack       = "skin-black"
	ColorschemeSkinBlackLight  = "skin-black-light"
	ColorschemeSkinBlue        = "skin-blue"
	ColorschemeSkinBlueLight   = "skin-blue-light"
	ColorschemeSkinGreen       = "skin-green"
	ColorschemeSkinGreenLight  = "skin-green-light"
	ColorschemeSkinPurple      = "skin-purple"
	ColorschemeSkinPurpleLight = "skin-purple-light"
	ColorschemeSkinRed         = "skin-red"
	ColorschemeSkinRedLight    = "skin-red-light"
	ColorschemeSkinYellow      = "skin-yellow"
	ColorschemeSkinYellowLight = "skin-yellow-light"
)

type Theme struct {
	Name string
	components.Base
}

var Adminlte = Theme{
	Name: "adminlte",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: TemplateList,
		},
	},
}

func init() {
	adminTemplate.Add("adminlte", &Adminlte)
}

func Get() *Theme {
	return &Adminlte
}

func (*Theme) GetTmplList() map[string]string {
	return TemplateList
}

func (*Theme) GetTemplate(isPjax bool) (tmpl *template.Template, name string) {
	var err error

	if !isPjax {
		name = "layout"
		tmpl, err = template.New("layout").Funcs(adminTemplate.DefaultFuncMap).
			Parse(TemplateList["layout"] +
				TemplateList["head"] + TemplateList["header"] + TemplateList["sidebar"] +
				TemplateList["footer"] + TemplateList["js"] + TemplateList["menu"] +
				TemplateList["admin_panel"] + TemplateList["content"])
	} else {
		name = "content"
		tmpl, err = template.New("content").Funcs(adminTemplate.DefaultFuncMap).
			Parse(TemplateList["admin_panel"] + TemplateList["content"])
	}

	if err != nil {
		panic(err)
	}

	return
}

func (*Theme) GetAsset(path string) ([]byte, error) {
	path = "resource" + path
	return resource.Asset(path)
}

func (*Theme) GetAssetList() []string {
	return resource.AssetsList
}
