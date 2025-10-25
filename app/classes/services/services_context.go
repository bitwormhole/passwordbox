package services

import (
	"github.com/bitwormhole/passwordbox/app/classes/authx"
	"github.com/bitwormhole/passwordbox/app/classes/binarydata"
	"github.com/bitwormhole/passwordbox/app/classes/hubproviders"
	"github.com/bitwormhole/passwordbox/app/classes/keypairs"
	"github.com/bitwormhole/passwordbox/app/classes/passwords"
)

type Context struct {
	SS Service

	PasswordService      passwords.Service
	PasswordChainService passwords.ChainService
	PasswordBlockService passwords.BlockService
	PasswordFastService  passwords.FastService

	KeyPairService keypairs.Service

	AuthService authx.Service

	HubProviderService hubproviders.Service

	BinaryDataService binarydata.Service
}
