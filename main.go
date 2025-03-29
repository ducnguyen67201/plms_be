package main

import (
	"database/sql"
	user_app "plms_be/internal/application/user"
	user_domain "plms_be/internal/domain/user"
	user_oracle_db "plms_be/internal/infrastructure/persistence/user"
	user_http "plms_be/internal/interfaces/http/user"

	"github.com/gin-gonic/gin"
	_ "github.com/alexbrainman/odbc"
)

func main() {

	r := gin.Default()

	db, err := sql.Open("odbc", "DSN=OracleXE;UID=damg7275_final;PWD=damg7275_final")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := &user_oracle_db.OracleUserRepository{DB: db}
	userDomain := user_domain.NewService(repo)
	userService := &user_app.UserAppService{UserService: userDomain}

	user_http.RegisterRoutes(r, userService)

	r.Run(":8080")
}