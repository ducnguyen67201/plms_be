package main

import (
	"database/sql"
	"log"
	discussion_app "plms_be/internal/application/discussion"
	learning_app "plms_be/internal/application/learning"
	problem_app "plms_be/internal/application/problem"
	user_app "plms_be/internal/application/user"
	discussion_domain "plms_be/internal/domain/discussion"
	learning_domain "plms_be/internal/domain/learning"
	problem_domain "plms_be/internal/domain/problem"
	user_domain "plms_be/internal/domain/user"
	discussion_db "plms_be/internal/infrastructure/persistence/discussion"
	learning_db "plms_be/internal/infrastructure/persistence/learning"
	problem_db "plms_be/internal/infrastructure/persistence/problem"
	user_oracle_db "plms_be/internal/infrastructure/persistence/user"
	discussion_http "plms_be/internal/interfaces/http/dicussion"
	learning_http "plms_be/internal/interfaces/http/learning_material"
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

	// Discussion Service
	discussionRepo := &discussion_db.OracleDiscussionRepository{DB: db}
	dicussDomain := discussion_domain.NewDiscussionService(discussionRepo)
	discussionService := &discussion_app.DiscussionAppService{DiscussionService: dicussDomain}
	
	// Learning Service
	leraningRepo := &learning_db.OracleLearningRepository{DB: db}
	learningDomain := learning_domain.NewLearningService(leraningRepo)
	learningService := &learning_app.LearningAppService{LearningService: learningDomain}

	user_http.RegisterRoutes(r, userService)
	problem_http.RegisterProblemRoutes(r, problemService)
	discussion_http.RegisterDiscussionRoutes(r, discussionService)
	learning_http.RegisterLearningRoutes(r, learningService)

	r.Run(":8080")
}