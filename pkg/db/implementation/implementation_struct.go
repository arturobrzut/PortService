package implementation

type DbParam int

const (
	DbType DbParam = iota
	DbUrl
	DbName
	DbUser
	DbPassword
	DbCollection
)
