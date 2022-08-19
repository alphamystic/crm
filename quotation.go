package main

type Quotation struct{
  QuotationId string
  OwnerId string
  ProductId string
  Name string
  Email string
  phoneNumber string
  QuotReq string
  Viewed bool
  CreatedAt string
  UpdatedAt string
}

func ListAllMyQuotations(ownerId string)([]Quotation,error){
  return nil,nil
}

func ListNotViewedQuotations(ownerId string)([]Quotation,error){
  return nil,nil
}

func MarkQuotationAsViewed(quoteId,ownerId string)error{
  return nil
}
