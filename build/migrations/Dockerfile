FROM gomicro/goose:3.7.0

ADD ./migrations/*.sql .

ENV GOOSE_DRIVER="postgres"
ENV GOOSE_DBSTRING="postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"

CMD goose up
