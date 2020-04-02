package db

type App struct {
	Appname  string
	Password string
}

func (app *App) CreateApp() error {

	return createapp(app.Appname, app.Password)

}

func (app *App) Authenticate() bool {

	password, err := getapp(app.Appname)

	if err != nil {
		return false
	}

	if password == app.Password {
		return true
	}

	return false

}
