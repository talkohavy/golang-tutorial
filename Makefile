.PHONY: progress test solutions scratch notes

# Show your progress. Add topic=03_loops for an exercise-by-exercise view.
progress:
	@go run ./cmd/progress $(topic)

# Run the tests of one topic with full failure details: make test topic=03_loops
test:
	@if [ -z "$(topic)" ]; then echo "usage: make test topic=03_loops"; exit 1; fi
	go test ./exercises/$(topic)/...

# Check that the reference solutions pass every test.
solutions:
	@go run ./cmd/progress -solutions

scratch:
	@go run ./scratch

notes:
	@go run ./notes
