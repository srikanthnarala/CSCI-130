package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age  string
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func upload(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("main.html")
	if err != nil {
		log.Fatalln(err)
	}

	name := req.FormValue("name")
	age := req.FormValue("age")

	CurrUser := User{
		Name: name,
		Age:  age,
	}

	bs, err := json.Marshal(currentUser)
	if err != nil {
		fmt.Println(err)
	}

	json := base64.StdEncoding.EncodeToString(bs)

	cookie, err := req.Cookie("session")
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name:     "session",
		Value:    id.String() + name + age + json + getCode(id.String()),
		HttpOnly: true,
	}
	http.SetCookie(res, cookie)
	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", upload)
	http.ListenAndServe(":8040", nil)
}
