# recvcheck-lint

Custom linter to check that the name of the method receivers match the type of the receiver (e.g., the name of teh receiver of type `MyStruct` should be `m`).

# Reference

 - [Add new linters to `golangci-lint`](https://golangci-lint.run/contributing/new-linters/)
 - [Use `go/analysis` to write a new linter](https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/)
 - [Example linter with `nocopy`](https://github.com/cockroachlabs/release-staging/tree/d89a1b98b590923f708c5fd06cea0c3aa074ac02/pkg/testutils/lint/passes/nocopy)
