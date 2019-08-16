# GraphQL_Go
Simple graphQL backend on Golang

To test it, run "go build main.go" and run the compiled file, then go to the endpoint: http://localhost:8080/zradlo

To **get all** items go to http://localhost:8080/zradlo?query={zradla{name,id,price}}

To **get one** item by ID go to http://localhost:8080/zradlo?query={zradlo(id:%YOURID%){name,id,price}}

To **add** new item go to http://localhost:8080/zradlo?query=mutation+_{create(name:"%YOURNAME%,price:%YOURPRICE%"){id,name,price}}

To **update** one of items go to http://localhost:8080/zradlo?query=mutation+_{update(id:%YOURID%,name:"%YOURNAME%",price:%YOURPRICE%){id,name,price}}

To **delete** one of items by ID go to http://localhost:8080/zradlo?query=mutation+_{delete(id:%YOURID%){id,name,price}}
