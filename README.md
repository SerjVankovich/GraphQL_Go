# GraphQL_Go
Simple graphQL backend on Golang

## DataBase
To use this backend you need PostgreSQL database on your PC with table **"zradlo"** with columns:
* **ID** (integer) primary key autoincrement
* **Name** (character varying(255))
* **Price** (real)

And table **users** with columns:
* **id** (integer) primary key autoincrement
* **email** (character varying(255))
* **password** (character varying(255))
* **confirmed** (boolean)

To connect DB make **config.json** file in **"db"** directory:

    {

      "user": %PGUSER%,
  
      "password": %PGPASSWORD%,
  
      "dbname": %PDDATABASE%,
  
      "sslmode": "disable" | "enable"

    }

## JWT Token
To be able to get **JWT Token** make **keys.json** file in project directory:

    {
        "json-secret": %YOURSECRET%
    }

To test it, run "go build main.go" and run the compiled file, then go to the endpoint: http://localhost:8080/zradlo

To go to each of endpoints you need to get **JWT Token**. To do it register to service

## API
### Registration
To **register** on service go to <br> http://localhost:8080/register?query=mutation+_{register(email:"%YOUREMAIL%",password:"%YOURPASSWORD%"){regLink}} <br>
A letter will be sent on your Email with confirmation link. Confirmation link will return **token**. Save it.

With each request provide your Token in Header:
    
    "Authorization": "bearer %YOURTOKEN%"
    
If you have already registered on service you can login.

### Login
To **login** go to <br> http://localhost:8080/login?query=mutation+_{emPass(email:"%YOUREMAIL%",password:"%YOURPASSWORD%"){token}} <br>
It returns **JWT Token**. Save it.

### Get all
To **get all** items go to <br> http://localhost:8080/zradlo?query={zradla{name,id,price}} <br>
Don't forget to provide **JWT Token**

### Get one
To **get one** item by ID go to <br> http://localhost:8080/zradlo?query={zradlo(id:%YOURID%){name,id,price}} <br>
Don't forget to provide **JWT Token**

### Add one
To **add** new item go to <br> http://localhost:8080/zradlo?query=mutation+_{create(name:"%YOURNAME%,price:%YOURPRICE%"){id,name,price}}
<br>
Don't forget to provide **JWT Token**

### Update one
To **update** one of items go to <br> http://localhost:8080/zradlo?query=mutation+_{update(id:%YOURID%,name:"%YOURNAME%",price:%YOURPRICE%){id,name,price}} <br>
Don't forget to provide **JWT Token**

### Delete one
To **delete** one of items by ID go to <br> http://localhost:8080/zradlo?query=mutation+_{delete(id:%YOURID%){id,name,price}} <br>
Don't forget to provide **JWT Token**

### Delete many
To **delete multiple items** go to <br> http://localhost:8080/zradlo?query=mutation+_{deleteMore(ids:"%YOURID1%,%YOURID2%"){id,name,price}}. <br> Note that you mustn't write spaces between ids. Use comma separator.
Don't forget to provide **JWT Token**
