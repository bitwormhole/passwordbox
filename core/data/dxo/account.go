package dxo

type AccountType string

const (
	AccountTypeMaster AccountType = "master" // 表示主账号
	AccountTypeSlave  AccountType = "slave"  // 表示被托管的账号
)
