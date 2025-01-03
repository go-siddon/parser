package parser

const (
	databaseTag = "db"
)

const (
	propertyRequired = "required"
	propertyIndex    = "index"
	propertyUnique   = "unique"
)

type E struct {
	Key      string
	Value    interface{}
	Required bool
	Unique   bool
	Index    bool
}

type M []E
