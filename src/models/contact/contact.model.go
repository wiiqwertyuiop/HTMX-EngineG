package contact_model

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Contact struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

// For demo purposes, uses cookie as "database"
func GetContact(r *http.Request) Contact {
	cookie, err := r.Cookie("contact")
	if err != nil {
		return Contact{
			FirstName: "Joe", LastName: "Blow", Email: "joe@blow.com",
		}
	}
	// Decode saved contact
	saved := Contact{}
	decoded, _ := url.QueryUnescape(cookie.Value)
	json.Unmarshal([]byte(decoded), &saved)
	return saved
}

// For demo purposes, uses cookie as "database"
func UpdateContact(w http.ResponseWriter, r *http.Request) Contact {
	// Read form data
	if err := r.ParseForm(); err != nil {
		return GetContact(r) // Return orginal contact if there is an error
	}

	updatedContact := Contact{
		FirstName: r.FormValue("firstName"),
		LastName:  r.FormValue("lastName"),
		Email:     r.FormValue("email"),
	}

	// Save data to cookie
	data, _ := json.Marshal(updatedContact)
	http.SetCookie(w, &http.Cookie{Name: "contact", Value: url.QueryEscape(string(data))})

	return updatedContact
}
