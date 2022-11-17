package main

import (
	"context"
	"log"
	_driverFactory "mini-project/drivers"
	"mini-project/util"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	// _userUseCase "mini-project/businesses/users"
	// _userController "mini-project/controllers/users"

	_studentUseCase "mini-project/businesses/students"
	_studentController "mini-project/controllers/students"

	_moduleUseCase "mini-project/businesses/modules"
	_moduleController "mini-project/controllers/modules"

	_teacherUseCase "mini-project/businesses/teachers"
	_teacherController "mini-project/controllers/teachers"

	_assignmentUsecase "mini-project/businesses/assignments"
	_assignmentController "mini-project/controllers/assignments"

	_dbDriver "mini-project/drivers/mysql"

	_middleware "mini-project/app/middlewares"
	_routes "mini-project/app/routes"

	echo "github.com/labstack/echo/v4"
)

type operation func(ctx context.Context) error

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: "bffa768996b0ac",
		DB_PASSWORD: "58d94e32",
		DB_HOST:     "us-cdbr-east-06.cleardb.net",
		DB_PORT:     "3306",
		DB_NAME:     "heroku_706e7c18af76f4d",
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	studentRepo := _driverFactory.NewStudentRepository(db)
	studentUsecase := _studentUseCase.NewStudentUsecase(studentRepo, &configJWT)
	studentCtrl := _studentController.NewStudentController(studentUsecase)

	moduleRepo := _driverFactory.NewModuleRepository(db)
	moduleUsecase := _moduleUseCase.NewModuleUsecase(moduleRepo, &configJWT)
	moduleCtrl := _moduleController.NewModuleController(moduleUsecase)

	teacherRepo := _driverFactory.NewTeacherRepository(db)
	teacherUsecase := _teacherUseCase.NewTeacherUsecase(teacherRepo, &configJWT)
	teacherCtrl := _teacherController.NewTeacherController(teacherUsecase)

	assignmentRepo := _driverFactory.NewAssignmentRepository(db)
	assignmentUsecase := _assignmentUsecase.NewAssignmentUsecase(assignmentRepo, &configJWT)
	assignmentCtrl := _assignmentController.NewAssignmentController(assignmentUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:     configLogger.Init(),
		JWTMiddleware:        configJWT.Init(),
		StudentController:    *studentCtrl,
		ModuleController:     *moduleCtrl,
		TeacherController:    *teacherCtrl,
		AssignmentController: *assignmentCtrl,
	}

	routesInit.RouteRegister(e)

	go func() {
		if err := e.Start(":80"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	wait := gracefulShutdown(context.Background(), 10*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return _dbDriver.CloseDB(db)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(context.Background())
		},
	})

	<-wait
}

// gracefulShutdown performs application shut down gracefully.
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
