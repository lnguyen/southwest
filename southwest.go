package main

import (
  "net/http"
  "strings"
  "log"
  "io/ioutil"
  "flag"
  "fmt"
  "net/url"
  "os"
)

//Southwest structure
type Southwest struct {
  FirstName string
  LastName string
  ConfirmationNumber string
  Url string
}

func NewSouthwest(firstName string, lastName string, 
                 confirmationNumber string, 
                  url string) *Southwest {
  southwest := Southwest{FirstName: firstName, LastName: lastName, ConfirmationNumber: confirmationNumber, Url: url}
  return &southwest
}

func (s *Southwest) CheckIn() error {
  //Create x-www-form-url-encoded
  // URL package
  v := url.Values{}
  v.Set("platform", "android")
  v.Set("firstName", s.FirstName)
  v.Set("lastName", s.LastName)
  v.Set("recordLocator", s.ConfirmationNumber)
  v.Set("serviceID", "flightcheckin_new")
  v.Set("appID", "swa")
  v.Set("appver", "2.4.1")
  v.Set("platformver", "5.0.GA_v201403042054")
  v.Set("channel", "rc")
  
  req, err := http.NewRequest("POST", s.Url, strings.NewReader(v.Encode()))
  
  if err != nil {
    log.Panic(err)
    return err
  }
  
  req.Header.Add("Content-type", "application/x-www-form-urlencoded")
 
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    log.Panic(err)
    return err
  }
  
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Panic(err)
    return err
  }
  fmt.Printf("Status code: %s\n", resp.StatusCode)
  fmt.Printf("Body: %s\n", string(body))
  return nil
}

func main() {
  var firstName string
  var lastName string
  var confirmationNumber string
  url := "http://mobile.southwest.com/middleware/MWServlet"
  
  flag.StringVar(&firstName, "firstName", "", "First name for check in")
  flag.StringVar(&lastName, "lastName", "", "Last name for check in")
  flag.StringVar(&confirmationNumber, "confirmationNumber", "", "Confirmation Number for check in")
  
  flag.Parse()
  
  if (firstName == "" || lastName == "" || confirmationNumber == "") {
    log.Panic("Please ensure first name, last name and confirmation number are filled out")
    os.Exit(1)
  }
  
  s := NewSouthwest(firstName, lastName, confirmationNumber, url)
  err := s.CheckIn()
  if err != nil {
    log.Panic(err)
    os.Exit(1)
  }
}

