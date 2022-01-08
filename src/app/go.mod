module github.com/hrs-o/docker-go

go 1.16

require (
    github.com/gin-gonic/gin v1.7.4
    local.data/data v0.0.0
)

replace local.data/data => ./data