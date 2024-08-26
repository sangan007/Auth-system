package main
import(
    "encoding/json"
    "log"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)
type User struct{
    ID        uint      `json:"id" gorm:"primaryKey"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"password" gorm:"not null"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
var db *gorm.DB
func init(){
    dsn := "host=localhost user=postgres password=sangan007 dbname=auth-system port=5432 sslmode=disable"
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != ni{
        log.Fatalf("Failed to connect to database: %v", err)
    }
    db.AutoMigrate(&User{})
}
func RegisterUser(w http.ResponseWriter, r *http.Request){
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err!= nil{
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := db.Create(&user).Error; err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
func LoginUser(w http.ResponseWriter, r *http.Request){
    var input User
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    var user User
    if err := db.Where("email = ? AND password = ?", input.Email, input.Password).First(&user).Error; err != nil{
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    json.NewEncoder(w).Encode(user)
}
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/register", RegisterUser).Methods("POST")
    r.HandleFunc("/login", LoginUser).Methods("POST")
    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
