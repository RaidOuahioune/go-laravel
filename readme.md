# My Collection of the awsome go packages 

1. Gin  (the api router)
2. Gorm (the orm)
3. github.com/go-playground/validator  (the  form validator)
4. https://github.com/golang-jwt/jwt  &  github.com/appleboy/gin-jwt/v2 (the authenticator)
5. Look for a cron job library(async)
6. Air for hot-reloading the server
7. Graph QL integration(github.com/99designs/gqlgen)

7. DO translation and pagination(pagination doen with meta data using scopes)
8. DO termination middelware(done some research ; i've found nothing)



# for the graphql (follow steps on github.com/99designs/gqlgen)

# go generate the needed go files afer schema modification : 

```sh
go run github.com/99designs/gqlgen generate
```


## TODOS

# u need to solve the problem with the eager loading 
# adding authentication 
# Swag for the docs 
1. to visit the docs check swagger/index.html 
2. to generate the docs after each new commoents to do type 
```sh
swag init --parseDependency --parseInternal
```
3. for any documnetation u need to comment just above the funcion binded to that route and not in the actual declaration of that route



