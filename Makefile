POSTGRES_CONTAINER_NAME=my-postgres
POSTGRES_PASSWORD=senha123
POSTGRES_PORT=5432
POSTGRES_DB=userdb

postgres-up:
	docker run --name $(POSTGRES_CONTAINER_NAME) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-p $(POSTGRES_PORT):5432 \
		-d --rm postgres

postgres-down:
	docker stop $(POSTGRES_CONTAINER_NAME)


docker-sqlc:
	docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate

# --rm 					Remove automaticamente o container ao parar 
# -p 5432:5432  		Porta externa 5432 -> interna 5432
# POSTGRES_PASSWORD		Define senha do postgres
# POSTGRES_DB 			Define o nome do banco
# make postgres-up		Sobe o banco
# make postgres-down	Derruba o banco e ele desaparece

# Testa entrado no banco

# docker exec -it my-postgres psql -U postgres -d userdb
