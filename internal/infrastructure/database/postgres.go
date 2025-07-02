package database

import (
	"fmt"
	"log"
	"os"

	"clean-arch-rest/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresConnection cria uma nova conexão com o PostgreSQL
func NewPostgresConnection() *gorm.DB {
	// Obtém as variáveis de ambiente com fallback para valores padrão
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "ramon")
	password := getEnv("DB_PASSWORD", "1234")
	dbname := getEnv("DB_NAME", "clean_arch_db")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port)

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

// getEnv retorna o valor da variável de ambiente ou o valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
