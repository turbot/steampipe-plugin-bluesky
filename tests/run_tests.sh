#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Initialize counters
total_tests=0
passed_tests=0
failed_tests=0
failed_files=()

# Function to run a test query
run_test() {
    local test_file=$1
    echo "Running test: $test_file"
    
    # Run the query and capture both stdout and stderr
    output=$(steampipe query "$test_file" 2>&1)
    exit_code=$?
    
    # Increment total tests
    ((total_tests++))
    
    if [ $exit_code -eq 0 ]; then
        echo -e "${GREEN}✓ Test passed${NC}"
        ((passed_tests++))
    else
        echo -e "${RED}✗ Test failed${NC}"
        echo "Error output:"
        echo "$output"
        echo "----------------------------------------"
        ((failed_tests++))
        failed_files+=("$test_file")
    fi
}

# Find all test files
test_files=$(find tests/queries -name "*.sql" | sort)

# Run each test
for test_file in $test_files; do
    run_test "$test_file"
done

# Print summary
echo -e "\nTest Summary:"
echo "----------------------------------------"
echo -e "Total tests: $total_tests"
echo -e "${GREEN}Passed: $passed_tests${NC}"
echo -e "${RED}Failed: $failed_tests${NC}"

# Print list of failed tests if any
if [ ${#failed_files[@]} -gt 0 ]; then
    echo -e "\nFailed Tests:"
    echo "----------------------------------------"
    for file in "${failed_files[@]}"; do
        echo -e "${RED}$file${NC}"
    done
fi

# Exit with non-zero status if any tests failed
if [ $failed_tests -gt 0 ]; then
    exit 1
fi 