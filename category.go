package main

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Category struct{
  Name string
  CategoryId string
  Description string
  CreatedAt string
  UpdatedAt string
}

func EditCategory(c Category)error{
  upStmt := "UPDATE `crm`.`category` SET `name` = ?, `description` = ?, `updated_at` = ? WHERE (`categoryid` = ?);"
  stmt, err := db.Prepare(upStmt)
  res, err := stmt.Exec(c.Name,c.Description,currentTime)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("EEC with id %s %s",c.CategoryId,err))
    logError(e)
    return errors.New("Server encountered an error while editing category.")
  }
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
    e := LogErrorToFile("sql",err)
    logError(e)
		return err
	}
  return nil
}

func CreateCategory(c Category)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `crm`.`category` (name,categoryid,description,created_at,updated_at) VALUES(?,?,?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert category: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating category, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(c.Name,c.CategoryId,c.Description,c.CreatedAt,c.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error on rows affected while creating category: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating category.")
  }
  return nil
}

func ListCategories()([]Category,error){
  stmt := "SELECT * FROM `crm`.`categories`;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all categories.")
  }
  defer rows.Close()
  var categories []Category
  for rows.Next(){
    var c Category
    err = rows.Scan(&c.Name,&c.CategoryId,&c.Description,&c.CreatedAt,&c.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig category rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all categories.")
    }
    categories = append(categories,c)
  }
  return nil,nil
}

func ViewCategory(ctgId string)(*Category,error){
  var c Category
  row := db.QueryRow("SELECT * FROM `crm`.`category` WHERE categoryid	 = ?;",ctgId)
  err := row.Scan(&c.Name,&c.CategoryId,&c.Description,&c.CreatedAt,&c.UpdatedAt)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("ESC id: %s, %s",ctgId,err))
    logError(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing agent with id of %s",ctgId))
  }
  return &c,nil
}

func DeleteCategory(ctgId string)error{
  err := CheckIfCategoryHasProduct(ctgId)
  if err != nil{
    return err
  }
  del, err := db.Prepare("DELETE FROM `crm`.`category` WHERE (`idproducts` = ?);")
	if err != nil {
		e := LogErrorToFile("sql",fmt.Sprintf("EDC with ID %s %s",ctgId,err))
    logError(e)
    return errors.New("Server encountered an error while deleting category")
	}
  defer del.Close()
  var res sql.Result
  res,err = del.Exec(ctgId)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("EDC with ID %s %s",ctgId,err))
    logError(e)
    return errors.New("Server encountered an error while deleting category")
  }
  return nil
}
