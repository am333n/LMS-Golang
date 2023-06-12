package managers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"github.com/gorilla/mux"
// 	"gorm.io/gorm"
// 	dc"lms/Database"
// 	l"lms/Login"
// )



// func GetManagers(w http.ResponseWriter, r *http.Request) {
// 	l.ValidateToken(w,r)
// 	w.Header().Set("Content-Type", "application/json")
// 	var managers []Manager
// 	dc.DB.Find(&managers)
// 	json.NewEncoder(w).Encode(managers)
// }
// func GetManager(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var manager Manager
// 	params := mux.Vars(r)
// 	dc.DB.First(&manager, params["id"])
// 	json.NewEncoder(w).Encode(manager)
// }
// func DeleteManager(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var manager Manager
// 	dc.DB.Delete(&manager, params["id"])
// 	json.NewEncoder(w).Encode("the manager has been deleted")
// }

// func PostManager(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var manager Manager
// 	json.NewDecoder(r.Body).Decode(&manager)
// 	dc.DB.Create(&manager)
// 	json.NewEncoder(w).Encode(manager)
// }
// func PutManager(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var manager Manager
// 	dc.DB.First(&manager, params["id"])
// 	json.NewDecoder(r.Body).Decode(&manager)
// 	dc.DB.Save(&manager)
// 	json.NewEncoder(w).Encode(manager)
// }
