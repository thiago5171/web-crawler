# Web crawler Backend

## O que é o projeto?

Esse é um projeto com o proposito de navegar pelos links encontrado e armazena-los em uma base em postgres. Ademais, o projeto utiliza a arquitetura hexagonal e até o momento tem uma única porta de entrada: um projeto BackEnd que utiliza o Framework [Echo](https://echo.labstack.com/).

 
### **Guia rápido de execução**

1. Execute o comando `go mod tidy` para baixar as dependências do projeto;
2. Copie todo o conteúdo do arquivo [src/ui/api/app/.env.example](src/ui/api/app/.env.example) e cole em um novo arquivo chamado `.env` na mesma pasta ([src/ui/api/app/](src/ui/api/app/));
3. Execute o banco de dados e instância redis com o seguinte comando: `docker compose up database  --build -d`

Pronto! O projeto está configurado. A partir de agora, toda vez que quiser iniciar o projeto basta executar o comando `go run main.go` dentro da pasta `src/ui/api/app`. Assim, o projeto estará disponível no endereço `http://localhost:8000`.

 