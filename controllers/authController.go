package controllers
import(
    "auth-system/models"
    "auth-system/utils"
    "encoding/json"
    "gorm.io/gorm"
    "net/http"
)

var db *gorm.DB
func SetDB(database *gorm.DB){
    db = database
    db.AutoMigrate(&models.User{})
}
func RegisterUser(w http.ResponseWriter, r *http.Request){
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }
    user.Password = hashedPassword
    result := db.Create(&user)
    if result.Error != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
func LoginUser(w http.ResponseWriter, r *http.Request){
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)
    var dbUser models.User
    result := db.Where("email = ?", user.Email).First(&dbUser)
    if result.Error != nil{
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    if !utils.CheckPasswordHash(user.Password, dbUser.Password){
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }
    token, err := utils.GenerateJWT(dbUser.Username)
    if err != nil{
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
