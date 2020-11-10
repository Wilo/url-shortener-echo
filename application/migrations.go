package app

// PersistenceMigration interface definition to handle persistence migrations
type PersistenceMigration interface {
	Migrate() error
}
