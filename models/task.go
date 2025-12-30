package models

type TaskStatus string

const (
    StatusStarted    TaskStatus = "started"
    StatusInProgress TaskStatus = "in Progress"
    StatusCompleted  TaskStatus = "completed"
    StatusCanceled   TaskStatus = "canceled"
)

type Task struct {
    ID          *int64      `json:"id,omitempty"` 
    // Field Skalar (Foreign Key)
    UserID      int64   `json:"user_id"` 
    // Field Relasional (Struct Association)
    User        *User   `json:"User,omitempty"` 
    Title       *string `json:"title"`
    Description *string `json:"description"`
    Deadline    *string `json:"deadline"`
    Status      TaskStatus `json:"status"`
}

/*
Catatan:
- Relasi dalam Go direpresentasikan 2 bentuk:
  1. Field skalar: menyimpan nilai kunci (contoh: UserID int64)
  2. Field relasional: menyimpan seluruh data terkait (contoh: User *User)
- Semua field pointer digunakan untuk optional dan mendukung omitempty saat JSON marshal
*/
