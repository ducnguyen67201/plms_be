package main

import (
	"database/sql"
	"log"
	problem_app "plms_be/internal/application/problem"
	user_app "plms_be/internal/application/user"
	problem_domain "plms_be/internal/domain/problem"
	user_domain "plms_be/internal/domain/user"
	problem_db "plms_be/internal/infrastructure/persistence/problem"
	user_oracle_db "plms_be/internal/infrastructure/persistence/user"
	problem_http "plms_be/internal/interfaces/http/problem"
	user_http "plms_be/internal/interfaces/http/user"
	"time"

	_ "github.com/alexbrainman/odbc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	connStr := "Driver={Oracle in OraDB21Home1};Dbq=localhost:1521/xe;Uid=damg7275_final;Pwd=damg7275_final;"
	db, err := sql.Open("odbc", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("db.Ping failed:", err)
	}

	// User Service
	repo := &user_oracle_db.OracleUserRepository{DB: db}
	userDomain := user_domain.NewService(repo)
	userService := &user_app.UserAppService{UserService: userDomain}

	// Problem Service
	problemRepo := &problem_db.OracleProblemRepository{DB: db}
	problemDomain := problem_domain.NewProblemService(problemRepo)
	problemService := &problem_app.ProblemAppService{ProblemService: problemDomain}


	user_http.RegisterRoutes(r, userService)
	problem_http.RegisterProblemRoutes(r, problemService)
	
	r.Run(":8080")
}