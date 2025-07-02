#!/bin/sh

# Script para aguardar o banco de dados estar pronto
echo "🚀 Iniciando aplicação Clean Architecture REST API..."
echo "⏳ Aguardando o banco de dados PostgreSQL..."

# Aguarda até que o PostgreSQL esteja pronto
until pg_isready -h postgres -p 5432 -U ramon -d clean_arch_db; do
  echo "📊 PostgreSQL ainda não está pronto - aguardando..."
  sleep 3
done

echo "✅ PostgreSQL está pronto!"
echo "🔄 Iniciando a aplicação Go..."

# Executa a aplicação
exec ./main 