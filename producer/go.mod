module github.com/sriharibtl/producer

go 1.15

require (
    github.com/gorilla/mux v1.8.0
    httpserver v0.0.0
)

replace (
    httpserver v0.0.0 => ./httpserver
)
