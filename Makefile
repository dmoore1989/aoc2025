get_day:
	go run main.go $(filter-out $@,$(MAKECMDGOALS))

run_pt1_sample:
	DAY=$(word 2,$(MAKECMDGOALS)) && go run cmd/day$$DAY/main.go sample 1

run_pt2_sample:
	DAY=$(word 2,$(MAKECMDGOALS)) && go run cmd/day$$DAY/main.go sample 2

run_pt1_real:
	DAY=$(word 2,$(MAKECMDGOALS)) && go run cmd/day$$DAY/main.go real 1

run_pt2_real:
	DAY=$(word 2,$(MAKECMDGOALS)) && go run cmd/day$$DAY/main.go real 2

%:
	@:
