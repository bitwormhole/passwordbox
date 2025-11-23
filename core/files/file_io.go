package files

import (
	"io/fs"
	"os"
	"time"
)

type IO struct {
}

func (inst *IO) innerPrepareOptions(opt *Options, method rune) *Options {

	if opt == nil {
		opt = new(Options)
	}

	if opt.Flag == 0 {
		switch method {
		case 'w':
			opt.Flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
		case 'r':
			opt.Flag = os.O_RDONLY
		case 'm':
			opt.Flag = os.O_CREATE
		}
	}

	if opt.Mode == 0 {
		switch method {
		case 'w':
			opt.Mode = 0644
		case 'r':
			opt.Mode = 0
		case 'm':
			opt.Mode = 0755
		}
	}

	return opt
}

func (inst *IO) Mkdir(dir Path, opt *Options) error {

	opt = inst.innerPrepareOptions(opt, 'm')

	name := dir.String()
	perm := opt.Mode
	return os.Mkdir(name, perm)
}

func (inst *IO) Mkdirs(dir Path, opt *Options) error {

	opt = inst.innerPrepareOptions(opt, 'm')

	name := dir.String()
	perm := opt.Mode
	return os.MkdirAll(name, perm)

}

func (inst *IO) ReadText(file Path, opt *Options) (string, error) {
	name := file.String()
	data, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (inst *IO) ReadBinary(file Path, opt *Options) ([]byte, error) {
	name := file.String()
	return os.ReadFile(name)
}

func (inst *IO) WriteText(text string, file Path, opt *Options) error {

	opt = inst.innerPrepareOptions(opt, 'w')

	name := file.String()
	perm := opt.Mode
	data := []byte(text)

	return os.WriteFile(name, data, perm)
}

func (inst *IO) WriteBinary(data []byte, file Path, opt *Options) error {

	opt = inst.innerPrepareOptions(opt, 'w')

	name := file.String()
	perm := opt.Mode
	// data := []byte(text)

	return os.WriteFile(name, data, perm)
}

func (inst *IO) GetMeta(path Path) (Meta, error) {
	name := path.String()
	info, err := os.Stat(name)
	if err != nil {
		// return nil, err
		mock := new(innerFileMetaMock)
		return mock, nil
	}
	ada := new(innerFileMetaAdapter)
	ada.info = info
	return ada, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerFileMetaAdapter struct {
	info fs.FileInfo
}

// Length implements Meta.
func (i *innerFileMetaAdapter) Length() int64 {
	return i.info.Size()
}

// Time implements Meta.
func (i *innerFileMetaAdapter) UpdatedAt() time.Time {
	return i.info.ModTime()
}

// Exists implements Meta.
func (i *innerFileMetaAdapter) Exists() bool {
	return true
}

// IsDir implements Meta.
func (i *innerFileMetaAdapter) IsDir() bool {
	return i.info.IsDir()
}

// IsFile implements Meta.
func (i *innerFileMetaAdapter) IsFile() bool {
	return !i.info.IsDir()
}

// IsLink implements Meta.
func (i *innerFileMetaAdapter) IsLink() bool {
	panic("unimplemented")
}

// Mode implements Meta.
func (i *innerFileMetaAdapter) Mode() Mode {
	return i.info.Mode()
}

////////////////////////////////////////////////////////////////////////////////

type innerFileMetaMock struct {
}

// Exists implements Meta.
func (i *innerFileMetaMock) Exists() bool {
	return false
}

// IsDir implements Meta.
func (i *innerFileMetaMock) IsDir() bool {
	return false
}

// IsFile implements Meta.
func (i *innerFileMetaMock) IsFile() bool {
	return false
}

// IsLink implements Meta.
func (i *innerFileMetaMock) IsLink() bool {
	return false
}

// Length implements Meta.
func (i *innerFileMetaMock) Length() int64 {
	return 0
}

// Mode implements Meta.
func (i *innerFileMetaMock) Mode() Mode {
	return 0
}

// UpdatedAt implements Meta.
func (i *innerFileMetaMock) UpdatedAt() time.Time {
	return time.UnixMilli(0)
}

////////////////////////////////////////////////////////////////////////////////
