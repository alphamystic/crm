package main

import (
  "fmt"
  "log"
  "strconv"
  "net/http"
  "database/sql"
  "html/template"
  "github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
  _ "github.com/go-sql-driver/mysql"
)



var store = sessions.NewCookieStore([]byte("SAM CRM "))

func main(){
  tpl,_ = template.ParseGlob("templates/*.html")
  var err error
  db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/crm")
  if err != nil {
    fmt.Println("[-] Huge error bud, failed to connect to the crm db")
		panic(err.Error())
	}
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/static/",http.StripPrefix("/static",fs))
  http.HandleFunc("/test",BlankTest)
  http.HandleFunc("/",Home)
  http.HandleFunc("/login",Login)
  http.HandleFunc("/logout",Logout)
  http.HandleFunc("searchcontact",Searchforcontact)
  http.HandleFunc("/createagent",Createagent)
  http.HandleFunc("/deletecontact/", Deletecontact)
  http.HandleFunc("/listcontacts",ListAgentsContacts)
  http.HandleFunc("/createcontact",Createcontact)
  http.HandleFunc("/updatecontact/",UpdateContactHandle)
  http.HandleFunc("/updtcontres/",UpdateContactResult)
  http.HandleFunc("/createdeals",Createdeal)
  http.HandleFunc("/listallmydeals",Listallmydeals)
  http.HandleFunc("/listmycompletedeals",Listmycompdeals)
  http.HandleFunc("/listincompletedeals",Listallmyincompletedeals)
  http.HandleFunc("/createtodolist",Createtodolist)
  http.HandleFunc("/incompletetodolist",Listincompletetodolist)
  http.HandleFunc("/todolist",ListTodos)
  http.HandleFunc("/markasdone/",Marktodasdone)
  http.HandleFunc("/createappointment",Createappoitment)
  http.HandleFunc("/listallappointment",Listallappointments)
  http.HandleFunc("/listpendingappointment",Listpendingappointments)
  //Addons
  http.HandleFunc("/enq",Createenquiry)
  fmt.Println("[+] Starting CRM server at 3000!!!")
  err = http.ListenAndServe("0.0.0.0:3000",nil)
  if err != nil {
    log.Fatal("[+] Error starting HTTP server: ",err)
  }
}

//type userData map[string]interface
func BlankTest(res http.ResponseWriter, req *http.Request){
  tpl.ExecuteTemplate(res,"blank.html",nil)
  return
}
func Home(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    d,err := GetDashboardData(agentId)
    if err != nil{
      fmt.Println(err)
    }
    //fmt.Println(" number %d,",pa)
    tpl.ExecuteTemplate(res,"dashboard.html",d)
    return
  }
  http.Redirect(res,req,"/",200)//302
  return
}

func Listallmyincompletedeals(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    deals,err := ListAllMyIncompleteDeals(agentId)
    if err  != nil{
    tpl.ExecuteTemplate(res,"result.html","Server encountered an error while trying to list your incomplete deals, try again later :)")
    return
    }
    if deals == nil || len(deals) <= 0{
      tpl.ExecuteTemplate(res,"result.html"," You don't have any incomplete deals:)")
      return
    }
    tpl.ExecuteTemplate(res,"listincompletedeals.html",deals)
    return
  }
  http.Redirect(res,req,"/listincompletedeals",http.StatusFound)//302
  return
}

type Enquiry struct{
  Name string
  Email string
  Subject string
  Message string
}

func Listenquiry(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  _,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    ///do some stuff here
  }
  http.Redirect(res,req,"/",http.StatusFound)//302
  return
}

func Createenquiry(res http.ResponseWriter, req *http.Request){
  if req.Method != "POST"{
    res.Write([]byte("Create a good request please."))
    return
    }
  req.ParseForm()
  enq := Enquiry{
    Name: req.FormValue("name"),
    Email: req.FormValue("email"),
    Subject: req.FormValue("subject"),
    Message: req.FormValue("message"),
  }
  err := CreateEnquiry(enq)
  if err != nil{
    res.Write([]byte("Inernal server error, please try again later"))
    return
  }
  res.Write([]byte("Successfullty created enquiry...."))
  return
}

