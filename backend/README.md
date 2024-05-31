# My BarberShop Backend

Este é o backend da aplicação My BarberShop, desenvolvido com Go e PostgreSQL.

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/ThroneCode/my-barbershop.git
   ```
2. Navegue até o diretório do projeto:
   ```bash
   cd backend
   ```
3. Instale as dependências:
   ```bash
   go mod download
   ```

## Configuração

1. Ajuste um arquivo `config.json` na raz do projeto com as configurações do seu banco postgres e sua key JWT:
   ```json
    {
      "auth": {
        "secret": "your_secret_here"
      },
      "databases": [
        {
          "driver": "postgres",
          "host": "localhost",
          "port": "5432",
          "user": "user1",
          "password": "password1",
          "dbname": "db1",
          "readonly": false
        },
        {
          "driver": "postgres",
          "host": "localhost",
          "port": "5432",
          "user": "user1",
          "password": "password1",
          "dbname": "db1",
          "readonly": true
        }
      ]
    }
   ```

## Executando o Projeto

Para iniciar o servidor, execute:

```bash
go run main.go
```

## Testes

Para executar os testes, use o comando:

```bash
go test ./...
```
