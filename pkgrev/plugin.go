package pkgrev

import (
	"github.com/aghape-pkg/admin"
	"github.com/aghape/aghape"
	"github.com/aghape/db"
	"github.com/aghape/helpers"
	"github.com/aghape/plug"
)

type Plugin struct {
	plug.EventDispatcher
	db.DBNames
	admin_plugin.AdminNames
}

func (p *Plugin) OnRegister(options *plug.Options) {
	admin_plugin.Events(p).InitResources(func(e *admin_plugin.AdminEvent) {
		Admin(e.Admin, options.GetInterface(aghape.AGHAPE).(*aghape.Aghape))
	})

	db.Events(p).
		DBOnMigrate(func(e *db.DBEvent) error {
			return helpers.CheckReturnE(func() (key string, err error) {
				return "Migrate", e.AutoMigrate(&Project{}, &Version{}, &File{}).Error
			})
		})
}
