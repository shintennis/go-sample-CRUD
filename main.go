package main

import (
	"fmt"
	"log"
	"todo_go/app/controllers"
	"todo_go/app/models"
)

func main() {
  fmt.Println(models.Db)

  controllers.StartMainServer()
  user, _ := models.GetUserByEmail("test@test.com")
  fmt.Println(user)

  session, err := user.CreateSession()
  if err != nil {
    log.Println(err)
  }
  fmt.Println(session)

  valid, _ := session.CheckSession()
  fmt.Println(valid)
}

