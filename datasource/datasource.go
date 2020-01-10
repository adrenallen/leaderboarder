package datasource

// DataSource is an interface defining 
// what must be present for a potential back-end storage
type DataSource interface {
	GetAll() []Entry
	SaveNew(entry Entry)
}
