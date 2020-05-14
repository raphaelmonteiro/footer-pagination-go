# Exercise Footer Pagination

This project is for Worten. It was done in golang and its goal is to return a footer pagination.

For the execution of the test it is necessary to inform some parameters. The expected parameters are those defined in the exercise sent by Worten. They are: current_page, total_pages, boundaries, around

### Execute
```
$ git clone https://github.com/raphaelmonteiro/footer-pagination-go.git
$ cd ./footer-pagination-go
$ go test
$ current_page=4 total_pages=10 boundaries=2 around=2 go run main.go
```
