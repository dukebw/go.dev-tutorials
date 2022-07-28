module github.com/dukebw/go.dev-tutorials/create-module/hello

go 1.18

require github.com/dukebw/go.dev-tutorials/create-module/greetings v0.0.0-20220727205344-794b525d2f1e

replace github.com/dukebw/go.dev-tutorials/create-module/greetings => ../greetings
