create readme.md with code

gofiber-boilerplate

Layer of app
--- 
src
    api
        configs
        handlers
        helpers
        middleware
        models
        routes
    .env
    go.mod
    go.sum
    main.go
.gitignore
...dockerfile
readme.md
test.http

Layer 1
--- 
    src
    .gitignore
    ...dockerfile
    readme.md
    test.http

Layer 2
--- 
    api
    .env
    go.mod
    go.sum
    main.go

Layer 3
--- 
    configs
    handlers
    helpers
    middleware
    models
    routes