func Listmycompdeals(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    deals,err := ListAllMyCompleteDeals(agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error while listing all your appointments. Please try again later :)")
      return
    }
    if deals == nil || len(deals) <= 0{
      tpl.ExecuteTemplate(res,"result.html"," You don't have any complete deals:)")
      return
    }
    tpl.ExecuteTemplate(res,"listmycompletedeals.html",deals)
    return
  }
  http.Redirect(res,req,"/listmycompletedeals",http.StatusFound)//302
  return
}

func Listpendingappointments(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    apps,err := ListAllPendingAppointmentsForAPerticularAgent(agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error while listing all your pending appointments. Please try again later :)")
      return
    }
    if apps == nil || len(apps) <= 0{
      tpl.ExecuteTemplate(res,"result.html","You have no pending appointments")
      return
    }
    tpl.ExecuteTemplate(res,"listpendingappointment.html",apps)
    return
  }
  http.Redirect(res,req,"/listpendingappointment",200)
  return
}

func Listallappointments(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    apps,err := ListAllAgentsAppointments(agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error while listing all your appointments. Please try again later :)")
      return
    }
    tpl.ExecuteTemplate(res,"listallagentsappointments.html",apps)
    return
  }
  http.Redirect(res,req,"/listallappointment",200)
  return
}

func Createappoitment(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "POST"{
    req.ParseForm()
    title := req.FormValue("title")
    description := req.FormValue("description")
    if title == "" || len(title) <= 0 {
      tpl.ExecuteTemplate(res,"createappointment.html","Check on your title.")
      return
    }
    if description == "" || len(description) <= 0 {
      tpl.ExecuteTemplate(res,"createappointment.html","Check on your description.")
      return
    }
    app := Appointment{
      AgentId:agentId,
      AppId:RandString(12),
      Title:title,
      Description:description,
      Done:false,
      CreatedAt:currentTime,
      UpdatedAt:currentTime,
    }
    err := CreateAppointment(app)
    if err != nil{
      tpl.ExecuteTemplate(res,"createappointment.html","Server encountered an error while creating appointment. Please try again later on :).")
      return
    }
    tpl.ExecuteTemplate(res,"createappointment.html","Successfully created appointment.")
    return
  }
  tpl.ExecuteTemplate(res,"createappointment.html",nil)
  return
}

//this is incomplete (errors out a nil erro for whatever the fucking reason)
func Marktodasdone(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  req.ParseForm()
  todoid := req.FormValue("todoid")
  if todoid == "" || len(todoid) <= 0{
    tpl.ExecuteTemplate(res,"result.html","Check on your todoid, refresh page if error persists. LOL :)")
    return
  }
  err := MarkTodoAsComplete(todoid,agentId)
  if err != nil {
    tpl.ExecuteTemplate(res,"result.html","Server encountered an error while marking todo as complete, Please try again later :)")
    return
  }
  tpl.ExecuteTemplate(res,"result.html","Succesfully marked to do as done.")
  return
}

func Listincompletetodolist(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    inctodos,err := ViewIncompleteTodos(agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error while listing all your icomplete todos. Please try again later :)")
      return
    }
    if inctodos == nil || len(inctodos) <= 0{
      tpl.ExecuteTemplate(res,"result.html","You don't have any incomplete todo's.")
      return
    }
    tpl.ExecuteTemplate(res,"linctodo.html",inctodos)
    return
  }
  tpl.ExecuteTemplate(res,"linctodo.html",nil)
  return
}

func ListTodos(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    todos,err := ListAllTodos(agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error while listing your todos, Please try again later :)")
      return
    }
    tpl.ExecuteTemplate(res,"listalltodo.html",todos)
    return
  }
  http.Redirect(res,req,"/",200)//302
  return
}
func Createtodolist(res http.ResponseWriter, req *http.Request){
  if req.Method != "POST"{
    tpl.ExecuteTemplate(res,"createtodo.html",nil)
    return
  }
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  req.ParseForm()
  description := req.FormValue("description")
  if description == "" || len(description) <= 0{
    tpl.ExecuteTemplate(res,"createtodo.html","C'mon man it's just one form, fill it. Literally just put any random data incide it.")
    return
  }
  todo := Todo{
    AgentId:agentId,
    TodoId:RandNoLetter(12),
    Description:description,
    Complete:false,
    CreatedAt:currentTime,
    UpdatedAt:currentTime,
  }
  err := CreateTodo(todo)
  if err != nil{
    tpl.ExecuteTemplate(res,"result.html","Server encountered an error while creating your todo, Please try again later :)")
    return
  }
  tpl.ExecuteTemplate(res,"createtodo.html","Created todo Successfully. Remeber to mark it as complete when done.")
  return
}
/*
func ListincompleteTodos(){}
func DeleteTodo(){}
func UpdatetodoDescription(){}*/
func Listallmydeals(res http.ResponseWriter,req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    deals,err := ViewAgentDeals(agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error. Please try again later :)")
     return
    }
    tpl.ExecuteTemplate(res,"listagentsdeals.html",deals)
    return
  }
  http.Redirect(res,req,"/listallmydeals",200)
  return
}

