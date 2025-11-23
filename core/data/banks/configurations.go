package banks

import "github.com/bitwormhole/passwordbox/core/data/properties"

// ConfigName 表示配置文件的文件名
type ConfigName string

// Configurations 表示由多个配置文件构成的集合, 也就是一个配置文件夹
type Configurations interface {
	Get(name ConfigName) ConfigurationHolder
}

type ConfigurationHolder interface {
	GetName() ConfigName

	Exists() bool

	Fetch() (*Configuration, error)

	Put(c *Configuration) error
}

type Configuration struct {
	Name ConfigName

	props properties.Table
}
