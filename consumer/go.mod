module github.com/sriharibtl/consumer

go 1.15

require (
    httpserver v0.0.0
)

replace (
    httpserver v0.0.0 => ./httpserver
)