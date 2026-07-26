mod app 'app'

set default-list := true

# Run code generation and start dev server
dev:
    just app dev

# Run code generation and build the app
build:
    just app build

# Run all linters and type checkers
check:
    just app check

# Format all source code
fmt:
    just app fmt

# Run all tests
test:
    just app test
