package main

type ContactMe struct{
  CntMeId string
  OwnerId string
  Name string
  Email string
  phoneNumber string
  Viewed bool
  CreatedAt string
  UpdatedAt string
}

type NewsLetter struct{
  OwnerId string
  Email string
  Subscribed bool
}

type Feedback struct{
  OwnerId string
  Name string
  Email string
  phoneNumber string
  Feedback string
  Viewed bool
  CreatedAt string
  UpdatedAt string
}


//lists only the none viewed
func ListContacMes(ownerId string)([]ContactMe,error){
  return nil,nil
}

//lists all contact mes
func ListAllMyContactMes(ownerId string)([]ContactMe,error){
  return nil,nil
}

func MarkContactmeAsViewed(ownerid,contactMeId string)(error){
  return nil
}


//lists only the none viewed feedbacks
func ListFeedBacks()([]Feedback,error){
  return nil,nil
}

func MarkFeedBackAsViewed(feedbackId,ownerId string)error{
  return nil
}
