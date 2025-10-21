package passwordbox

import (
	"embed"

	"github.com/bitwormhole/passwordbox/gen/main4pwbox"
	"github.com/bitwormhole/passwordbox/gen/test4pwbox"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"

	// "github.com/starter-go/module-gorm-sqlite/modules/sqlite"

	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
)

////////////////////////////////////////////////////////////////////////////////

const (
	theModuleName     = "github.com/bitwormhole/passwordbox"
	theModuleVersion  = "v0.0.0"
	theModuleRevision = 0
)

////////////////////////////////////////////////////////////////////////////////

const (
	theSrcMainResPath = "src/main/resources"
	theSrcTestResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theSrcMainResFS embed.FS

//go:embed "src/test/resources"
var theSrcTestResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

func Module() application.Module {

	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#main")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theSrcMainResFS, theSrcMainResPath)

	mb.Components(main4pwbox.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

func ModuleForTest() application.Module {

	mb := new(application.ModuleBuilder)
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theSrcTestResFS, theSrcTestResPath)

	mb.Components(test4pwbox.ExportComponents)

	mb.Depend(Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
