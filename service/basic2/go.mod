module basic2

go 1.20

require (
	github.com/gorilla/mux v1.8.1
	github.com/sourabh/db v0.0.1
	github.com/sourabh/handlers v0.0.1
)

require (
	github.com/lib/pq v1.10.9 // indirect
	github.com/sourabh/models v0.0.1 // indirect
)

replace github.com/sourabh/db v0.0.1 => ./db

replace github.com/sourabh/handlers v0.0.1 => ./handlers

replace github.com/sourabh/models v0.0.1 => ./models
