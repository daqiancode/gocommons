package states

const (
	OK             = 0
	InternalError  = 1 //eg. sql error
	ServiceError   = 2
	InvalidParam   = 3
	NoAuthority    = 4
	Forbidden      = 5
	LoginRequired  = 6
	NotImplemented = 7
	NotFound       = 8 // No record with corresponding id found
)
