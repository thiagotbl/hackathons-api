package interfaces

type IDb interface {
	Query(statement string) (IRows, error)
}

type IRows interface {
	Scan(dest ...interface{}) error
	Next() bool
}
