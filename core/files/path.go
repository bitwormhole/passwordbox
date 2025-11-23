package files

import "strings"

type Path string

func (p Path) String() string {
	return string(p)
}

func (p Path) Name() string {
	elist := p.Elements()
	return elist.GetLastName()
}

func (p Path) Normalize() Path {
	elist := p.Elements()
	el2, err := elist.Normalize()
	if err != nil {
		return p
	}
	return el2.Path()
}

func (p Path) IsAbsolute() bool {
	elist := p.Elements()
	return elist.IsAbsolute()
}

func (p Path) IsRelative() bool {
	elist := p.Elements()
	return elist.IsRelative()
}

func (p Path) IsFile() bool {
	meta, err := p.GetMeta()
	if err != nil {
		return false
	}
	return meta.IsFile()
}

func (p Path) IsDir() bool {
	meta, err := p.GetMeta()
	if err != nil {
		return false
	}
	return meta.IsDir()
}

func (p Path) Exists() bool {
	meta, err := p.GetMeta()
	if err != nil {
		return false
	}
	return meta.Exists()
}

func (p Path) GetParent() (parent Path, ok bool) {
	p2, err := p.Resolve("..")
	if err != nil {
		return "", false
	}
	return p2, true
}

func (p Path) GetChild(name string) Path {
	sub := "./" + name
	p2, err := p.Resolve(Path(sub))
	if err != nil {
		return p
	}
	return p2
}

func (p Path) Resolve(subpath Path) (Path, error) {

	elist2 := subpath.Elements()
	if elist2.IsAbsolute() {
		// if: as abs
		elist2n, err := elist2.Normalize()
		if err != nil {
			return "", err
		}
		return elist2n.Path(), nil
	}

	// else: as rel
	elist1 := p.Elements()
	tmp := append(elist1, elist2...)
	tmp2n, err := tmp.Normalize()
	if err != nil {
		return "", err
	}
	return tmp2n.Path(), nil
}

func (p Path) Elements() PathElementArray {

	const (
		sep1 = '\\'
		sep2 = '/'
		sep3 = '\n'
		sep  = string(sep3)
	)

	str := p.String()
	str = strings.ReplaceAll(str, string(sep1), sep)
	str = strings.ReplaceAll(str, string(sep2), sep)

	src := strings.Split(str, sep)
	dst := make(PathElementArray, 0)

	for _, item := range src {
		dst = append(dst, PathElement(item))
	}

	return dst
}

func (p Path) GetMeta() (Meta, error) {
	io := p.GetIO()
	return io.GetMeta(p)
}

func (p Path) GetIO() *IO {
	return new(IO)
}
