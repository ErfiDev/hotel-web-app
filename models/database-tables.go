package models

import "time"

type Reservations struct {
	ID int
	First_name , Last_name , Email , Phone string
	Start_date , End_date time.Time
	Room_id int
	Created_at , Updated_at time.Time
}

type Users struct {
	ID int
	First_name , Last_name , Email , Password string
	Access_level int
	Created_at , Updated_at time.Time
}

type Rooms struct {
	ID int
	Room_name string
	Created_at , Updated_at time.Time
}

type RoomRestrictions struct {
	ID int
	Start_date , End_date time.Time
	Created_at , Updated_at time.Time
	Reservation_id , Restriction_id , Room_id int
}

type Restrictions struct {
	ID int
	Restriction_name string
	Created_at , Updated_at time.Time
}