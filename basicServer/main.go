package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)


type user struct{
	ID string `json:"id"`
	UserName  string `json:"userName"`
	Password string `json:"password"`
}

var Users =[]user {
	{ID:"1",UserName:"hero",Password:"heroine"},
}

// send all users in json format
func getUser(c *gin.Context){
	c.IndentedJSON(http.StatusOK,Users)
}

func addUser(c *gin.Context){
	var newUser user
	c.BindJSON(&newUser)
	Users = append(Users,newUser)
	c.IndentedJSON(http.StatusCreated,newUser)
}

// func getUserById(id string) (*user,error){
// 	for i,u := range Users{
// 		if u.ID == id {
// 			return &Users[i],nil
// 		}
// 	}
// }


func main(){
	router := gin.Default()
	router.GET("/users",getUser)
	router.POST("/user",addUser)
	router.Run("localhost:3000")
}

