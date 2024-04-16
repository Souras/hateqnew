package db

import (
	"database/sql"
	"time"

	"github.com/Souras/hateqnew/service/hateqService/internal/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "192.168.1.12"
	port     = 5432
	user     = "casaos"
	password = "casaos"
	dbname   = "hateq"
)

func InitDb() error {
	connectionString := "user=casaos dbname=hateq password=casaos host=192.168.1.12 port=5432 sslmode=disable" // Replace with your credentials

	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	return nil

	// Connect to the PostgreSQL database
	// dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", dbinfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// // Ensure the database connection is valid
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	// return nil
}

func GetProducts(date string) ([]models.QueueData, error) {
	var query = ""
	var rows *sql.Rows
	var err error
	if date == "today" {
		query = "SELECT * FROM public.tokens WHERE insert_time >= CURRENT_DATE AND insert_time < CURRENT_DATE + interval '1 day' ORDER BY id ASC;"
		rows, err = db.Query(query)
	} else if date != "" {
		query = "SELECT * FROM public.tokens WHERE insert_time > $1 ORDER BY id ASC"
		rows, err = db.Query(query, date)
	} else {
		query = "SELECT * FROM public.tokens ORDER BY id ASC"
		rows, err = db.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.QueueData
	for rows.Next() {
		var p models.QueueData
		err := rows.Scan(&p.ID, &p.TokenNur, &p.Name, &p.IsActive, &p.IsCancelled, &p.TimeSlot, &p.AdminID, &p.MobileNo, &p.InsertTime, &p.StartTime, &p.EndTime, &p.Operating, &p.OsVersion, &p.Duration)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProduct(id int64) (models.QueueData, error) {
	row := db.QueryRow("SELECT * FROM tokens WHERE id = $1", id)
	var p models.QueueData
	err := row.Scan(&p.ID, &p.TokenNur, &p.Name, &p.IsActive, &p.IsCancelled, &p.TimeSlot, &p.AdminID, &p.MobileNo, &p.InsertTime, &p.StartTime, &p.EndTime, &p.Operating, &p.OsVersion, &p.Duration)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.QueueData{}, err
		}
		return models.QueueData{}, err
	}
	return p, nil
}

func CreateProduct(p models.QueueData, insertTime time.Time, startTime time.Time, endTime time.Time) (models.QueueData, error) {

	// query := `INSERT INTO tokens (tokennur, name, isactive, iscancelled, timeslot, adminid, mobileno, inserttim, starttime, endtime, operating, osversion, duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING tokennur`
	// var ID int
	// err := db.QueryRow(query, p.TokenNur, p.Name, p.IsActive, p.IsCancelled, p.TimeSlot, p.AdminID, p.MobileNo, insertTime, startTime, endTime, p.Operating, p.OsVersion, p.Duration).Scan(&ID)
	// if err != nil {

	// 	return models.QueueData{}, err
	// }

	stmt, err := db.Prepare("INSERT INTO tokens (token_number, name, is_active, is_cancelled, time_slot, admin_id, mobile_no, insert_time, start_time, end_time, operating, os_version, duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id")
	if err != nil {
		return models.QueueData{}, err
	}
	defer stmt.Close()

	var ID int64
	err = stmt.QueryRow(p.TokenNur, p.Name, p.IsActive, p.IsCancelled, p.TimeSlot, p.AdminID, p.MobileNo, insertTime, startTime, endTime, p.Operating, p.OsVersion, p.Duration).Scan(&ID)
	if err != nil {
		return models.QueueData{}, err
	}

	p.ID = ID
	return p, nil
}

func UpdateProduct(id int64, p models.QueueData) (models.QueueData, error) {
	stmt, err := db.Prepare("UPDATE tokens SET name = $1 WHERE id = $2 RETURNING id")
	if err != nil {
		return models.QueueData{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, id)
	if err != nil {
		return models.QueueData{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.QueueData{}, err
	}

	if rowsAffected == 0 {
		return models.QueueData{}, sql.ErrNoRows
	}

	p.ID = id
	return p, nil
}

func DeleteProduct(id int64) error {
	stmt, err := db.Prepare("DELETE FROM tokens WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
