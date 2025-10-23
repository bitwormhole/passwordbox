package entity

func ListAll(prefix string) []any {

	if prefix != "" {
		theTableNamePrefix = prefix
	}

	list := make([]any, 0)

	list = append(list, new(Example))
	list = append(list, new(KeyPair))
	list = append(list, new(Password))
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

func (Password) TableName() string {
	return theTableNamePrefix + "passwords"
}

func (User) TableName() string {
	return theTableNamePrefix + "users"
}
