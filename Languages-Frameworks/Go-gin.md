Go Gin Notes

1. Create a go.mod file (to initialise a new go module which is a collection of related go packages) go mod init [module-name]
2. Compile Daemon package for invoking go build if .go file changes go get github.com/githubnemo/CompileDaemon go install github.com/githubnemo/CompileDaemon
3. Godotenv for environment variables go get github.com/joho/godotenv
4. Download and install gin package go get -u github.com/gin-gonic/gin
5. Installing GORM (ORM library for go) [Optional] go get -u gorm.io/gorm go get -u gorm.io/gorm/driver/postgres
6. Create main.go file package main import “fmt” func main() {      fmt.Println(“Hello”) }
7. Run main.go via using CompileDaemon /[Path to CompileDaemon] -command=“./[module-name]” Run “go env GOPATH” and add “/bin” after it /Users/rahul436.kumar/go/bin/CompileDaemon -command=“./[module-name]”
8. Update the main.go file with this 
    router := gin.Default() 
    router.GET("/greet", func(cxt *gin.Context) {       
        cxt.JSON(200, gin.H{"message": "Hello, how are you?”} 
    }) 
    router.Run()
9. Create a function init() like this and define PORT = 3000 in .env file func init() {
    	  err := godotenv.Load()
	  if err != nil {
             log.Fatal("Error loading .env file")
          }
      }        The app should now run on Port 3000. init() function runs before main() function in go. Gin 	framework checks for PORT variable before running the app.
10. Create a folder initialisers and create a file loadEnvVariables.go with the content package initializers  function LoadEnvVariables() { // contents of init() function should be here and change the init() function content to  // initializers.LoadEnvVariables() }
11. For adding MongoDB, go get go.mongodb.org/mongo-driver/mongo     GO/GIN Important Points

1. In Go, struct fields must be exported (i.e., start with a capital letter) in order to be accessible outside of the package.
2. 
3. Interface is a way to group together types based on the functions that they have associated with those types  


