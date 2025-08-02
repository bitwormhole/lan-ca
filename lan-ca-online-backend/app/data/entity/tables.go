package entity

import "github.com/bitwormhole/lan-ca/backend/app/data/dxo"

// GetDataGroupInfo  ...
func GetDataGroupInfo() dxo.DataGroupInfo {
	return new(dgInfo)
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix string

type dgInfo struct{}

func (inst *dgInfo) Prototypes() []any {

	list := make([]any, 0)

	list = append(list, new(Domain))
	list = append(list, new(Certificate))
	list = append(list, new(Scene))

	return list
}

func (inst *dgInfo) SetTableNamePrefix(prefix string) {
	if theTableNamePrefix == "" {
		theTableNamePrefix = prefix
	}
}

////////////////////////////////////////////////////////////////////////////////

// TableName 。。。
func (Example) TableName() string {
	return theTableNamePrefix + "examples"
}

// TableName 。。。
func (Domain) TableName() string {
	return theTableNamePrefix + "domains"
}

// TableName 。。。
func (Scene) TableName() string {
	return theTableNamePrefix + "scenes"
}

// TableName 。。。
func (Certificate) TableName() string {
	return theTableNamePrefix + "certificates"
}
