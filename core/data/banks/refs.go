package banks

type Refname string

type Refs interface {
	Get(name Refname) RefHolder
}

type RefHolder interface {
	GetName() Refname

	Fetch() (*Ref, error)

	Put(ref *Ref) error

	Exists() bool
}

type Ref struct {
	Name   Refname
	Target ObjectID
}
