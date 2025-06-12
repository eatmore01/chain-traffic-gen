package main

type dd interface {
  hello(string)int
  method2(int)int
}

type DD struct {
  dd_field1 string
  dd_field2 int
  client &dd
}

var user = DD{}

user.client.hello("qwe")