func Createdeal(res http.ResponseWriter, req *http.Request){
  if req.Method == "POST"{
    session,_ := store.Get(req,"session")
    agentId,ok := session.Values["AgentId"].(string)
    if !ok {
      http.Redirect(res,req,"/login",http.StatusFound)//302
      return
    }
    req.ParseForm()
    contactid := req.FormValue("contactid")
    description := req.FormValue("description")
    if contactid == "" || len(contactid) <= 0{
      tpl.ExecuteTemplate(res,"createdeal.html","Check on your contact Id.")
      return
    }
    if description == "" || len(description) <= 0{
      tpl.ExecuteTemplate(res,"createdeal.html","Check on your description.")
      return
    }
    deal := Deal{
      ContactId:contactid,
      AgentId:agentId,
      DealId:RandString(10),
      Description:description,
      Complete:true,
      Approoved:false,
      CreatedAt:currentTime,
      UpdatedAt:currentTime,
    }
    err := CreateDeal(deal)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error creting your deal, Please try again later :)")
      return
    }
    tpl.ExecuteTemplate(res,"result.html","Deal was created Succefully")
    return
  }
  tpl.ExecuteTemplate(res,"createdeal.html",nil)
  return
}
func UpdateContactHandle(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    contactid := req.FormValue("contactid")
    if contactid == "" || len(contactid) <= 0{
      http.Redirect(res,req,"/listcontacts",307)
      return
    }
    contact,err := ViewMyContact(contactid,agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error while viewing contact. Please try again later.")
      return
    }
    tpl.ExecuteTemplate(res,"updatecontact.html",contact)
    return
  }
  http.Redirect(res,req,"/",307)
  return
}

//email` = ?, `PhoneNo` = ?, `dateoflastcontact` = ?, `typeofcontact` = ?, `description` = ?, `company` = ?, `notes` = ?, `professinaltitle` = ?, `address` = ?, `city` = ?, `state` = ?, `zipcode` = ?, `updated_at
func UpdateContactResult(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "POST"{
    req.ParseForm()
    contactid := req.FormValue("contactid")
    email := req.FormValue("email")
    phonenumber := req.FormValue("phonenumber")
    typeOfContact := req.FormValue("typeOfContact")
    description := req.FormValue("description")
    company := req.FormValue("company")
    notes := req.FormValue("notes")
    proftitle := req.FormValue("proftitle")
    address := req.FormValue("address")
    city := req.FormValue("city")
    state := req.FormValue("state")
    zipcode := req.FormValue("zipcode")
    if email == "" || len(email) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your email.")
      return
    }
    if phonenumber == "" || len(phonenumber) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your phonenumber.")
      return
    }
    if typeOfContact == "" || len(typeOfContact) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your Type of contact.")
      return
    }
    if description == "" || len(description) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your description.")
      return
    }
    if company == "" || len(company) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your company.")
      return
    }
    if notes == "" || len(notes) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your notes.")
      return
    }
    if proftitle == "" || len(proftitle) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your Professional title.")
      return
    }
    if address == "" || len(address) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your address.")
      return
    }
    if city == "" || len(city) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your city.")
      return
    }
    if state == "" || len(state) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on First your state.")
      return
    }
    zc,err := strconv.Atoi(zipcode)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Zip Code must be an integer.")
      return
    }
    contact := Contact{
      CreatorId:agentId,
      ContactId:contactid,
      Email:email,
      PhoneNo:phonenumber,
      DateOfLastContact:currentTime,
      TypeOfContact:typeOfContact,
      Description:description,
      Company:company,
      Notes:notes,
      ProffesionalTitle:proftitle,
      Address:address,
      City:city,
      State:state,
      ZipCode:zc,
      UpdatedAt:currentTime,
    }
    err = UpdateContact(contact)
    if err != nil{
      fmt.Println("[-] ERROR updating contact. ",err)
      http.Redirect(res,req,"/listcontacts",307)
      return
    }
    _ = LogErrorToFile("updates",fmt.Sprintf("Updated contact with ID: %s",contactid))
    tpl.ExecuteTemplate(res,"result.html",fmt.Sprintf("Successfully Updated Contact %s %s.",contact.FirstName, contact.LastName))
    return
  }
  http.Redirect(res,req,"/listcontacts",307)
  return
}

