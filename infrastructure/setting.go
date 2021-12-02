package infrastructure

type app struct {
	AppName  string
	RunMode  string
	HTTPPort int
	LogPath  string
}

var AppSetting = &app{
	AppName:  "bookshop",
	RunMode:  "debug",
	HTTPPort: 8000,
	LogPath:  "/home/chuyang/tmp",
}

type database struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
}

var DatabaseSetting = &database{
	DBUser:     "postgres",
	DBPassword: "postgres",
	DBHost:     "127.0.0.1",
	DBPort:     5432,
	DBName:     "bookshop",
}

type ldapSetting struct {
	Url      string
	User     string
	Password string
	DC       string //domain component
}

var LDAPSetting = &ldapSetting{
	Url:      "ldap://127.0.0.1:10389",
	User:     "uid=admin,ou=system",
	Password: "secret",
	DC:       "dc=chuyang,dc=org",
}
