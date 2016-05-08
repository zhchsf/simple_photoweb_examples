package handlers

import (
  "io"
  // "log"
  "net/http"
  "html/template"
  // "path/filepath"
  "os"
  "io/ioutil"
  "fmt"
)

const (
  UPLOAD_DIR = "./uploads"
)

func UploadHandler(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET" {
    // cwd, _ := os.Getwd()
    t, err := template.ParseFiles("html/upload.html")
    check(err)
    t.Execute(w, nil)
    return
  }

  if r.Method == "POST" {
    f, h, err := r.FormFile("image")
    check(err)
    filename := h.Filename
    defer f.Close()

    t, err := os.Create(UPLOAD_DIR + "/" + filename)
    check(err)
    defer t.Close()

    _, err = io.Copy(t, f)
    check(err)
    http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
  }
}


func ViewHandler(w http.ResponseWriter, r *http.Request) { 
  imageId := r.FormValue("id")
  imagePath := UPLOAD_DIR + "/" + imageId
  if exists := isExists(imagePath);!exists {
    http.NotFound(w, r)
    return
  }
  w.Header().Set("Content-Type", "image")
  http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
  _, err := os.Stat(path)
  if err == nil {
    return true
  }
  return os.IsExist(err)
}

func ListHandler(w http.ResponseWriter, r *http.Request){
  fileInfoArr, err := ioutil.ReadDir("./uploads")
  check(err)
  locals := make(map[string]interface{})
  images := []string{}
  for _, fileInfo := range fileInfoArr{
    fmt.Printf(fileInfo.Name())
    images = append(images, fileInfo.Name())
  }
  // for _, vvv := range images {
  //   fmt.Printf(vvv)
  // }
  locals["images"] = images

  t, err := template.ParseFiles("html/list.html")
  check(err)
  t.Execute(w, locals)
  return
}

func check(err error){
  if err != nil {
    // http.Error(w, err.Error(), http.StatusInternalServerError)
    panic(err)
    // return
  }
}