func Createcontact(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method != "POST"{
    tpl.ExecuteTemplate(res,"createcontacts.html",nil)
    return
  }
  req.ParseForm()
  firstname := req.FormValue("firstname")
  lastname := req.FormValue("lastname")
  email := req.FormValue("email")
  phonenumber := req.FormValue("phonenumber")
  typeOfContact := req.FormValue("typeOfContact")
  description := req.FormValue("description")
  company := req.FormValue("company")
  notes := req.FormValue("notes")
  proftitle := req.FormValue("proftitle")
  address := req.FormValue("address")
  city := req.FormValue("city")
  state := req.FormValue("state")
  zipcode := req.FormValue("zipcode")
  if firstname == "" || len(firstname) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your firstname.")
    return
  }
  if lastname == "" || len(lastname) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your lastname.")
    return
  }
  if email == "" || len(email) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your email.")
    return
  }
  if phonenumber == "" || len(phonenumber) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your phonenumber.")
    return
  }
  if typeOfContact == "" || len(typeOfContact) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your Type of contact.")
    return
  }
  if description == "" || len(description) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your description.")
    return
  }
  if company == "" || len(company) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your company.")
    return
  }
  if notes == "" || len(notes) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your notes.")
    return
  }
  if proftitle == "" || len(proftitle) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your Professional title.")
    return
  }
  if address == "" || len(address) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your address.")
    return
  }
  if city == "" || len(city) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your city.")
    return
  }
  if state == "" || len(state) <= 0{
    tpl.ExecuteTemplate(res,"createcontacts.html","Check on First your state.")
    return
  }
  zc,err := strconv.Atoi(zipcode)
  if err != nil{
    tpl.ExecuteTemplate(res,"createcontacts.html","Zip Code must be an integer.")
    return
  }
  contact := Contact{
    CreatorId:agentId,
    ContactId:RandLetters(10),
    FirstName:firstname,
    LastName:lastname,
    Email:email,
    PhoneNo:phonenumber,
    DateOfLastContact:currentTime,
    TypeOfContact:typeOfContact,
    Description:description,
    Company:company,
    Notes:notes,
    ProffesionalTitle:proftitle,
    Address:address,
    City:city,
    State:state,
    ZipCode:zc,
    CreatedAt:currentTime,
    UpdatedAt:currentTime,
  }
  err = CreateContact(contact)
  if err != nil{
    fmt.Println("[-] Error creating contact: ",err)
    tpl.ExecuteTemplate(res,"createcontacts.html","Server encountered an error while creating contact. Please try again later  :)")
    return
  }
  //_ = fmt.Sprintf("updates","Created new contact with the name: %s",contact.ContactId)
  tpl.ExecuteTemplate(res,"createcontacts.html","Contact created Succefully.")
  return
}

func Deletecontact(res http.ResponseWriter,req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    cntId := req.FormValue("contactid")
    err := DeleteContact(cntId,agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error deleting contact. Please try again later on.")
      return
    }
    tpl.ExecuteTemplate(res,"result.html","Successfully deleted contact")
    return
  }
  http.Redirect(res,req,"/",200)
  return
}
func ListAgentsContacts(res http.ResponseWriter,req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "GET"{
    contacts,err := ListMyContacts(agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Serever encountered an error listing your contacts. Please try again later")
      return
    }
    if contacts == nil || len(contacts) <= 0 {
      tpl.ExecuteTemplate(res,"result.html","You have zero contacts, create some.")
      return
      }
    tpl.ExecuteTemplate(res,"listagentscontacts.html",contacts)
    return
  }
  tpl.ExecuteTemplate(res,"/",nil)
  return
}

