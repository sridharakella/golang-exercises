package main

type dbcontract interface {
	close()
	InsertUser(username string) error
	SelectSingleUser(username string) (string, error)
}

type Application struct {
	db dbcontract
}

func NewApplication(db dbcontract) *Application {
	return &Application{db: db}
}

func main() {
	// implementation logic
	/*db, err := mysqldb.New()
	if err != nil {
		panic(err)
	}
	app := NewApplication(db)
	app.db.InsertUser("user1")*/
}
