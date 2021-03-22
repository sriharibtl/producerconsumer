module httpserver

go 1.15

require (
    github.com/gorilla/mux v1.8.0
    kafka v0.0.0
    cbapi v0.0.0
)

replace (
    kafka v0.0.0 => ../kafka
    cbapi v0.0.0 => ../cbapi
)
