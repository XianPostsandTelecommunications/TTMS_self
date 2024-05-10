package setting

type group struct {
	Log    log
	Va     va
	Dao    mdao
	Worker worker
	Maker  maker
	Auto   auto
}

var Group = new(group)

func AllInit() {
	Group.Log.Init()
}
