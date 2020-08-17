package env


type Mysql struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"dbname"`
}
