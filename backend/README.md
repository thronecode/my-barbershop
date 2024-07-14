
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

1. Defina as variáveis de ambiente no seu sistema ou crie um arquivo `.env` na raiz do projeto com as configurações do seu banco postgres e sua key JWT:

   ```env
   AUTH_SECRET=your_secret_here
   DB1_DRIVER=postgres
   DB1_HOST=localhost
   DB1_PORT=5432
   DB1_USER=user1
   DB1_PASSWORD=password1
   DB1_DBNAME=db1
   DB1_READONLY=false
   DB2_DRIVER=postgres
   DB2_HOST=localhost
   DB2_PORT=5432
   DB2_USER=user1
   DB2_PASSWORD=password1
   DB2_DBNAME=db1
   DB2_READONLY=true
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