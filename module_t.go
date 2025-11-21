package passwordbox

const (
	theModuleName     = "github.com/bitwormhole/passwordbox"
	theModuleVersion  = "v0.0.0"
	theModuleRevision = 0
)

func GetCurrentModule() *Module {
	m := new(Module)
	m.name = theModuleName
	m.version = theModuleVersion
	m.revision = theModuleRevision
	return m
}
