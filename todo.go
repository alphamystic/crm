package main

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Todo struct {
  AgentId string
  TodoId string
  Description string
  Complete bool
  CreatedAt string
  UpdatedAt string
}

func CreateTodo(t Todo)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `crm`.`todo` (agentid,todoid,description,complete,created_at,updated_at) VALUES(?,?,?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert todo: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating appointment, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(t.AgentId,t.TodoId,t.Description,t.Complete,t.CreatedAt,t.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error on rows affected while creating todo: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating todo.")
  }
  return nil
}

//mark todo as complete
func MarkTodoAsComplete(todoId,agentId string)error{
  upStmt := "UPDATE `crm`.`todo` SET `complete` = ? AND `updated_at` = ? WHERE (`agentid` = ? AND `todoid` = ?);"
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := LogErrorToFile("sql","Error marking todo as complete: ",err)
    logError(e)
    return errors.New("Server encountered an error while making todo complete. Please try again later :)")
  }
  defer stmt.Close()
	var result sql.Result
  result,err = stmt.Exec(true,currentTime,agentId,todoId)
  fmt.Println("Executing the stuff")
  fmt.Println("Executing  %s %s %s",currentTime,agentId,todoId)
  fmt.Println("Executed the stuff")
  rowsAffec, _ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    fmt.Sprintf("error two %s rows affected %d",err,rowsAffec)
    e := LogErrorToFile("sql",fmt.Sprintf("Error marking oyyfthdjtdy %s todo complete: %s",todoId,err))
    fmt.Println(err)
    logError(e)
    return errors.New("Server encountered an error while marking todo as complete. Please try again later :)")
  }
  return nil
}

//Update todo description
func UpdateTodoDescription(todoId,agentId,description string)error{
  upStmt := "UPDATE `crm`.`todo` SET `description` = ? WHERE (`agentid` = ? AND `todoid` = ?);"
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := LogErrorToFile("sql","Error updating todo as description: ",err)
    logError(e)
    return errors.New("Server encountered an error while updating todo description. Please try again later :)")
  }
  defer stmt.Close()
	var result sql.Result
  result,err = stmt.Exec(description,agentId,todoId)
  rowsAffec, _ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql","Error updating todo: ",err)
    logError(e)
    return errors.New("Server encountered an error while updating todo. Please try again later :)")
  }
  return nil
}

func ViewIncompleteTodos(agentId string)([]Todo,error){
  stmt := "SELECT * FROM `crm`.`todo` WHERE agentid = ? AND complete = ?;"
  rows,err := db.Query(stmt,agentId,false)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all todos.")
  }
  defer rows.Close()
  var todos []Todo
  for rows.Next(){
    var t Todo
    err = rows.Scan(&t.AgentId,&t.TodoId,&t.Description,&t.Complete,&t.CreatedAt,&t.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig todo rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all incomplete todos.")
    }
    todos = append(todos,t)
  }
  return todos,nil
}

//No one should view someone elses todo
func ListAllTodos(agentId string)([]Todo,error){
  stmt := "SELECT * FROM `crm`.`todo` WHERE agentid = ?;"
  rows,err := db.Query(stmt,agentId)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all appointments.")
  }
  defer rows.Close()
  var todos []Todo
  for rows.Next(){
    var t Todo
    err = rows.Scan(&t.AgentId,&t.TodoId,&t.Description,&t.Complete,&t.CreatedAt,&t.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig todo rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all todos.")
    }
    todos = append(todos,t)
  }
  return todos,nil
}
