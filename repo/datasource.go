package repo

// DataSource is the interface for datasources
type DataSource interface {
	Repo
	Close()
}
