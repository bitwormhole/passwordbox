package banks

import "github.com/bitwormhole/passwordbox/core/files"

type StoreContext struct {
	Dir files.Path

	Store Store
}

type BankContext struct {
	Dir files.Path

	Name BankName

	Bank Bank
}

type ObjectsContext struct {
	OwnerBank *BankContext

	Dir files.Path

	Objects Objects
}

type RefsContext struct {
	OwnerBank *BankContext

	Dir files.Path

	Refs Refs
}

type ConfigurationsContext struct {
	OwnerBank *BankContext

	Dir files.Path
}
