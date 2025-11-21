package entity

func ListAll(prefix string) []any {

	if prefix != "" {
		theTableNamePrefix = prefix
	}

	list := make([]any, 0)

	list = append(list, new(Example))
	list = append(list, new(KeyPair))
	list = append(list, new(PasswordChain))
	list = append(list, new(PasswordBlock))
	list = append(list, new(BinaryData))
	list = append(list, new(User))

	return list
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix = "pbox_"

////////////////////////////////////////////////////////////////////////////////

func (Example) TableName() string {
	return theTableNamePrefix + "examples"
}

func (KeyPair) TableName() string {
	return theTableNamePrefix + "key_pairs"
}

func (PasswordBlock) TableName() string {
	return theTableNamePrefix + "password_blocks"
}

func (PasswordChain) TableName() string {
	return theTableNamePrefix + "password_chains"
}

func (BinaryData) TableName() string {
	return theTableNamePrefix + "binary_data"
}

func (User) TableName() string {
	return theTableNamePrefix + "users"
}
