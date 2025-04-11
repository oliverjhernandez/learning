export PGPASSWORD='secret'
psql -U postgres -d greenlight -h localhost -p 5432 -c "DROP TABLE movies"
psql -U postgres -d greenlight -h localhost -p 5432 -c "DROP TABLE schema_migrations"
migrate -path=./migrations -database='postgres://postgres:secret@localhost/greenlight?sslmode=disable' up

BODY='{"title":"Moana","year":2016,"runtime":"107 mins", "genres":["animation","adventure"]}'
curl -d "$BODY" localhost:3000/v1/movie

BODY='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["action","adventure"]}'
curl -d "$BODY" localhost:3000/v1/movie

BODY='{"title":"Deadpool","year":2016, "runtime":"108 mins","genres":["action","comedy"]}'
curl -d "$BODY" localhost:3000/v1/movie

BODY='{"title":"The Breakfast Club","year":1986, "runtime":"96 mins","genres":["drama"]}'
curl -d "$BODY" localhost:3000/v1/movie

BODY='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["sci-fi","action","adventure"]}'
curl -X PUT -d "$BODY" localhost:3000/v1/movie/2

BODY='{"email": "alice@example.com" , "password": "pa55word"}'
TOKEN=$(curl -d "$BODY" localhost:4000/v1/tokens/authentication | jq -r '.authentication_token.token')
