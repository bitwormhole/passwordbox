package entity

func ListAll(prefix string) []any {

	if prefix != "" {
		theTableNamePrefix = prefix
	}

	list := make([]any, 0)

	list = append(list, new(Example))
	list = append(list, new(Password))

	return list
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix = "passwordbox_"

////////////////////////////////////////////////////////////////////////////////

func (Example) TableName() string {
	return theTableNamePrefix + "examples"
}

func (Password) TableName() string {
	return theTableNamePrefix + "passwords"
}
