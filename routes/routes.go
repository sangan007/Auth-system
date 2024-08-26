package routes
import (
    "auth-system/controllers"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)
func InitializeRoutes()*mux.Router{
    router := mux.NewRouter()
    db, err := gorm.Open(postgres.Open("your-postgres-connection-string"), &gorm.Config{})
    if err != nil{
        log.Fatal("Failed to connect to database")
    }
    controllers.SetDB(db)
    router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
    router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
    return router
}
