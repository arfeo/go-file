package file

import (
  "io/ioutil"
  "log"
  "os"
)

/**
 * Function checks whether a file or directory with the given name exists and handles error (if any);
 * returns true if file/directory exists, otherwise returns false
 */
func Exists(name string) bool {
  if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
      return false
    }
  }

  return true
}

/**
 * Function creates a new file or directory with the given name and handles error (if any);
 * returns true if the file successfully created, otherwise returns false
 */
func Create(name string) bool {
  if _, err := os.Create(name); err != nil {
    log.Println(err)

    return false
  }

  return true
}

/**
 * Function tries to remove a file with the given name and handles error (if any);
 * returns true if the file deleted, otherwise returns false
 */
func Remove(name string) bool {
  if err := os.Remove(name); err != nil {
    log.Println(err)

    return false
  }

  return true
}

/**
 * Function tries to read data from a file with the given name and handles error (if any);
 * returns data as the first value and true as the second one in case of success,
 * otherwise returns nil as the first value and false as the second one
 */
func Read(name string) ([]byte, bool) {
  file, err := ioutil.ReadFile(name)

  if err != nil {
    log.Println(err)

    return nil, false
  }

  return file, true
}

/**
 * Function tries to write given data to a file with the given name and handles error (if any);
 * returns true if data was written to the file, otherwise returns false
 */
func Write(name string, data []byte) bool {
  if err := ioutil.WriteFile(name, data, 0644); err != nil {
    log.Println(err)

    return false
  }

  return true
}
