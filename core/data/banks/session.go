package banks

import "io"

type Session interface {
	io.Closer

	// objects

	FetchObject(o *Object) error

	PutObject(o *Object) error

	// refs

	FetchRef(r *Ref) error

	PutRef(r *Ref) error

	// configs

	FetchConfiguration(c *Configuration) error

	PutConfiguration(c *Configuration) error
}
