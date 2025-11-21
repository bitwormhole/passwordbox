package dto

import "github.com/bitwormhole/passwordbox/app/data/dxo"

type User struct {
	ID dxo.UserID `json:"id"`

	Base

	Name       string           `json:"name"`   // the user-name
	Email      dxo.EmailAddress `json:"email"`  // the email-address of user
	ProviderDN dxo.DomainName   `json:"domain"` // the domain-name of online-account provider
}
