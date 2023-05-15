package migrations

type Migration2 struct {
	// Migration2 için alanlar burada tanımlanır
}

func NewMigration2() *Migration2 {
	return &Migration2{}
}

func (m *Migration2) Run() {
	// Migration2 için çalıştırma işlemleri burada yer alır
}
