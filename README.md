# Go-phone contacts REST API
This is phone contacts manager that was built using Golang & its standrard library **(net/http)**, used to demonstrate **CRUD** (Create, Read, Update, Delete) operations. The project here provides a RESTful API for managing phone contacts, interacting with a **PostgreSQL** database using the **pgxpool** PostgreSQL driver. 

Aims of this project:
- Showcase clean architecture in a Go application.
- Ability to implement RESTful API architecture.
- Database interaction through pgxpool Go driver.

# Features of API
- Create a new contect record. `/api/v1/addContact`
- Get a contact by id. `api/v1/getContact`
- Remove a contact by id. `api/v1/removeContact`
- Favorite a contact. `api/v1/favoriteContact`
- Unfavorite a contact `api/v1/unfavoriteContact`

 # Structure of project
.
├ /handlers
|  └── handlers.go
├ /modles
|  └── contact.go
├ /server
|  └── server.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── main.go
├── README.md

