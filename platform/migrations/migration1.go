package migrations

type Migration1 struct {
	// Migration1 için alanlar burada tanımlanır
}

func NewMigration1() *Migration1 {
	return &Migration1{}
}

func (m *Migration1) Run() {
	// Migration1 için çalıştırma işlemleri burada yer alır
}
