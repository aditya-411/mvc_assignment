package models


func RequestAdminAccess(username string) string {
	db, err:= Connection()
	if err != nil {
		return "error connecting to db"
	}
	defer db.Close()
	var isAdmin bool
	err = db.QueryRow("SELECT isadmin FROM users WHERE username=?", username).Scan(&isAdmin)
	if err != nil {
		return "error checking user access"
	}
	if isAdmin {
		return "User is already an admin"
	}

	db.QueryRow("SELECT admin_request FROM users WHERE username=?", username).Scan(&isAdmin)
	if isAdmin {
		return "Admin access already requested"
	}

	_, err = db.Exec("UPDATE users SET admin_request=1 WHERE username=?", username)
	if err != nil {
		return "error requesting admin access in db"
	}
	return "Admin access requested successfully"
}