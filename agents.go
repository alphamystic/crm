package main

import (
  "fmt"
  "errors"
  "database/sql"
  "golang.org/x/crypto/bcrypt"
  _ "github.com/go-sql-driver/mysql"
)

type Agent struct{
  FirstName string
  SecondName string
  Email string
  PhoneNo string
  AgentId string
  Active bool
  CreatedAt string
  UpdatedAt string
}

//transaction with users to create a password for authentication
func CreateAgent(a Agent)error{
  tx,err := db.Begin()
  //defer tx.Rollback()
  if err !=  nil{
     _ = tx.Rollback()
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert agent: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating agent.")
  }
  //create the agent
  var result sql.Result
  result,err = tx.Exec("INSERT INTO `crm`.`agent` (firstname,secondname,email,phoneno,agentid,active,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?);",a.FirstName,a.SecondName,a.Email,a.PhoneNo,a.AgentId,a.Active,a.CreatedAt,a.UpdatedAt)
  rowsAffec,_ := result.RowsAffected()
  if err != nil || rowsAffec != 1  {
    _ = tx.Rollback()
    e := LogErrorToFile("sql",fmt.Sprintf("Error inserting into agent: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating agent.")
  }
  //create a password hash for the temporary password
  var passwordHash []byte
  passwordHash,err = bcrypt.GenerateFromPassword([]byte(a.AgentId),bcrypt.DefaultCost)
  if err != nil{
     _ = tx.Rollback()
    e := LogErrorToFile("bcrypt",fmt.Sprintf("Error creating password hash: %s\n",err))
    logError(e)
    return  errors.New("Server encountered an error while trying to create agent")
  }
  //Create the user (for login purposes)
  result,err = tx.Exec("INSERT INTO `crm`.`users` (email,agentid,password) VALUES(?,?,?);",a.Email,a.AgentId,passwordHash);
  rowsAffec,_ = result.RowsAffected()
  if err != nil || rowsAffec != 1 {
    _ = tx.Rollback()
    e := LogErrorToFile("sql",fmt.Sprintf("Error inserting into users: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating agent.")
  }
  //commit transaction
  if err = tx.Commit(); err != nil{
    _ = tx.Rollback()
    e := LogErrorToFile("sql","Error comiting create agent: ",err)
    logError(e)
    return errors.New("Server encountered an error while trying to create agent please try again later")
  }
  return nil
}

func ViewAgent(aid string)(*Agent,error){
  var a Agent
  row := db.QueryRow("SELECT * FROM `crm`.`agent` WHERE agentid	 = ?;",aid)
  err := row.Scan(&a.FirstName,&a.SecondName,&a.Email,&a.PhoneNo,&a.AgentId,&a.Active,&a.CreatedAt,&a.UpdatedAt)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("ERROR viewing agent with id of %s \nERROR: %s\n",aid,err))
    logError(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing agent with id of %s",aid))
  }
  return &a,nil
}

func ListAgents()([]Agent,error){
  stmt := "SELECT * FROM `crm`.`agent`;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all Agents.")
  }
  defer rows.Close()
  var agents []Agent
  for rows.Next(){
    var a Agent
    err = rows.Scan(&a.FirstName,&a.SecondName,&a.Email,&a.PhoneNo,&a.AgentId,&a.Active,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig agent rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all agents.")
    }
    agents = append(agents,a)
  }
  return agents,nil
}
