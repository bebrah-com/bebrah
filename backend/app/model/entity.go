package model

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName  string    `gorm:"type:VARCHAR(256);unique;not null" json:"user_name"`
	Email     string    `gorm:"type:VARCHAR(256);unique;not null" json:"email"`
	Password  string    `gorm:"type:VARCHAR(128);not null" json:"password"`
	Avatar    string    `gorm:"type:BLOB" json:"avatar"`
	Info      string    `gorm:"type:TEXT" json:"info"`
	Banner    string    `gorm:"type:BLOB" json:"banner"`
	CreatedAt time.Time `gorm:"created_at;autoCreateTime" json:"created_at"`
	Token     string    `gorm:"type:TEXT" json:"token"`
}

type Work struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64     `gorm:"index;not null" json:"user_id"`
	User      *User      `gorm:"foreignKey:UserID" json:"user"`
	Data      string     `gorm:"type:BLOB;not null" json:"data"`
	WorkName  string     `gorm:"type:VARCHAR(256);not null" json:"work_name"`
	WorkDesc  string     `gorm:"type:TEXT" json:"work_desc"`
	CreatedAt time.Time  `gorm:"created_at;autoCreateTime" json:"created_at"`
	DeletedAt *time.Time `gorm:"deleted_at;type:TIMESTAMP" json:"deleted_at"`
	Viewed    uint64     `gorm:"default:0" json:"viewed"`
}

type Like struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64     `gorm:"index;not null" json:"user_id"`
	User      *User      `gorm:"foreignKey:UserID" json:"user"`
	WorkID    uint64     `gorm:"index;not null" json:"work_id"`
	Work      *Work      `gorm:"foreignKey:WorkID" json:"work"`
	CreatedAt time.Time  `gorm:"created_at;autoCreateTime" json:"created_at"`
	DeletedAt *time.Time `gorm:"deleted_at;type:TIMESTAMP" json:"deleted_at"`
}

type Follow struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	FollowedID uint64     `gorm:"index;not null" json:"followed_id"`
	Followed   *User      `gorm:"foreignKey:FollowedID" json:"followed"`
	FollowerID uint64     `gorm:"index;not null" json:"follower_id"`
	Follower   *User      `gorm:"foreignKey:FollowerID" json:"follower"`
	CreatedAt  time.Time  `gorm:"created_at;type:TIMESTAMP;autoCreateTime" json:"created_at"`
	DeletedAt  *time.Time `gorm:"deleted_at;type:TIMESTAMP" json:"deleted_at"`
}
