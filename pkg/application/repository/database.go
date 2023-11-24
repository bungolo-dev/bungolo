package repository

type DbConnection interface {
	Initialize() error
	RoomsRepository
}
