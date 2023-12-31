package repos

import "time"

type AutoGenerated []struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	FullName   string    `json:"full_name"`
	Fork       bool      `json:"fork"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	PushedAt   time.Time `json:"pushed_at"`
	Language   string    `json:"language"`
	ForksCount int       `json:"forks_count"`
	Forks      int       `json:"forks"`
}

type License struct {
	Name string `json:"name"`
}
