module kafka

go 1.15

require (
    github.com/segmentio/kafka-go v0.4.12
    cbapi v0.0.0
)

replace (
    cbapi v0.0.0 => ../cbapi
)
