package models


type User struct {
	ID             int64   `json:"id" example:"1109659429"`
	FullName       string  `json:"full_name" example:"Jumaniyozov Ibodulla"`
	Username       string  `json:"username" example:"Ibodulla"`
	Language       string  `json:"language" enums:"uz, ru, en"`
	IsPremium      bool    `json:"is_premium" example:"false"`
	CashbackBalans float64 `json:"cashback_balans" example:"0"`
	IsAdmin        bool    `json:"is_admin" example:"false"`
	CreatedAt      string  `json:"created_at"`
	DeletedAt      string  `json:"deleted_at"`
}

type GetUser struct {
	ID             int64   `json:"id" example:"1109659429"`
	FullName       string  `json:"full_name" example:"Jumaniyozov Ibodulla"`
	Username       string  `json:"username" example:"Ibodulla"`
	Language       string  `json:"language" enums:"uz, ru, en"`
	IsPremium      bool    `json:"is_premium" example:"false"`
	CashbackBalans float64 `json:"cashback_balans" example:"0"`
	IsAdmin        bool    `json:"is_admin" example:"false"`
	CreatedAt      string  `json:"created_at"`
}

type AddUser struct {
	ID        int64  `json:"id" example:"1109659429"`
	FullName  string `json:"full_name" example:"Jumaniyozov Ibodulla"`
	Username  string `json:"username" example:"Ibodulla"`
	Language  string `json:"language" enums:"uz, ru, en"`
	IsPremium bool   `json:"is_premium" example:"false"`
}

type GetAllUsersRequest struct {
	SearchByFullName string `json:"search_by_full_name"`
	SearchByUsername string `json:"search_by_username"`
	SearchByID       int64  `json:"search_by_id"`
	Page             uint64 `json:"page"`
	Limit            uint64 `json:"limit"`
}

type GetAllUsersResponse struct {
	Users []GetUser `json:"users"`
	Count uint64    `json:"count"`
}

