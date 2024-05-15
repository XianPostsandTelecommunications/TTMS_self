package setting

type mdao struct {
}

func (d *mdao) Init() {
	mysql.InitMysql()

}
