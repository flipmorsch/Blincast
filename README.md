## Teste BlinCast

Este projeto usa as seguintes tecnologias:

#### Backend
- Golang
- Gin
- Gorm
- Go Validator
- CORS

#### Frontend
- React Router
- Mantine UI

### Como rodar o backend
- Acesse a pasta do backend
- Execute o comando abaixo para instalar as dependências:
> go mod tidy
- Execute o comando abaixo para executar o projeto:
> go run main.go
- Como alternativa você pode compilar e executar o binário diretamente:
> go build .
> ./Blincast
- O servidor será executado na porta **8080** por padrão

### Como rodar o frontend
- Acesse a pasta do frontend
- Execute o comando abaixo para instalar as dependências:
> npm i
- Execute o comando abaixo para executar o projeto:
> npm run dev
- O projeto por padrão será executado na porta **5173**