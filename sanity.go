package main

/*
  * This are various functions for doing sanity checks to maintain system and dataintegrity
  * Some functions sanitize data to be more human readable rather presentable
*/

import (
  "fmt"
  "errors"
  //"database/sql"
//  _ "github.com/go-sql-driver/mysql"
)

type Dashboard struct{
  PendingAppointments int64
  PendingTodos int64
  ApproovedDeals int64
  PendingDeals int64
}
//struct defines a client who contains different kinds of data
type Client struct{
  FirstName string
  LastName string
  Email string
  PhoneNo string
  ClientId string//this is the contactid
  //DataId string//will be initialized on html to match its representation
  Data string//range from an appointment description or a deal deal description
  CreatedAt string
  UpdatedAt string
}

func ViewAgentDeals(agentId string)([]Client,error){
  var clients []Client
  deals,err := ListAllMyDeals(agentId)
  if err != nil{
    e := LogErrorToFile("sanity","Error listing agent deals: ",err)
    logError(e)
    return nil,errors.New("Server encountered an error while viewing client deals")
  }
  var c Client
  //for each deal view the contact and append the data
  for _,deal := range deals{
    //put it into client data
    fmt.Println("first contatct id: ",deal.ContactId)
    cd,err := ViewMyContact(deal.ContactId,agentId)
    fmt.Println("second contatct id: ",deal.ContactId)
    if err != nil{
      e := LogErrorToFile("sanity","Server encountered an error while viewing contact. ",err)
      logError(e)
      return nil,errors.New("Server encountered an error while viewing contact.")
    }
    c = Client{
      FirstName:cd.FirstName,
      LastName:cd.LastName,
      Email:cd.Email,
      PhoneNo:cd.PhoneNo,
      ClientId:cd.ContactId,
      Data:deal.Description,
      CreatedAt:deal.CreatedAt,
      UpdatedAt: deal.UpdatedAt,
    }
    fmt.Println(clients)
    clients = append(clients,c)
  }
  return clients,nil
}

func DoesContactIdExist(id string)bool{
  return true
}
//SELECT COUNT(*) AS num  FROm `crm`.`appointments` WHERE agentid = 'LQYPVW5QRV';
func IsManager(id string)bool{
  return true
}

func GetDashboardData(agentId string)(Dashboard,error){
  var d Dashboard
  pa,err := GetPendingAppoiintments(agentId)
  logError(err)
  pt,err := GetPendingTodos(agentId)
  logError(err)
  ad,err := GetApproovedDeals(agentId)
  logError(err)
  pd,err := GetPendingDeals(agentId)
  logError(err)
  d = Dashboard{
    PendingAppointments:pa,
    PendingTodos:pt,
    ApproovedDeals:ad,
    PendingDeals:pd,
  }
  return d,nil
}

func GetPendingAppoiintments(agentId string)(int64,error){
  var pa int64
  stmt := "SELECT COUNT(*) AS num  FROM `crm`.`appointments` WHERE agentid = ? AND `done` = ?;"
  rows,err := db.Query(stmt,agentId,false)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return pa,errors.New("Server encountered an error while counting pending appointments.")
  }
  defer rows.Close()
  for rows.Next(){
    err = rows.Scan(&pa)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig appointment rows: %s\n",err))
      logError(e)
      return pa,errors.New("Server encountered an error while listing all deals.")
    }
  }
  return pa,nil
}

func GetPendingTodos(agentId string)(int64,error){
  var pt int64
  stmt := "SELECT COUNT(*) AS num  FROM `crm`.`todo` WHERE agentid = ? AND `complete` = ?;"
  rows,err := db.Query(stmt,agentId,false)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return pt,errors.New("Server encountered an error while counting pending todos.")
  }
  defer rows.Close()
  for rows.Next(){
    err = rows.Scan(&pt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig todo rows: %s\n",err))
      logError(e)
      return pt,errors.New("Server encountered an error while listing all deals.")
    }
  }
  return pt,nil
}

func GetApproovedDeals(agentId string)(int64,error){
  var ad int64
  stmt := "SELECT COUNT(*) AS num  FROM `crm`.`deals` WHERE agentid = ? AND `approoved` = ?;"
  rows,err := db.Query(stmt,agentId,true)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return ad,errors.New("Server encountered an error while counting all approoved deals.")
  }
  defer rows.Close()
  for rows.Next(){
    err = rows.Scan(&ad)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig deal rows: %s\n",err))
      logError(e)
      return ad,errors.New("Server encountered an error while listing all approoved deals.")
    }
  }
  return ad,nil
}

func GetPendingDeals(agentId string)(int64,error){
  var pd int64
  stmt := "SELECT COUNT(*) AS num  FROM `crm`.`deals` WHERE agentid = ? AND `approoved` = ?;"
  rows,err := db.Query(stmt,agentId,false)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return pd,errors.New("Server encountered an error while counting pending appointments.")
  }
  defer rows.Close()
  for rows.Next(){
    err = rows.Scan(&pd)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig deal rows: %s\n",err))
      logError(e)
      return pd,errors.New("Server encountered an error while listing all deals.")
    }
  }
  return pd,nil
}


func ViewClientDeals(agentId,clientId string)([]Client,error){
  return nil,nil
}
