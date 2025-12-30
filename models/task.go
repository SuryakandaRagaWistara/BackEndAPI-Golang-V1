package models

type TaskStatus string

const (
    StatusStarted    TaskStatus = "started"
    StatusInProgress TaskStatus = "in Progress"
    StatusCompleted  TaskStatus = "completed"
    StatusCanceled   TaskStatus = "canceled"
)

type Task struct {
    ID          *int64      `json:"id,omitempty"` // Wajib Pointer + omitempty	CreatedAt *string  `json:"created_at,omitempty"` //engunakan pointer nama
    
    // Field Skalar (Foreign Key)
    UserID      int64   `json:"user_id"` 
    
    // Field Relasional (Struct Association)
    User        *User   `json:"User,omitempty"` 
    
    Title       *string `json:"title"`
    Description *string `json:"description"`
    Deadline    *string `json:"deadline"`
    Status      TaskStatus `json:"status"`
}


// Relasi dalam go di representasikan dalam 2 bentuk
// Field skalar:untuk menyimpan nilai kunci misalnya int8
// Field Relasional:untuk menyimpan seluruh data terkait (misalnya struck user)