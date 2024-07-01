package userdb

import (
	DBConnection "Ecomm/pkg/db/DbConnection"
	"Ecomm/pkg/models"
	"errors"
)

func AddNewUser(user models.User) (int, error) {

	DB, err := DBConnection.DBConnection()

	if err != nil {
		return 0, errors.New(err.Error())
	}

	defer DB.Close()

	query := "INSERT INTO users (First_Name,Last_Name,User_Name,Password_Hash,Email_Id,isAdmin) VALUES(?,?,?,?,?,?)"

	result, err := DB.Exec(query, user.First_name, user.Last_name, user.User_name, user.Password, user.Email, user.IsAdmin)

	if err != nil {
		return 0, errors.New("Unable to insert data into database " + err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("Unable to fetch last inserted id" + err.Error())
	}

	return int(id), nil
}

func GetAllUsers() ([]models.User, error) {

	var Users []models.User

	DB, err := DBConnection.DBConnection()
	if err != nil {
		return nil, errors.New("unable to connect to Database")
	}

	query := "SELECT id,First_name,Last_name,User_name,Password_Hash,Email_id,isAdmin FROM users"

	rows, err := DB.Query(query)
	if err != nil {
		return nil, errors.New("Falied to execute the query " + err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.User_name, &user.Password, &user.Email, &user.IsAdmin)
		Users = append(Users, user)
	}

	return Users, nil

}

func UpdateUserDetails(user models.User) error {

	DB, err := DBConnection.DBConnection()
	if err != nil {
		return errors.New("unable to connect to Database")
	}

	defer DB.Close()

	query := "UPDATE users SET First_name=?,Last_name=?,User_name=?,Password_Hash=?,Email_id=?,Updated_at=CURRENT_TIMESTAMP WHERE id=?"

	_, err = DB.Exec(query, user.First_name, user.Last_name, user.User_name, user.Password, user.Email, user.Id)
	if err != nil {
		return err
	}

	return nil
}
