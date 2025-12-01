package backend

import (
	"time"

	aesd "github.com/bitwormhole/passwordbox/core/algorithms/aes.d"
	rsad "github.com/bitwormhole/passwordbox/core/algorithms/rsa.d"
	sha1d "github.com/bitwormhole/passwordbox/core/algorithms/sha1.d"
	sha256d "github.com/bitwormhole/passwordbox/core/algorithms/sha256.d"
	"github.com/bitwormhole/passwordbox/core/boot"
	"github.com/bitwormhole/passwordbox/core/components"
	"github.com/bitwormhole/passwordbox/core/pemapi"
)

////////////////////////////////////////////////////////////////////////////////
// main

type config struct {
}

func (inst *config) configure(target *boot.Configuration) {
	inst.configureEnvironment(target)
	inst.configureProperties(target)
	inst.configureComponents(target)

	target.Timeout = time.Second * 300
}

func (inst *config) configureProperties(target *boot.Configuration) {

	sub := new(innerPropertiesConfig)
	sub.target = target

	sub.set("a", "1")
	sub.set("b", "2")
	sub.set("c", "3")

}

func (inst *config) configureEnvironment(target *boot.Configuration) {

	sub := new(innerEnvironmentConfig)
	sub.target = target

	sub.set("EA", "1x")
	sub.set("EB", "2x")
	sub.set("EC", "3x")

}

func (inst *config) configureComponents(target *boot.Configuration) {

	sub := new(innerComponentsConfig)
	sub.target = target

	sub.register(sub.sha1driver)
	sub.register(sub.sha256driver)
	sub.register(sub.router)

}

////////////////////////////////////////////////////////////////////////////////
// properties

type innerPropertiesConfig struct {
	target *boot.Configuration
}

func (inst *innerPropertiesConfig) set(name, value string) {
	inst.target.SetProperty(name, value)
}

////////////////////////////////////////////////////////////////////////////////
// environment

type innerEnvironmentConfig struct {
	target *boot.Configuration
}

func (inst *innerEnvironmentConfig) set(name, value string) {
	inst.target.SetEnv(name, value)
}

////////////////////////////////////////////////////////////////////////////////
// components

type innerComponentsConfig struct {
	target *boot.Configuration
}

func (inst *innerComponentsConfig) register(fn func() *components.ComponentRegistration) {
	inst.target.RegisterComponentFn(fn)
}

func (inst *innerComponentsConfig) sha1driver() *components.ComponentRegistration {

	r := new(components.ComponentRegistration)

	r.ID = "sha1"
	r.Scope = components.ScopeSingleton
	r.SetAliases("sha-1")
	r.SetClasses("hash", "algorithm")

	r.OnNew = func() any {
		return new(sha1d.Driver)
	}

	r.OnLoad = func(l *components.Loading) {
	}

	return r
}

func (inst *innerComponentsConfig) sha256driver() *components.ComponentRegistration {

	r := new(components.ComponentRegistration)

	r.ID = "sha256"
	r.Scope = components.ScopeSingleton
	r.SetAliases("sha-256")
	r.SetClasses("hash", "algorithm")

	r.OnNew = func() any {
		return new(sha256d.Driver)
	}

	r.OnLoad = func(l *components.Loading) {
	}

	return r
}

func (inst *innerComponentsConfig) AesDriver() *components.ComponentRegistration {

	r := new(components.ComponentRegistration)

	r.ID = "aes"
	r.Scope = components.ScopeSingleton
	r.SetAliases("AES")
	// r.SetClasses("hash", "algorithm")

	r.OnNew = func() any {
		return new(aesd.Driver)
	}

	r.OnLoad = func(l *components.Loading) {
	}

	return r
}

func (inst *innerComponentsConfig) RsaDriver() *components.ComponentRegistration {

	r := new(components.ComponentRegistration)

	r.ID = "rsa"
	r.Scope = components.ScopeSingleton
	r.SetAliases("RSA")
	// r.SetClasses("hash", "algorithm")

	r.OnNew = func() any {
		return new(rsad.Driver)
	}

	r.OnLoad = func(l *components.Loading) {
	}

	return r
}

func (inst *innerComponentsConfig) router() *components.ComponentRegistration {

	r := new(components.ComponentRegistration)

	r.ID = "router"
	r.Scope = components.ScopeSingleton
	r.SetAliases("pemapi-router")
	r.SetClasses("router", "pemapi")

	r.OnNew = func() any {
		return new(pemapi.Router)
	}

	r.OnLoad = func(l *components.Loading) {

		sub := l.Component.(*pemapi.Router)
		sub.Controllers = nil

		list := sub.Controllers
		all := l.Loader.ListAllComponents()
		for _, com := range all {
			ctrl, ok := com.(pemapi.Controller)
			if ok {
				list = append(list, ctrl)
			}
		}
		sub.Controllers = list

	}

	return r
}

////////////////////////////////////////////////////////////////////////////////
// EOF
