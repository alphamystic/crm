package main

/*
  * Contains data definers and manipulators for deals
*/

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Deal struct{
  ContactId string
  AgentId string
  DealId string
  Description string
  Complete bool
  Approoved bool
  CreatedAt string
  UpdatedAt string
}

func CreateDeal(d Deal)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `crm`.`deals` (contactid,agentid,dealid,description,complete,approoved,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert deal: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating deal, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(d.ContactId,d.AgentId,d.DealId,d.Description,d.Complete,d.Approoved,d.CreatedAt,d.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error on rows affected while creating deal: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating deal.")
  }
  return nil
}

func CreateEnquiry(e Enquiry)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `crm`.`enquiry` (name,email,subject,message) VALUES(?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert enquiry: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating enquiry, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(e.Name,e.Email,e.Subject,e.Message)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error on rows affected while creating deal: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating enquiry.")
  }
  return nil
}
func UpdateDeal(dealId,agentId,description string) error{
  upStmt := "UPDATE `crm`.`deals` SET  `description` = ? AND `updated_at` = ? WHERE (`agentid` = ? AND `dealid` = ?);"
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := LogErrorToFile("sql","Error preparing update deal: ",err)
    logError(e)
    return errors.New("Server encountered an error while updating deal. Please try again later :)")
  }
  defer stmt.Close()
	var result sql.Result
  result,err = stmt.Exec(description,currentTime,agentId,dealId)
  rowsAffec, _ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql","Error updating deal: ",err)
    logError(e)
    return errors.New("Server encountered an error while trying to update deal. Please try again later :)")
  }
  return nil
}

func MarkDealDone(dealId,agentId string) error{
  upStmt := "UPDATE `crm`.`deals` SET  `complete` = ?, `approoved` = ? AND `updated_at` = ? WHERE (`agentid` = ? AND `dealid` = ?);"
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := LogErrorToFile("sql","Error preparing deal done deal: ",err)
    logError(e)
    return errors.New("Server encountered an error while making deal done. Please try again later :)")
  }
  defer stmt.Close()
	var result sql.Result
  result,err = stmt.Exec(true,true,currentTime,agentId,dealId)
  rowsAffec, _ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql","Error updating deal: ",err)
    logError(e)
    return errors.New("Server encountered an error while trying to update deal. Please try again later :)")
  }
  return nil
}

func ListAllDealls()([]Deal,error){
  stmt := "SELECT * FROM `crm`.`deals`;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all deals.")
  }
  defer rows.Close()
  var deals []Deal
  for rows.Next(){
    var d Deal
    err = rows.Scan(&d.ContactId,&d.AgentId,&d.DealId,&d.Description,&d.Complete,&d.Approoved,&d.CreatedAt,&d.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig deal rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all deals.")
    }
    deals = append(deals,d)
  }
  return deals,nil
}

func ListAllMyDeals(agntId string)([]Deal,error){
  stmt := "SELECT * FROM `crm`.`deals` WHERE (`agentid` = ?);"
  rows,err := db.Query(stmt,agntId)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all agents deals.")
  }
  defer rows.Close()
  var deals []Deal
  for rows.Next(){
    var d Deal
    err = rows.Scan(&d.ContactId,&d.AgentId,&d.DealId,&d.Description,&d.Complete,&d.Approoved,&d.CreatedAt,&d.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig deal rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all agents deals. Please try again later :)")
    }
    deals = append(deals,d)
  }
  return deals,nil
}

func ViewDeal(agntId,dealId string)(*Deal,error){
  var d Deal
  row := db.QueryRow("SELECT * FROM `crm`.`deals` WHERE agentid = ? AND `dealid` = ?;",agntId,dealId)
  err := row.Scan(&d.ContactId,&d.AgentId,&d.DealId,&d.Description,&d.Complete,&d.Approoved,&d.CreatedAt,&d.UpdatedAt)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("ERROR viewing deal with deal id of %s :ERROR: %s\n",dealId,err))
    logError(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing deal with id of %s",dealId))
  }
  return &d,nil
}

func ListAllCompleteDeals()([]Deal,error){
  stmt := "SELECT * FROM `crm`.`deals` WHERE (`approoved` = ?);"
  rows,err := db.Query(stmt,true)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all complete deals.")
  }
  defer rows.Close()
  var deals []Deal
  for rows.Next(){
    var d Deal
    err = rows.Scan(&d.ContactId,&d.AgentId,&d.DealId,&d.Description,&d.Complete,&d.Approoved,&d.CreatedAt,&d.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig deal rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all complete deals.")
    }
    deals = append(deals,d)
  }
  return deals,nil
}

func ListAllIncompleteDeals()([]Deal,error){
  stmt := "SELECT * FROM `crm`.`deals` WHERE `approoved` = ?;"
  rows,err := db.Query(stmt,true)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all complete deals.")
  }
  defer rows.Close()
  var deals []Deal
  for rows.Next(){
    var d Deal
    err = rows.Scan(&d.ContactId,&d.AgentId,&d.DealId,&d.Description,&d.Complete,&d.CreatedAt,&d.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig deal rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all complete deals.")
    }
    deals = append(deals,d)
  }
  return deals,nil
}

func ListAllMyIncompleteDeals(agntId string)([]Deal,error){
  stmt := "SELECT * FROM `crm`.`deals` WHERE (`approoved` = ? AND agentid = ?)ORDER BY updated_at ASC;"
  rows,err := db.Query(stmt,false,agntId)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all my incomplete deals.")
  }
  defer rows.Close()
  var deals []Deal
  for rows.Next(){
    var d Deal
    err = rows.Scan(&d.ContactId,&d.AgentId,&d.DealId,&d.Description,&d.Complete,&d.Approoved,&d.CreatedAt,&d.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig deal rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all incomplete deals.")
    }
    deals = append(deals,d)
  }
  return deals,nil
}

func ListAllMyCompleteDeals(agntId string)([]Deal,error){
  stmt := "SELECT * FROM `crm`.`deals` WHERE (`approoved` = ? AND `agentid` = ?)  ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,true,agntId)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all agents complete deals.")
  }
  defer rows.Close()
  var deals []Deal
  for rows.Next(){
    var d Deal
    err = rows.Scan(&d.ContactId,&d.AgentId,&d.DealId,&d.Description,&d.Complete,&d.Approoved,&d.CreatedAt,&d.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig deal rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all complete deals.")
    }
    deals = append(deals,d)
  }
  return deals,nil
}
