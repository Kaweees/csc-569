alias r := run
alias b := bench
alias cd := client_dev

# Run the project
run:
  go run .

# Run the project with hot reloading
client_dev:
  reflex -r '\.go$$' -s go run client.go

# Benchmark the project
bench:
  go test -bench=. -benchmem -cpu=1