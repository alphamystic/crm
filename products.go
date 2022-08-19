package main

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)


type Product struct{
  Name string
  ProductId string
  CategoryId string
  Number int //the number of products available
  Description string
  Amount float64
  CreatedAt string
  UpdatedAt string
}

func EditProduct(prdId,name,description string,num int,amount float64)error{
  upStmt := "UPDATE `crm`.`products` SET `name` = ?,`number` = ?,`description` = ?,`amount` = ?, `updated_at` = ? WHERE (`productid` = ?);"
  stmt, err := db.Prepare(upStmt)
  res, err := stmt.Exec(name,num,description,amount,currentTime)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("EEP with id %s %s",prdId,err))
    logError(e)
    return errors.New("Server encountered an error while editing product.")
  }
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
    e := LogErrorToFile("sql",err)
    logError(e)
		return err
	}
  return nil
}

func CheckIfCategoryHasProduct(crgId string)error{
  stmt := "SELECT * FROM `crm`.`products` WHERE categoryid = ?;"
  _,err := db.Query(stmt,crgId)
  if err != nil{
    if err == sql.ErrNoRows{
      return nil
    }
    e := LogErrorToFile("sql",err)
    logError(e)
    return errors.New("Server encountered an error while getting all products.")
  }
  return err
}

func CreateProduct(p Product)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `crm`.`products` (name,productid,categoryid,number,description,amount ,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?);")
  if err !=  nil{
    e := LogErrorToFile("sql",fmt.Sprintf("Error preparing to insert product: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating a product, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(p.Name,p.ProductId,p.CategoryId,p.Number,p.Description,p.Amount,p.CreatedAt,p.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("Error on rows affected while creating product: %s\n",err))
    logError(e)
    return errors.New("Server encountered an error while creating product.")
  }
  return nil
}

func ListProductsOfAPerticularCategory(ctgId string)([]Product,error){
  stmt := "SELECT * FROM `crm`.`products` WHERE categoryid = ?;"
  rows,err := db.Query(stmt,ctgId)
  if err != nil{
    e := LogErrorToFile("sql",err)
    logError(e)
    return nil,errors.New("Server encountered an error while listing all products.")
  }
  defer rows.Close()
  var products []Product
  for rows.Next(){
    var p Product
    err = rows.Scan(&p.Name,&p.ProductId,&p.CategoryId,&p.Number,&p.Description,&p.Amount,&p.CreatedAt,&p.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scannig product rows: %s\n",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing requested products.")
    }
    products = append(products,p)
  }
  return nil,nil
}
/*
func SaleProduct(prdId string,number int)error{
  tx,err := db.Begin()
  of err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("EPST %s",err))
    logError(e)
    return errors.New("Server encountered an error saling product. Try again later")
  }
  var
  return nil
}
*/
func ViewProduct(prdId string)(*Product,error){
  var p Product
  row := db.QueryRow("SELECT * FROM `crm`.`products` WHERE productid = ?;",prdId)
  err := row.Scan(&p.Name,&p.ProductId,&p.CategoryId,&p.Number,&p.Description,&p.Amount,&p.CreatedAt,&p.UpdatedAt)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("ESP id: %s, %s",prdId,err))
    logError(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing product with id of %s",prdId))
  }
  return &p,nil
}

func DeleteProduct(prdId string)error{
  del, err := db.Prepare("DELETE FROM `crm`.`products` WHERE (`productid` = ?);")
	if err != nil {
		e := LogErrorToFile("sql",fmt.Sprintf("EDP with ID %s %s",prdId,err))
    logError(e)
    return errors.New("Server encountered an error while deleting product")
	}
  defer del.Close()
  var res sql.Result
  res,err = del.Exec(prdId)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := LogErrorToFile("sql",fmt.Sprintf("EDC with ID %s %s",prdId,err))
    logError(e)
    return errors.New("Server encountered an error while deleting product")
  }
  return nil
}
