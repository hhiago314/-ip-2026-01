package utils

// Importa os pacotes necessários para a conexão com o banco de dados
import (
    "database/sql" // Usado para interagir com o banco de dados
    "fmt"          // Usado para formatar strings
    "log"          // Usado para registrar mensagens de log e erros
    "os"           // Usado para acessar variáveis de ambiente
    "strings"      // Usado para limpar espaços das variáveis de ambiente

    "github.com/joho/godotenv" // Usado para carregar variáveis de ambiente de um arquivo .env
    _ "github.com/lib/pq"      // Driver PostgreSQL para o pacote database/sql
)

// Declara uma variável global para armazenar a conexão com o banco de dados
var DB *sql.DB

// Função responsável por conectar ao banco de dados
func ConnectToDB() {
    // Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load(".env")
    if err != nil {
        log.Println("Aviso: arquivo .env não encontrado. Tentando usar variáveis de ambiente do sistema.")
    }

    // Obtém as variáveis de ambiente necessárias para a conexão
    user := strings.TrimSpace(os.Getenv("DB_USER"))         // Usuário do banco de dados
    password := strings.TrimSpace(os.Getenv("DB_PASSWORD")) // Senha do banco de dados
    dbname := strings.TrimSpace(os.Getenv("DB_NAME"))       // Nome do banco de dados
    host := strings.TrimSpace(os.Getenv("DB_HOST"))         // Host do banco de dados
    port := strings.TrimSpace(os.Getenv("DB_PORT"))         // Porta do banco de dados

    if user == "" || password == "" || dbname == "" || host == "" || port == "" {
        log.Fatalf("Erro: faltam variáveis de ambiente de conexão ao banco de dados. Verifique o arquivo .env ou as variáveis de ambiente do sistema.")
    }

    // Cria a string de conexão com base nas variáveis de ambiente
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

	// Abre a conexão com o banco de dados usando a string de conexão
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		// Encerra o programa caso ocorra um erro ao abrir a conexão
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Testa a conexão com o banco de dados
	err = DB.Ping()
	if err != nil {
		// Encerra o programa caso ocorra um erro ao verificar a conexão
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	// Imprime uma mensagem de sucesso no terminal caso a conexão seja estabelecida
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}
