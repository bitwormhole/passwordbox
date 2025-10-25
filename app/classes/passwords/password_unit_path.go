package passwords

// func MakeUniquePathWithDTO(item *dto.Password) dxo.UniformPath {
// 	mk := new(innerUniquePathMaker)
// 	mk.initWithDTO(item)
// 	return mk.create()
// }

// func MakeUniquePathWithEntity(item *entity.PasswordRef) dxo.UniformPath {
// 	mk := new(innerUniquePathMaker)
// 	mk.initWithEntity(item)
// 	return mk.create()
// }

// ////////////////////////////////////////////////////////////////////////////////

// type innerUniquePathMaker struct {
// 	builder strings.Builder
// }

// func (inst *innerUniquePathMaker) create() dxo.UniformPath {
// 	builder := &inst.builder
// 	str := builder.String()
// 	return dxo.UniformPath(str)
// }

// func (inst *innerUniquePathMaker) addElement(el string) {

// 	const sep = '/'

// 	el = strings.ReplaceAll(el, string(sep), "-")
// 	el = strings.TrimSpace(el)

// 	if len(el) < 1 {
// 		el = "undefined"
// 	}

// 	builder := &inst.builder
// 	builder.WriteRune(sep)
// 	builder.WriteString(el)
// }

// func (inst *innerUniquePathMaker) init(it *dto.Password) {

// 	rev := it.Revision.String()

// 	inst.addElement(it.Email.String())
// 	inst.addElement(it.Domain1.String())
// 	inst.addElement(it.Domain2.String())
// 	inst.addElement(it.UserName)
// 	inst.addElement(it.Scene)
// 	inst.addElement(rev)
// }

// func (inst *innerUniquePathMaker) initWithDTO(it1 *dto.Password) {
// 	it2 := new(dto.Password)
// 	if it1 != nil {
// 		*it2 = *it1
// 	}
// 	inst.init(it2)
// }

// func (inst *innerUniquePathMaker) initWithEntity(it1 *entity.PasswordRef) {
// 	it2 := new(dto.Password)
// 	if it1 != nil {
// 		ConvertE2D(it1, nil, it2)
// 	}
// 	inst.init(it2)
// }
