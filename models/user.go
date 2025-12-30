package models

type User struct {
	ID        *int64    `json:"id,omitempty"` 
	Name      string    `json:"name"`
    Email     string    `json:"email"`
    Role      string    `json:"role"`
    //  tempat untuk menyimpan hasil join dari tabel Task.
    Tasks     []Task    `json:"tasks,omitempty"` 
}
