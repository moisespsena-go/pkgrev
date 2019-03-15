package pkgrev

import (
	"github.com/aghape/admin"
	"github.com/aghape/aghape"
	"github.com/moisespsena-go/xroute"
)

func Admin(Admin *admin.Admin, agp *aghape.Aghape) {
	prj := Admin.AddResource(&Project{}, &admin.Config{
		Setup: func(prj *admin.Resource) {

		},
	})

	oss := prj.AddResourceField("Oss", &Os{}, func(res *admin.Resource) {
		res.NewAttrs("ID", res.NewAttrs())
	})
	arch := oss.AddResourceField("Archs", &Arch{}, func(res *admin.Resource) {
		res.NewAttrs("ID", res.NewAttrs())
	})
	vrs := arch.AddResourceField("Versions", &Version{}, func(res *admin.Resource) {
		res.SetMeta(&admin.Meta{Name: "Version", Type: "string"})
	})
	vrs.AddResourceField("Files", &File{})

	Admin.OnRouter(func(r xroute.Router) {
		r.Put("/register/{project}/{os}/{arch}", admin.NewHandler(func(c *admin.Context) {
			project, os, arch := c.URLParam("project"), c.URLParam("os"), c.URLParam("arch")
			println(project)
			println(os)
			println(arch)
			c.Write([]byte("hello"))
		}))
	})
}

