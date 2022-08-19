package main

/*
  * Contains data for defining and manipulating contacts
*/

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Contact struct{
  CreatorId string
  ContactId string
  FirstName string
  LastName string
  Email string
  PhoneNo string
  DateOfLastContact string
  TypeOfContact string
  Description string
  Company string
  Notes string
  ProffesionalTitle string
  Address string
  City string
  State string
  ZipCode int
  CreatedAt string
  UpdatedAt string
}

func SeachContact(name,agentId string)([]Contact,error){
  fmt.Println("seracjknnuihn")
  stmt := "SELECT * FROM `contact` WHERE firstname OR lastname LIKE = ? AND creatorid  = ?;"
  rows,err := db.Query(stmt,name,agentId)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while searching for contatc. Please try again later :)")
  }
  defer rows.Close()
  var contacts []Contact
  for rows.Next(){
    var c Contact
    err = rows.Scan(&c.CreatorId,&c.ContactId,&c.FirstName,&c.LastName,&c.Email,&c.PhoneNo,&c.DateOfLastContact,&c.TypeOfContact,&c.Description,&c.Company,&c.Notes,&c.ProffesionalTitle,&c.Address,&c.City,&c.State,&c.ZipCode,&c.CreatedAt,&c.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig conntact rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while searching for contact.")
    }
    contacts = append(contacts,c)
  }
  return contacts,nil
}

func CreateContact(cnt Contact)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `crm`.`contact` (creatorid,contactid,firstname,lastname,email,phoneno,dateoflastcontact,typeofcontact,description,company,notes,professionaltitle, address,city,state,zipcode,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert contact: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating contact, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(cnt.CreatorId,cnt.ContactId,cnt.FirstName,cnt.LastName,cnt.Email,cnt.PhoneNo,cnt.DateOfLastContact,cnt.TypeOfContact,cnt.Description,cnt.Company,cnt.Notes,cnt.ProffesionalTitle,cnt.Address,cnt.City,cnt.State,cnt.ZipCode,cnt.CreatedAt,cnt.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error on rows affected while creating contact: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating contact. Please try again later :)")
  }
  return nil
}

func UpdateContact(cnt Contact)error{
  upStmt := "UPDATE `crm`.`contact` SET  `email` = ?, `PhoneNo` = ?, `dateoflastcontact` = ?, `typeofcontact` = ?, `description` = ?, `company` = ?, `notes` = ?, `professionaltitle` = ?, `address` = ?, `city` = ?, `state` = ?, `zipcode` = ?, `updated_at` = ? WHERE (`creatorid` = ? AND `contactid` = ?);"
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := LogErrorToFile("sql","Error preparing update contact: ",err)
    logError(e)
    return errors.New("Server encountered an error while updating contact. Please try again later :)")
  }
  defer stmt.Close()
	var result sql.Result
  result,err = stmt.Exec(cnt.Email,cnt.PhoneNo,cnt.DateOfLastContact,cnt.TypeOfContact,cnt.Description,cnt.Company,cnt.Notes,cnt.ProffesionalTitle,cnt.Address,cnt.City,cnt.State,cnt.ZipCode,cnt.UpdatedAt,cnt.CreatorId,cnt.ContactId)
  rowsAffec, _ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql","Error updating contact: ",err)
    logError(e)
    return errors.New("Server encountered an error while trying to update contact. Please try again later :)")
  }
  return nil
}

func ListMyContacts(agntId string)([]Contact,error){
  stmt := "SELECT * FROM `crm`.`contact` WHERE (`creatorid` = ?);"
  rows,err := db.Query(stmt,agntId)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all agents contatcs. Please try again later :)")
  }
  defer rows.Close()
  var contacts []Contact
  for rows.Next(){
    var c Contact
    err = rows.Scan(&c.CreatorId,&c.ContactId,&c.FirstName,&c.LastName,&c.Email,&c.PhoneNo,&c.DateOfLastContact,&c.TypeOfContact,&c.Description,&c.Company,&c.Notes,&c.ProffesionalTitle,&c.Address,&c.City,&c.State,&c.ZipCode,&c.CreatedAt,&c.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig conntact rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all agents contacts.")
    }
    contacts = append(contacts,c)
  }
  return contacts,nil
}

func ListAllContacts()([]Contact,error){
  stmt := "SELECT * FROM `crm`.`contact`;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all contatcs. Please try again later :)")
  }
  defer rows.Close()
  var contacts []Contact
  for rows.Next(){
    var c Contact
    err = rows.Scan(&c.CreatorId,&c.ContactId,&c.FirstName,&c.LastName,&c.Email,&c.PhoneNo,&c.DateOfLastContact,&c.TypeOfContact,&c.Description,&c.Company,&c.Notes,&c.ProffesionalTitle,&c.Address,&c.City,&c.State,&c.ZipCode,&c.CreatedAt,&c.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig conntact rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all contacts.")
    }
    contacts = append(contacts,c)
  }
  return contacts,nil
}

func ViewMyContact(cntId,agntId string)(*Contact,error){
  var c Contact
  row := db.QueryRow("SELECT * FROM `crm`.`contact` WHERE (`creatorid` = ? AND `contactid` = ?);",agntId,cntId)
  err := row.Scan(&c.CreatorId,&c.ContactId,&c.FirstName,&c.LastName,&c.Email,&c.PhoneNo,&c.DateOfLastContact,&c.TypeOfContact,&c.Description,&c.Company,&c.Notes,&c.ProffesionalTitle,&c.Address,&c.City,&c.State,&c.ZipCode,&c.CreatedAt,&c.UpdatedAt)
  if err != nil{
    fmt.Println("")
    fmt.Println("")
    fmt.Println(err)
    fmt.Println("")
    fmt.Println("")
    e := LogErrorToFile("sql",fmt.Sprintf("ERROR viewing contact with id of %s :ERROR: \n",cntId,err))
    logError(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing contact with id of %s and agent id: %s",cntId,agntId))
  }
  return &c,nil
}

func DeleteContact(cntId,agntId string)error{
  del,err := db.Prepare("DELETE FROM `crm`.`contact` WHERE (`creatorid` = ? AND `contactid` = ?);")
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error deleting contact with id %s of agent id %s, ERROR: %s \n",cntId,agntId,err))
    logError(e)
    return errors.New("Server encountered an error while deleting contact.")
  }
  defer del.Close()
  var result sql.Result
  result,err = del.Exec(agntId,cntId)
  rowsAffec, _ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql","Error deleting contact: ",err)
    logError(e)
    return errors.New("Server encountered an error while deleting contact")
  }
  return nil
}
