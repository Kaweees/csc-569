alias r := run
alias b := bench

# Run the project
run:
  go run .

# Benchmark the project
bench:
  go test -bench=. -benchmem -cpu=1