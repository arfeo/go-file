# go-file

Very simple package for working with files in `Go`.

Implements the next methods:

* Exists
* Create
* Remove
* Read
* Write

## Installing

```
$ go get github.com/arfeo/go-file
```

## Usage

```
package main

import (
  "encoding/json"
  "github.com/arfeo/go-file"
  "log"
)

type Config struct {
  User          string        `json:"user"`
  Password      string        `json:"password"`
}
	
func main() {
  var config Config
  configFileName := "config.json"
	
  if !file.Exists(configFileName) {
    log.Fatalln("Fatal error: Cannot find configuration file")
  }
  
  if data, ok := file.Read(configFileName); ok {
    if err := json.Unmarshal([]byte(data), &config); err != nil {
      log.Fatalln(err)
    }
  } else {
    log.Fatalln("Fatal error: Cannot read configuration file")
  }
}
```