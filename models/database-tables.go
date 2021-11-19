package models

import "time"

type Reservation struct {
	ID int
	FirstName , LastName , Email , Phone string
	StartDate , EndDate time.Time
	RoomId int
	CreatedAt , UpdatedAt time.Time
	Room Room
}

type User struct {
	ID int
	FirstName , LastName , Email , Password string
	AccessLevel int
	CreatedAt , UpdatedAt time.Time
}

type Room struct {
	ID int
	RoomName string
	CreatedAt , UpdatedAt time.Time
}

type RoomRestriction struct {
	ID int
	StartDate , EndDate time.Time
	CreatedAt , UpdatedAt time.Time
	ReservationId , RestrictionId , RoomId int
	Reservation Reservation
	Restriction Restriction
}

type Restriction struct {
	ID int
	RestrictionName string
	CreatedAt , UpdatedAt time.Time
}