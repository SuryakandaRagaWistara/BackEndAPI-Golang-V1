package models

type User struct {
	ID        *int64    `json:"id,omitempty"` // Gunakan pointer	CreatedAt *string  `json:"created_at,omitempty"` // Gunakan pointer    Name      string    `json:"name"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Role      string    `json:"role"`
    // Ini tetap ada. Ini bukan kolom di DB, 
    // melainkan tempat untuk menyimpan hasil join dari tabel Task.
    Tasks     []Task    `json:"tasks,omitempty"` 
}