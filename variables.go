package main

import (
  "time"
  "database/sql"
  "html/template"
)

var db *sql.DB
var now = time.Now()
var currentTime = now.Format("2006-01-02 15:04:05")
var tpl *template.Template
