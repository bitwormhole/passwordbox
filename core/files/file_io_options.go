package files

import "io/fs"

type Options struct {
	Mode fs.FileMode

	Flag int
}
