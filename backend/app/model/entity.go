package model

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	UserName  string    `gorm:"type:VARCHAR(256);unique;not null"`
	Email     string    `gorm:"type:VARCHAR(256);unique;not null"`
	Password  string    `gorm:"type:VARCHAR(128);not null"`
	CreatedAt time.Time `gorm:"created_at;autoCreateTime"`
	Token     string    `gorm:"type:TEXT"`
}

type Work struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	UserID    uint64     `gorm:"index;not null"`
	User      *User      `gorm:"foreignKey:UserID"`
	Data      string     `gorm:"type:BLOB;not null"`
	WorkName  string     `gorm:"type:VARCHAR(256);not null"`
	WorkDesc  string     `gorm:"type:TEXT"`
	CreatedAt time.Time  `gorm:"created_at;autoCreateTime"`
	DeletedAt *time.Time `gorm:"deleted_at;type:TIMESTAMP"`
}

type Like struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	UserID    uint64     `gorm:"index;not null"`
	User      *User      `gorm:"foreignKey:UserID"`
	WorkID    uint64     `gorm:"index;not null"`
	Work      *Work      `gorm:"foreignKey:WorkID"`
	CreatedAt time.Time  `gorm:"created_at;autoCreateTime"`
	DeletedAt *time.Time `gorm:"deleted_at;type:TIMESTAMP"`
}

type Follow struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement"`
	FollowedID uint64     `gorm:"index;not null"`
	Followed   *User      `gorm:"foreignKey:FollowedID"`
	FollowerID uint64     `gorm:"index;not null"`
	Follower   *User      `gorm:"foreignKey:FollowerID"`
	CreatedAt  time.Time  `gorm:"created_at;type:TIMESTAMP;autoCreateTime"`
	DeletedAt  *time.Time `gorm:"deleted_at;type:TIMESTAMP"`
}
