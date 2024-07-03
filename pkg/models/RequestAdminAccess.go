package models


func RequestAdminAccess(username string) string {
	db, err:= Connection()
	if err != nil {
		return "error connecting to db"
	}
	var isAdmin bool
	err = db.QueryRow("SELECT isadmin FROM users WHERE username=?", username).Scan(&isAdmin)
	if err != nil {
		db.Close()
		return "error checking user access"
	}
	if isAdmin {
		db.Close()
		return "User is already an admin"
	}

	db.QueryRow("SELECT admin_request FROM users WHERE username=?", username).Scan(&isAdmin)
	if isAdmin {
		db.Close()
		return "Admin access already requested"
	}

	_, err = db.Exec("UPDATE users SET admin_request=1 WHERE username=?", username)
	db.Close()
	if err != nil {
		return "error requesting admin access in db"
	}
	return "Admin access requested successfully"
}