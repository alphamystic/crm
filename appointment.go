package main

/*
  * Contains data dfiners and manipulators for appointments
*/
import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type  Appointment struct{
  AgentId string
  AppId string
  Title string
  Description string
  Done bool
  CreatedAt string
  UpdatedAt string
}

func CreateAppointment(a Appointment)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `crm`.`appointments` (agentid,appointmentid,title,description,done,created_at,updated_at) VALUES(?,?,?,?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert appointment: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating appointment, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(a.AgentId,a.AppId,a.Title,a.Description,a.Done,a.CreatedAt,a.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error on rows affected while creating appointment: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating appointment.")
  }
  return nil
}

//mark appointment as complete
func MarkAppointmentAsComplete(agentId,appId string)error{
  upStmt := "UPDATE `crm`.`appointments` SET `done` = ? AND `updated_at` = ? WHERE (`agentid` = ? AND `appointmentid	` = ?);"
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := LogErrorToFile("sql","Error making appointment as done: ",err)
    logError(e)
    return errors.New("Server encountered an error while amrking appointment as done. Please try again later :)")
  }
  defer stmt.Close()
	var result sql.Result
  result,err = stmt.Exec(true,currentTime,agentId,appId)
  rowsAffec, _ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql","Error making apppointment as done: ",err)
    logError(e)
    return errors.New("Server encountered an error while marking appointment as done. Please try again later :)")
  }
  return nil
}

//update appointment
func UpdateAppointment(appId,agentId,title,description string) error{
  upStmt := "UPDATE `crm`.`appointments` SET `title` = ?, `description` = ? AND `updated_at` = ? WHERE (`agentid` = ? AND `appointmentid	` = ?);"
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := LogErrorToFile("sql","Error marking appointment as done: ",err)
    logError(e)
    return errors.New("Server encountered an error while amrking appointment as done. Please try again later :)")
  }
  defer stmt.Close()
	var result sql.Result
  result,err = stmt.Exec(title,description,currentTime,agentId,appId)
  rowsAffec, _ := result.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql","Error making apppointment as done: ",err)
    logError(e)
    return errors.New("Server encountered an error while marking appointment as done. Please try again later :)")
  }
  return nil
}

func ListAllAppointments()([]Appointment,error){
    stmt := "SELECT * FROM `crm`.`appointments`;"
    rows,err := db.Query(stmt)
    if err != nil{
      e := LogErrorToFile("sql",err)
      logError(e)
      return nil,errors.New("Server encountered an error while listing all appointments.")
    }
    defer rows.Close()
    var apps []Appointment
    for rows.Next(){
      var a Appointment
      err = rows.Scan(&a.AgentId,&a.AppId,&a.Title,&a.Description,&a.Done,&a.CreatedAt,&a.UpdatedAt)
      if err != nil{
        e := LogErrorToFile("sql",fmt.Sprintf("Error scannig appointment rows: %s\n",err))
        logError(e)
        return nil,errors.New("Server encountered an error while listing all appointments.")
      }
      apps = append(apps,a)
    }
    return apps,nil
}

func ListAllAgentsAppointments(agentId string)([]Appointment,error){
    stmt := "SELECT * FROM `crm`.`appointments` WHERE agentid = ? ORDER BY updated_at DESC;"
    rows,err := db.Query(stmt,agentId)
    if err != nil{
      e := LogErrorToFile("sql",err)
      logError(e)
      return nil,errors.New("Server encountered an error while listing all agents appointments.")
    }
    defer rows.Close()
    var apps []Appointment
    for rows.Next(){
      var a Appointment
      err = rows.Scan(&a.AgentId,&a.AppId,&a.Title,&a.Description,&a.Done,&a.CreatedAt,&a.UpdatedAt)
      if err != nil{
        e := LogErrorToFile("sql",fmt.Sprintf("Error scannig appointment rows: %s\n",err))
        logError(e)
        return nil,errors.New("Server encountered an error while listing all pending appointments.")
      }
      apps = append(apps,a)
    }
    return apps,nil
}

func ListAllPendingAppointments()([]Appointment,error){
    stmt := "SELECT * FROM `crm`.`appointments` WHERE done = ?;"
    rows,err := db.Query(stmt,false)
    if err != nil{
      e := LogErrorToFile("sql",err)
      logError(e)
      return nil,errors.New("Server encountered an error while listing all pending appointments.")
    }
    defer rows.Close()
    var apps []Appointment
    for rows.Next(){
      var a Appointment
      err = rows.Scan(&a.AgentId,&a.AppId,&a.Title,&a.Description,&a.Done,&a.CreatedAt,&a.UpdatedAt)
      if err != nil{
        e := LogErrorToFile("sql",fmt.Sprintf("Error scannig appointment rows: %s\n",err))
        logError(e)
        return nil,errors.New("Server encountered an error while listing all pending appointments.")
      }
      apps = append(apps,a)
    }
    return apps,nil
}

func ListAllPendingAppointmentsForAPerticularAgent(agntId string)([]Appointment,error){
  stmt := "SELECT * FROM `crm`.`appointments` WHERE agentid = ? AND done = ?;"
  rows,err := db.Query(stmt,agntId,false)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all pending appointments for a perticular user.")
  }
  defer rows.Close()
  var apps []Appointment
  for rows.Next(){
    var a Appointment
    err = rows.Scan(&a.AgentId,&a.AppId,&a.Title,&a.Description,&a.Done,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig appointment rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all pending appointments for a perticular user.")
    }
    apps = append(apps,a)
  }
  return apps,nil
}
