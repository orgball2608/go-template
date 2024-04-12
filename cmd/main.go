package main

import (
	"context"
	"elearning/config"
	"elearning/docs"
	"elearning/infra/mysql"
	"elearning/infra/mysql/repository"
	asynq "elearning/infra/queue"
	"elearning/infra/redis"
	cachedRepository "elearning/infra/redis/repository"
	"elearning/middlewares"
	"elearning/pkg/data"
	stringPkg "elearning/pkg/hasher"
	"elearning/routers"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	queueRepository "elearning/infra/queue/repository"
	jwtPkg "elearning/pkg/jwt"
	"elearning/usecase/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"reflect"
)

type App struct {
	config        *config.Environment
	database      *mysql.Database
	redisDatabase *redis.Database
	queueClient   *asynq.QueueClient
	logger        *logrus.Entry
}

var listEnvSecret = []string{
	"Constants",
	"MysqlPassword",
	"RedisPassword",
	"AccessTokenSecretKey",
	"RefreshTokenSecretKey",
	"AwsSecretAccessKey",
}

func main() {
	logger := initLog()
	loggerStartServer := initStartServerLog()
	cfg := loadEnvironment()
	v := reflect.ValueOf(cfg).Elem()
	for i := 0; i < v.NumField(); i++ {
		varName := v.Type().Field(i).Name
		varValue := v.Field(i).Interface()
		isLog := true
		for _, envSecret := range listEnvSecret {
			if varName == envSecret {
				isLog = false
				break
			}
		}
		if isLog {
			fmt.Printf("EnvKeyAndValue %s: '%v'\n", varName, varValue)
		}
	}

	//Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "DDD API"
	docs.SwaggerInfo.Description = "This is a DDD server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080/api/v1"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	gin.SetMode(cfg.RunMode)
	loggerStartServer.Infof("System is running with %s mode", cfg.RunMode)
	// Connect to database
	db, err := mysql.ConnectMysql(cfg)
	if err != nil {
		loggerStartServer.Fatalf("Connect Mysql Server Failed. Error: %s", err.Error())
	}
	loggerStartServer.Infof("Connect Mysql Server Successfully")

	app := &App{
		config:   cfg,
		database: db,
		logger:   logger,
	}

	if app.config.MysqlMigrateMode {
		loggerStartServer.Info("Start Migrating")
		err = app.database.AutoMigrate()
		if err != nil {
			loggerStartServer.Fatal("Migrating Failed. Error: " + err.Error())
		}
		loggerStartServer.Info("Migrating Successfully")
	}

	// Service
	dataService := data.NewDataService()
	stringService := stringPkg.NewStringService()
	jwtService := jwtPkg.NewJwtService(app.config)

	// Repository
	userRepository := repository.NewUserRepository(app.database, dataService)
	cacheRepository := cachedRepository.NewCacheRepository(app.redisDatabase)
	userTokenRepository := repository.NewUserTokenRepository(app.database, dataService)
	queueRepository := queueRepository.NewQueueRepository(app.queueClient)
	authUseCase := auth.NewAuthUseCase(
		*app.config,
		jwtService,
		stringService,
		userRepository,
		cacheRepository,
		userTokenRepository,
		queueRepository,
	)

	middleware := middlewares.NewMiddleware(
		jwtService,
		stringService,
		app.logger,
		userRepository,
		*app.config,
	)

	router := routers.InitRouter(
		app.config,
		middleware,
		authUseCase,
	)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.Port),
		Handler: router,
	}
	done := make(chan bool)
	go func() {
		if err := GracefulShutDown(app.config, done, server); err != nil {
			loggerStartServer.Infof("Stop server shutdown error: %v", err.Error())
			return
		}
		loggerStartServer.Info("Stopped serving on Services")
	}()
	loggerStartServer.Infof("Start HTTP Server Successfully")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		loggerStartServer.Fatalf("Start HTTP Server Failed. Error: %s", err.Error())
	}
	<-done
	loggerStartServer.Infof("Stopped backend application.")
}

func GracefulShutDown(config *config.Environment, quit chan bool, server *http.Server) error {
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.SystemTimeOutSecond)*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	close(quit)
	return nil
}

func initLog() *logrus.Entry {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		DisableQuote:    true,
		DisableColors:   true,
		FieldMap: logrus.FieldMap{
			"level": "logLevel",
		},
	})
	log := logrus.WithFields(logrus.Fields{
		"module": "backend",
	})
	return log
}

func initStartServerLog() *logrus.Entry {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		DisableQuote:    true,
		DisableColors:   true,
		DisableSorting:  true,
		FieldMap: logrus.FieldMap{
			"level": "logLevel",
		},
	})
	log := logrus.WithFields(logrus.Fields{
		"module":  "backend",
		"logType": "startServer",
	})
	return log
}

func loadEnvironment() *config.Environment {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal("Fail loading environment variables: ", err)
	}
	return cfg
}
