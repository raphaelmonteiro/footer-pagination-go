package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	currentPage, totalPages, boundaries, around, error := checkVariables()

	if len(error) > 0 {
		fmt.Printf("The params %v is invalid\n", strings.Join(error, ", "))
		os.Exit(1)
	}

	pagination := footerPagination(currentPage, totalPages, boundaries, around)

	fmt.Println(pagination)
}

//checkVariables() Checks if all variables were informed correctly.
func checkVariables() (int, int, int, int, []string) {
	var error []string

	currentPage, err := strconv.Atoi(os.Getenv("current_page"))
	if err != nil {
		error = append(error, "current_page")
	}
	totalPages, err := strconv.Atoi(os.Getenv("total_pages"))
	if err != nil {
		error = append(error, "total_pages")
	}
	boundaries, err := strconv.Atoi(os.Getenv("boundaries"))
	if err != nil {
		error = append(error, "boundaries")
	}
	around, err := strconv.Atoi(os.Getenv("around"))
	if err != nil {
		error = append(error, "around")
	}
	if currentPage > totalPages {
		error = append(error, "current_page")
	}

	return currentPage, totalPages, boundaries, around, error
}

//footerPagination() works the received information to return the footer pagination string.
func footerPagination(currentPage int, totalPages int, boundaries int, around int) string {
	var pagination string

	for i := 1; i <= totalPages; i++ {
		if i <= boundaries || i > totalPages-boundaries {
			pagination += strconv.Itoa(i) + " "
			continue
		}

		if i >= currentPage-around && i < currentPage || currentPage+around >= i && i >= currentPage {
			pagination += strconv.Itoa(i) + " "
		} else if !strings.HasSuffix(pagination, "... ") {
			pagination += "... "
		}
	}

	return strings.TrimSpace(pagination)
}
