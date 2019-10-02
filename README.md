## Prerequisites
[GoLang](https://golang.org/dl/) <br>

## Development setup
### Project setup
Set the go path in ~/.bash_profile
```
vi ~/.bash_profile
```

```
export GOPATH="<desired path>"
eg: export GOPATH="/Users/me/go"
```
Create the following folder structure from the go path
```
cd <go path>
eg:cd /Users/me/go
```

### Build and Run Locally
Go to the git repo and run dep to fetch vendor libraries
```
dep ensure
```
Build project
```
go build
```
Run project
```
./ecosia-application
```


###Developer Notes


###Swagger

API documentation

```$xslt
{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This documentation describes example APIs found under https://github.com/sunilnk17/Ecosia",
    "title": "Favorite tree",
    "contact": {
      "name": "Sunil NK",
      "email": "sunilnk95@gmail.com"
    },
    "version": "1.0.0"
  },
  "host": "127.0.0.1",
  "paths": {
    "/": {
      "get": {
        "description": "Get Favorite tree",
        "schemes": [
          "http"
        ],
        "summary": "Get Favorite tree from the URL",
        "parameters": [
          {
            "type": "string",
            "description": "Favorite tree name",
            "name": "favoriteTree",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The request is successful."
          },
          "400": {
            "description": "The request doesn't meet the requirements"
          },
          "404": {
            "description": "The request is invalid"
          }
        }
      }
    }
  }
}
```

####GET
```$xslt
http://127.0.0.1:8000/?favoriteTree=apple
```

The Application returns 200 for the successful request. 400 for Bad request. 404 invalid request.

######TestCase 1: 

```$xslt
http://localhost:8000/?favoriteTree=apple
```

Result : This is a valid request and returns 200 with HTML encoding text in the browser.

######TestCase 2: 

```$xslt
http://localhost:8000/?
```

Result : This is a Bad request and returns 400 with corresponding response.

######TestCase 3: 

```$xslt
http://localhost:8000/?favoriteTee=apple
```

Result : This is an invalid request returns 400 with corresponding response.

######TestCase 4: 

```$xslt
http://localhost:8000?fa=tree
```

Result : This is an invalid request returns 404.


###To view the test cases

```$xslt
ecosia-application/ecosia_test.go
```




