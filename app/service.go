package app

type AppService struct {
	//dbcon goes here
	//roles/permissions service gets added here
	//etc
}

func New() *AppService {
	return &AppService{} //normally has things like the auth service instantiated to validate specific permissions (seperate from middleware), database connections, mail clients, etc
}
