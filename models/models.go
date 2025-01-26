package models

type Thread struct {
	ID uint;
	Title string;
	Category string;
	Content string; 
	UserID uint; 
}

type Comment struct {
	ID uint;
	ThreadID uint; 
	Content string; 
}

type User struct {
	ID uint;
	Username string;
}

