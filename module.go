package passwordbox

import "strconv"

type Module struct {
	name     string
	version  string
	revision int
}

func (m Module) String() string {
	name := m.name
	ver := m.version
	rev := m.revision
	return "module:" + name + "#" + ver + "-r" + strconv.Itoa(rev)
}

func (m Module) GetName() string {
	return m.name
}

func (m Module) GetVersion() string {
	return m.version
}

func (m Module) GetRevision() int {
	return m.revision
}
