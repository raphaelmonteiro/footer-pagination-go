package main

import (
	"os"
	"strings"
	"testing"
)

func TestCheckVariablesFunctionIfAllVariablesIsNil(t *testing.T) {
	_, _, _, _, error := checkVariables()

	if len(error) != 4 {
		t.Errorf("It was expected to obtain error of the 4 expected variables because none was defined")
	}
}

func TestCheckVariablesFunctionIfCurrentPageIsBiggestThanTotalPages(t *testing.T) {
	os.Setenv("current_page", "12")
	os.Setenv("total_pages", "10")
	os.Setenv("boundaries", "2")
	os.Setenv("around", "2")

	_, _, _, _, error := checkVariables()

	if !contain(error, "current_page") {
		t.Errorf("The current page %v must not be larger than the total pages %v.", os.Getenv("current_page"), os.Getenv("total_pages"))
	}
}

func TestCheckVariablesFunctionIfAllParamsHaveBeenPassed(t *testing.T) {
	os.Setenv("current_page", "4")
	os.Setenv("total_pages", "10")
	os.Setenv("boundaries", "2")
	os.Setenv("around", "2")

	currentPage, totalPages, boundaries, around, error := checkVariables()

	if (contain(error, "current_page") || currentPage != 4) ||
		(contain(error, "total_pages") || totalPages != 10) ||
		(contain(error, "boundaries") || boundaries != 2) ||
		(contain(error, "around") || around != 2) {
		t.Errorf("The %v parameters are not in accordance with the expected result.", strings.Join(error, ", "))
	}
}

func TestFooterPaginationFunctionIfAllReceivedValuesAreExpected(t *testing.T) {
	result := footerPagination(3, 5, 1, 2)
	if result != "1 2 3 4 5" {
		t.Errorf("The expected value was '1 2 3 4 5 6 7 ... 9 10', but we received %v ", result)
	}

	result = footerPagination(1, 7, 0, 0)
	if result != "1 ..." {
		t.Errorf("The expected value was '1 2 3 4 5 6 7 ... 9 10', but we received %v ", result)
	}

	result = footerPagination(7, 7, 0, 0)
	if result != "... 7" {
		t.Errorf("The expected value was '1 2 3 4 5 6 7 ... 9 10', but we received %v ", result)
	}

	result = footerPagination(3, 5, 0, 0)
	if result != "... 3 ..." {
		t.Errorf("The expected value was '1 2 3 4 5 6 7 ... 9 10', but we received %v ", result)
	}

	result = footerPagination(5, 10, 2, 2)
	if result != "1 2 3 4 5 6 7 ... 9 10" {
		t.Errorf("The expected value was '1 2 3 4 5 6 7 ... 9 10', but we received %v ", result)
	}

	result = footerPagination(1, 7, 2, 2)

	if result != "1 2 3 ... 6 7" {
		t.Errorf("The expected value was '1 2 3 4 5 6 7 ... 9 10', but we received %v ", result)
	}

	result = footerPagination(3, 5, 10, 10)

	if result != "1 2 3 4 5" {
		t.Errorf("The expected value was '1 2 3 4 5 6 7 ... 9 10', but we received %v ", result)
	}

	result = footerPagination(10, 10, 2, 2)

	if result != "1 2 ... 8 9 10" {
		t.Errorf("The expected value was '1 2 3 4 5 6 7 ... 9 10', but we received %v ", result)
	}
}

func contain(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
