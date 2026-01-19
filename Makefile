query: sui_mainnet_introspection_query.graphql
	./generate_sui_schema.sh

types:
	go run github.com/99designs/gqlgen generate
