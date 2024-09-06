package main

import (
	"fmt"
	"goproj/utils"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)
type User struct {
	Name string
	Password int
	Authorized bool
	Urls []string
}
type Imag struct {
	Url []string
}
type Page struct {
	Usrautho bool
	Userurl []string
	Imageutl []string
}
var u User

func home_page(w http.ResponseWriter, r *http.Request){
	//http.ServeFile(page, r, "templates/index.html")
	var page Page
	var i Imag
	db, _ := utils.Connect()
	defer db.Close()
	i.Url, _ = utils.ShowProd(db)
	page.Usrautho = u.Authorized
	page.Userurl = u.Urls
	page.Imageutl = i.Url
	ch := u.Authorized 
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, page)
	fmt.Println(ch)

}
func gay_house(w http.ResponseWriter, r *http.Request){
	//http.ServeFile(w, r, "templates/gayhouse.html")
	//name := r.URL.Query().Get("Name")
    // if name == "" {
    //     http.Error(w, "Name is required", http.StatusBadRequest)
    //     return
    // }
    //u := User{Name: name}
	//u.Name = name
	db, _ := utils.Connect()
	defer db.Close()
	u.Urls, _ = utils.GetProd(db, u.Name)
	fmt.Println(u.Urls)
    t, _ := template.ParseFiles("templates/gayhouse.html")
    t.Execute(w, u)
}
func Save(w http.ResponseWriter, r *http.Request){
	Name := r.FormValue("Name")
	Passwordstr := r.FormValue("Password")
	Password, _ := utils.HashPassword(Passwordstr)
	db, err := utils.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = utils.InsertRow(db, Name, Password)
    if err != nil {
        log.Fatal("Error inserting row: ", err)
    }
	http.Redirect(w,r,"/",301)
    //fmt.Println("Row inserted successfully")
}
func Check(w http.ResponseWriter, r *http.Request){
	Name := r.FormValue("Name")
	Pass := r.FormValue("Password")
	db, err := utils.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	exists, err := utils.Checker(db, Name)
	if err != nil {
        log.Println("Error checking user:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    if utils.CheckPasswordHash(Pass, exists) {
		u.Authorized = true
		u.Name = Name
		//http.Redirect(w, r, fmt.Sprintf("/gayhouse?Name=%s", Name), http.StatusSeeOther)
		http.Redirect(w, r, "/gayhouse", http.StatusSeeOther)
    } 
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // 10MB
    file, handler, _:= r.FormFile("image")
    defer file.Close()

    // Создание файла на сервере
    filePath := "./uploads/" + handler.Filename
    dst, _ := os.Create(filePath)
    defer dst.Close()

    io.Copy(dst, file)

    // Формируем URL изображения
    imageURL := "./uploads/" + handler.Filename

    // Сохранение URL в базе данных
    db, _:= utils.Connect()
    defer db.Close()
	utils.InsertProd(db, imageURL)
	http.Redirect(w, r, "/gayhouse", http.StatusSeeOther)
}
func SaveImg(w http.ResponseWriter, r *http.Request){
	name := u.Name
    url := r.FormValue("Url")
	fmt.Println(url)
    db, err := utils.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    utils.AddProd(db, url, name)

    http.Redirect(w, r, "/gayhouse", http.StatusSeeOther) 
}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
    // Очистка cookie, которые хранят данные авторизации
    cookie := &http.Cookie{
        Name:    "session_token",
        Value:   "",
        Expires: time.Now().Add(-1 * time.Hour),
    }
    http.SetCookie(w, cookie)
    u.Authorized = false
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
func handreq(){
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/uploads/",http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))
	http.Handle("http://localhost:8080/uploads/",http.StripPrefix("http://localhost:8080/uploads/", http.FileServer(http.Dir("http://localhost:8080/uploads/"))))
	http.HandleFunc("/", home_page)
	http.HandleFunc("/gayhouse", gay_house)
	http.HandleFunc("/save", Save)
	http.HandleFunc("/check", Check)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/suser", SaveImg)
	http.HandleFunc("/logout", logoutHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handreq()
}