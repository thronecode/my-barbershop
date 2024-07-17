
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
   MYBARBERSHOP_AUTH_SECRET=your_secret_here
   MYBARBERSHOP_DB1_DRIVER=postgres
   MYBARBERSHOP_DB1_HOST=localhost
   MYBARBERSHOP_DB1_PORT=5432
   MYBARBERSHOP_DB1_USER=user1
   MYBARBERSHOP_DB1_PASSWORD=password1
   MYBARBERSHOP_DB1_DBNAME=db1
   MYBARBERSHOP_DB1_READONLY=false
   MYBARBERSHOP_DB2_DRIVER=postgres
   MYBARBERSHOP_DB2_HOST=localhost
   MYBARBERSHOP_DB2_PORT=5432
   MYBARBERSHOP_DB2_USER=user1
   MYBARBERSHOP_DB2_PASSWORD=password1
   MYBARBERSHOP_DB2_DBNAME=db1
   MYBARBERSHOP_DB2_READONLY=true
   ```

## Executando o Projeto

Para iniciar o servidor, execute:

```bash
cd cmd/server && go run main.go
```

## Testes

Para executar os testes, use o comando:

```bash
go test ./...
```