func Searchforcontact(res http.ResponseWriter,req *http.Request){
  session,_ := store.Get(req,"session")
  agentId,ok := session.Values["AgentId"].(string)
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "POST"{
    req.ParseForm()
    name := req.FormValue("contactname")
    if name == "" || len(name) <= 0{
      tpl.ExecuteTemplate(res,"result.html","Check on the name ypu are searching for.")
      return
    }
    contacts,err := SeachContact(name,agentId)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Serever encountered an error listing your contacts. Please try again later")
      return
    }
    if contacts == nil || len(contacts) <= 0 {
      tpl.ExecuteTemplate(res,"result.html","You have zero contacts, with search possible names, create some.")
      return
      }
    tpl.ExecuteTemplate(res,"listpossible.html",contacts)
    return
  }
  http.Redirect(res,req,"/logout",http.StatusFound)//302
  return
}

//Remember to start session
func Login(res http.ResponseWriter, req *http.Request){
  if req.Method == "POST"{
    req.ParseForm()
    password := req.FormValue("password")
    email := req.FormValue("email")
    var userEmail,hash,agentId string
    stmt := "SELECT email, agentid,password FROM `crm`.`users` WHERE email =?;"
    row := db.QueryRow(stmt,email)
    err := row.Scan(&userEmail,&agentId,&hash)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for auth %s\n",err))
      logError(e)
      tpl.ExecuteTemplate(res,"login.html","Server encountered an error during authentication. Please try again later :)")
      return
    }
    err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err == nil{
      session, _ := store.Get(req,"session")
      session.Values["AgentId"] = agentId
      session.Save(req,res)
      d,err := GetDashboardData(agentId)
      if err != nil{
        fmt.Println(err)
      }
      //fmt.Println(" number %d,",pa)
      tpl.ExecuteTemplate(res,"dashboard.html",d)
      return
    }
    fmt.Println("Error is not equal to nil: ",err)
    e := LogErrorToFile("auth",fmt.Sprintf("wrong login attempt for email %s with password %: ERROR: %s \n",email,password,err))
    logError(e)
    http.Redirect(res,req,"/login",200)
    return
  }
  tpl.ExecuteTemplate(res,"login.html",nil)
  return
}

func Logout(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  delete(session.Values,"AgentId")
  session.Save(req,res)
  tpl.ExecuteTemplate(res,"login.html","Logged Out ADIOS! :)")
  return
}

func Createagent(res http.ResponseWriter,req *http.Request){
  //Ensure user is logged in
  session,_ := store.Get(req,"session")
  _,ok := session.Values["AgentId"]
  if !ok {
    http.Redirect(res,req,"/login",http.StatusFound)//302
    return
  }
  if req.Method == "POST"{
    req.ParseForm()
    firstName := req.FormValue("firstname")
    secondName := req.FormValue("secondname")
    email := req.FormValue("email")
    phoneNumber := req.FormValue("phonenumber")
    if firstName == "" || len(firstName) <= 0{
      tpl.ExecuteTemplate(res,"createagent.html","Check on your firstname")
      return
    }
    if secondName == "" || len(secondName) <= 0{
      tpl.ExecuteTemplate(res,"createagent.html","Check on your second name")
      return
    }
    if email == "" || len(email) <= 0{
      tpl.ExecuteTemplate(res,"createagent.html","Check your email, can't be empty")
      return
    }
    if phoneNumber == "" || len(phoneNumber) <= 0{
      tpl.ExecuteTemplate(res,"createagent.html","Check phone number, can't be empty")
      return
    }
    agent := Agent{
      FirstName:firstName,
      SecondName:secondName,
      Email:email,
      PhoneNo:phoneNumber,
      AgentId:RandNoLetter(10),
      Active:true,
      CreatedAt:currentTime,
      UpdatedAt:currentTime,
    }
    err := CreateAgent(agent)
    if err != nil{
      tpl.ExecuteTemplate(res,"result.html","Server encountered an error while creating user. Internal server error.")
      return
    }
    tpl.ExecuteTemplate(res,"createagent.html","Succefully created agent")
    return
  }
  tpl.ExecuteTemplate(res,"createagent.html",nil)
  return
}


//@TODO
//create a check to ensure agent is employed before allowing login
