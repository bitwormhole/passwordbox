package files

import "time"

type Meta interface {
	UpdatedAt() time.Time

	Length() int64

	Mode() Mode

	IsDir() bool

	IsFile() bool

	IsLink() bool

	Exists() bool
}
