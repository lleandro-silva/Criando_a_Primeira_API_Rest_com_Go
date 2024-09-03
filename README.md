# Criando a Primeira API Rest com Go

A API oferece funcionalidades básicas de CRUD (criação, leitura, atualização e exclusão) para manipular informações de clientes. Os dados são armazenados em memória usando um mapa.

# Possui os seguintes endpoints:

Um para listar todos os clientes, outro para adicionar um novo cliente, um endpoint para obter os detalhes de um cliente específico pelo ID, um para atualizar um cliente existente e outro para remover um cliente pelo ID.

Para executar a API, você deve instalar o Gorilla Mux com o comando go get -u github.com/gorilla/mux, e então rodar o servidor com go run main.go. 

A API ficará disponível na porta 8080 para receber requisições.
