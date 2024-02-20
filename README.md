

Executa docker compose

 - docker-compose up -d --remove-orphans 

Cria o arquivo para salvar o último valor dolar
 - sqlite3 ./data/dolar_brl.db ".databases" 

Script para criar banco de dados
 - sqlite3 ./data/dolar_brl.db "CREATE TABLE dolar_brl (id VARCHAR(255) PRIMARY KEY,price VARCHAR(25), create_at VARCHAR(150);"  

Script para listar dados 
 - sqlite3 ./data/dolar_brl.db "SELECT * FROM dolar_brl;"  

Executar server - go run Server/server.go
 - http://localhost:8080/cotacao

Executar server - go run Server/server.go
- http://localhost:8080/cotacao

Executar client - go run Client/client.go
- http://localhost:8090/cotacao?code=BRL
