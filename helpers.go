package main

/*
  * Contains helper functions
*/

import (
  "os"
  "io"
  "log"
  "time"
  "strings"
  "math/rand"
)

//Check if a string is empty returns True if string is a string
func CheckifStringIsEmpty(data string) bool{
  if len(strings.TrimSpace(data)) == 0{
    return false
  }
  if len(data) == 0{
    return false
  }
  return true
}


func RandString(length int) string{
  var output strings.Builder
  rand.Seed(time.Now().Unix())
  charset := []rune("QWERTYUIOPLKJHGFDSAZXCVBNM123456789qwertyuioplkjhgfdsazxcvbnm")
  for i := 0; i < length; i++{
    random := rand.Intn(len(charset))
    randomChar := charset[random]
    output.WriteRune(randomChar)
  }
  id := output.String()
  id = strings.ToUpper(id)
  if !CheckifStringIsEmpty(id){
    RandString(length)
  }
  return id
}

//Retrns a random string with numbers and letters (caps on)
func RandNoLetter(length int) string{
  var output strings.Builder
  rand.Seed(time.Now().Unix())
  charset := []rune("QWERTYUIOPLKJHGFDSAZXCVBNM123456789")
  for i := 0; i < length; i++{
    random := rand.Intn(len(charset))
    randomChar := charset[random]
    output.WriteRune(randomChar)
  }
  id := output.String()
  id = strings.ToUpper(id)
  if !CheckifStringIsEmpty(id){
    RandNoLetter(length)
  }
  return id
}

//Returns A Random letters
func RandLetters(length int) string{
  var output strings.Builder
  rand.Seed(time.Now().Unix())
  charset := []rune("qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBBNM")
  for i := 0; i < length; i++{
    random := rand.Intn(len(charset))
    randomChar := charset[random]
    output.WriteRune(randomChar)
  }
  id := output.String()
  if !CheckifStringIsEmpty(id){
    RandLetters(length)
  }
  return id
}

//Returns a random number in string format
func RandNo(length int) string{
  var output strings.Builder
  rand.Seed(time.Now().Unix())
  charset := []rune("1234567890")
  for i := 0; i < length; i++{
    random := rand.Intn(len(charset))
    randomChar := charset[random]
    output.WriteRune(randomChar)
  }
  id := output.String()
  if !CheckifStringIsEmpty(id){
    RandNo(length)
  }
  return id
}

//Log and error to file Allows format string input
func LogErrorToFile(name string,text ...interface{}) error{
  name = "./.data/logs/"+name+".log"
  f,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
  if err != nil{
    return err
  }
  defer f.Close()
  writer := io.MultiWriter(os.Stdout,f)
  log.SetOutput(writer)
  log.Println(text)
  return nil
}

//log error of errors
func logError(e error){
  if e != nil {
    log.Println(e)
  }
}
