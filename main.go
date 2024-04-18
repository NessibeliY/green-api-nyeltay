package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Data struct {
	IdInstance       string
	ApiTokenInstance string
	Result           string
}

const APIUrl = "https://7103.api.greenapi.com"

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	mux.HandleFunc("/", home)

	log.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	action := r.FormValue("action")
	if action == "/getSettings" {
		getSettings(w, r)
	} else if action == "/getStateInstance" {
		getStateInstance(w, r)
	} else if action == "/sendMessage" {
		sendMessage(w, r)
	} else if action == "/sendFileByUrl" {
		sendFileByUrl(w, r)
	} else {
		ts, _ := template.ParseFiles("html/index.html")
		ts.Execute(w, nil)
	}
}

func getSettings(w http.ResponseWriter, r *http.Request) {
	idInstance := r.PostFormValue("idInstance")
	apiTokenInstance := r.PostFormValue("apiTokenInstance")
	if idInstance == "" || apiTokenInstance == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	url := fmt.Sprintf("%s/waInstance%s/getSettings/%s", APIUrl, idInstance, apiTokenInstance)

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintf(w, "Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading response: %v", err)
		return
	}
	ts, _ := template.ParseFiles("html/index.html")
	data := &Data{
		IdInstance:       idInstance,
		ApiTokenInstance: apiTokenInstance,
		Result:           responseBody.String(),
	}
	ts.Execute(w, data)
}

func getStateInstance(w http.ResponseWriter, r *http.Request) {
	idInstance := r.PostFormValue("idInstance")
	apiTokenInstance := r.PostFormValue("apiTokenInstance")

	if idInstance == "" || apiTokenInstance == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", APIUrl, idInstance, apiTokenInstance)
	url := fmt.Sprintf("https://7103.api.greenapi.com/waInstance%s/getStateInstance/%s", idInstance, apiTokenInstance)

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintf(w, "Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading response: %v", err)
		return
	}
	ts, _ := template.ParseFiles("html/index.html")
	data := &Data{
		IdInstance:       idInstance,
		ApiTokenInstance: apiTokenInstance,
		Result:           responseBody.String(),
	}
	ts.Execute(w, data)
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	idInstance := r.PostFormValue("idInstance")
	apiTokenInstance := r.PostFormValue("apiTokenInstance")
	chatId := r.PostFormValue("chatId")
	message := r.PostFormValue("message")

	if idInstance == "" || apiTokenInstance == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if _, err := strconv.Atoi(chatId); err != nil {
		chatId = chatId + "@g.us"
	} else {
		chatId = chatId + "@c.us"
	}

	url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", APIUrl, idInstance, apiTokenInstance)

	payload := map[string]string{
		"chatId":  chatId,
		"message": message,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintf(w, "Error encoding JSON payload: %v", err)
		return
	}

	// Create request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Fprintf(w, "Error creating request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading response: %v", err)
		return
	}
	ts, _ := template.ParseFiles("html/index.html")
	data := &Data{
		IdInstance:       idInstance,
		ApiTokenInstance: apiTokenInstance,
		Result:           responseBody.String(),
	}
	ts.Execute(w, data)
}

func sendFileByUrl(w http.ResponseWriter, r *http.Request) {
	idInstance := r.PostFormValue("idInstance")
	apiTokenInstance := r.PostFormValue("apiTokenInstance")
	chatId := r.PostFormValue("chatId2")
	urlFile := r.PostFormValue("urlFile")

	if idInstance == "" || apiTokenInstance == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if _, err := strconv.Atoi(chatId); err != nil {
		chatId = chatId + "@g.us"
	} else {
		chatId = chatId + "@c.us"
	}

	url := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", APIUrl, idInstance, apiTokenInstance)

	payload := map[string]string{
		"chatId":   chatId,
		"urlFile":  urlFile,
		"fileName": "picture",
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintf(w, "Error encoding JSON payload: %v", err)
		return
	}

	// Create request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Fprintf(w, "Error creating request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading response: %v", err)
		return
	}
	ts, _ := template.ParseFiles("html/index.html")
	data := &Data{
		IdInstance:       idInstance,
		ApiTokenInstance: apiTokenInstance,
		Result:           responseBody.String(),
	}
	ts.Execute(w, data)
}
