package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	User User `json:"data"`
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserCreated struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Job       string     `json:"job"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deleteAt,omitempty"`
}

const (
	baseUrl = "https://reqres.in/api"
)

func main() {

	request, err := GetReqExample(fmt.Sprintf("%s/users/2", baseUrl))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(request))

	fmt.Println()

	request, err = Get(fmt.Sprintf("%s/users/2", baseUrl))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(request))

	fmt.Println()

	user, err := GetUser("2")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
	fmt.Println("ID: ", user.ID)
	fmt.Println("Email: ", user.Email)
	fmt.Println("FirstName: ", user.FirstName)
	fmt.Println("LastName: ", user.LastName)

	fmt.Println()

	u, err := CreateUser("jkna", "enginer")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)
	fmt.Println("ID: ", u.ID)
	fmt.Println("Name: ", u.Name)
	fmt.Println("Job: ", u.Job)
	fmt.Println("Date: ", u.CreatedAt)
}

func GetReqExample(url string) ([]byte, error) {

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func Get(url string) ([]byte, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode > 299 {
		return nil, fmt.Errorf("Status code %d", response.StatusCode)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return responseData, nil
}

func GetUser(userId string) (*User, error) {

	request, err := Get(fmt.Sprintf("%s/users/%s", baseUrl, userId))

	if err != nil {
		return nil, err
	}

	var dataResponse Data

	if err := json.Unmarshal(request, &dataResponse); err != nil {
		return nil, err
	}

	return &dataResponse.User, nil
}

func Post(url string, data interface{}) ([]byte, error) {

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func CreateUser(name, job string) (*UserCreated, error) {

	user := &UserCreated{
		Name: name,
		Job:  job,
	}

	req, err := Post(fmt.Sprintf("%s/users", baseUrl), user)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(req, user); err != nil {
		return nil, err
	}

	return user, nil
}
