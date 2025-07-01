package database

import (
	"fmt"
	"log"

	"clean-arch-rest/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresConnection cria uma nova conexão com o PostgreSQL
func NewPostgresConnection() *gorm.DB {
	dsn := "host=localhost user=ramon password=1234 dbname=clean_arch_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	// Executa as migrações automaticamente
	err = db.AutoMigrate(&domain.Order{})
	if err != nil {
		log.Fatal("Erro ao executar migrações:", err)
	}

	fmt.Println("Banco de dados conectado com sucesso")
	return db
}
