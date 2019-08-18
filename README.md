# GraphQL_Go
Simple graphQL backend on Golang

To use this backend you need PostgreSQL database on your PC with table **"zradlo"** and columns:
* **ID** (integer)
* **Name** (character varying(255))
* **Price** (real)

To connect DB make **config.json** file in **"db"** directory:

    {

      "user": %PGUSER%,
  
      "password": %PGPASSWORD%,
  
      "dbname": %PDDATABASE%,
  
      "sslmode": "disable" | "enable"

    }
    
To be able to get JWT Token make **keys.json** file in project directory:

    {
        "json-secret": %YOURSECRET%
    }

To test it, run "go build main.go" and run the compiled file, then go to the endpoint: http://localhost:8080/zradlo

To get each of endpoints you need to get JWT Token. To do it go to
http://localhost:8080/get-token and save it.
With each request provide your Token in Header:
    
    "Authorization": "bearer %YOURTOKEN%"

To **get all** items go to http://localhost:8080/zradlo?query={zradla{name,id,price}}
Don't forget to provide **JWT Token**

To **get one** item by ID go to http://localhost:8080/zradlo?query={zradlo(id:%YOURID%){name,id,price}}
Don't forget to provide **JWT Token**

To **add** new item go to http://localhost:8080/zradlo?query=mutation+_{create(name:"%YOURNAME%,price:%YOURPRICE%"){id,name,price}}
Don't forget to provide **JWT Token**

To **update** one of items go to http://localhost:8080/zradlo?query=mutation+_{update(id:%YOURID%,name:"%YOURNAME%",price:%YOURPRICE%){id,name,price}}
Don't forget to provide **JWT Token**

To **delete** one of items by ID go to http://localhost:8080/zradlo?query=mutation+_{delete(id:%YOURID%){id,name,price}}
Don't forget to provide **JWT Token**

To **delete multiple items** go to http://localhost:8080/zradlo?query=mutation+_{deleteMore(ids:"%YOURID1%,%YOURID2%"){id,name,price}}. Note that you mustn't write spaces between ids. Use comma separator.
Don't forget to provide **JWT Token**
