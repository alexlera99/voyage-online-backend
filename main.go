package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	_ "github.com/rs/cors"
	"net/http"
)

var users map[string]string

var lastUser string

type User struct {
	Username string
	Password string
}

type Offers struct {
	Offers []Offer
}

type Offer struct {
	Image string
	Title string
	Text string
	Price string
}

func main (){
	users = make(map[string]string)
	users["alex"] = "alex"
	router := mux.NewRouter()
	router.HandleFunc("/", hi)
	router.HandleFunc("/images", images)
	router.HandleFunc("/lastuser", lastuser)
	router.HandleFunc("/login", login)
	router.HandleFunc("/register", register)
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}

func hi (response http.ResponseWriter, request *http.Request){
}

func login (response http.ResponseWriter, request *http.Request){
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	_, ok := users[user.Username]
	if ok {
		if users[user.Username] == user.Password {
			lastUser = user.Username
			response.Write([]byte("Login"))
			return
		}
		return
	} else {
		return
	}
}

func register (response http.ResponseWriter, request *http.Request){
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	_, ok := users[user.Username]
	if ok {
		return
	} else {
		users[user.Username] = user.Password
		response.Write([]byte("Register"))
		return
	}
}

func lastuser (response http.ResponseWriter, request *http.Request){
	response.Write([]byte(lastUser))
}

func images (response http.ResponseWriter, request *http.Request) {
	var offers Offers
	offer1 := Offer{
		Image: "https://images2.bovpg.net/wsdv/media/1/1/5/7/2/157261.jpg",
		Title: "TENERIFE / COSTA ADEJE",
		Text:  "Escápate a Costa Adeje, un lugar idílico en la costa sur de Tenerife donde el descanso y la diversión están asegurados.",
		Price: "Desde 122€ / pers",
	}
	offer2 := Offer{
		Image: "https://images2.bovpg.net/wsdv/media/1/1/7/4/4/174443.jpg",
		Title: "ISLANDIA / REYKJAVIK",
		Text:  "Disfruta del máximo confort urbano en un acogedor hotel 4* que se inauguró en junio de 2020 en el centro de Reykjavik.",
		Price: "Desde 177€ / pers",
	}
	offer3 := Offer{
		Image: "https://images2.bovpg.net/wsdv/media/1/1/2/3/7/123728.jpg",
		Title: "ANDORRA / EL SERRAT",
		Text:  "Desconecta de la rutina en un entorno encantador entre montañas y aprovecha para practicar esquí en las estaciones del Principado.",
		Price: "Desde 103€ / pers",
	}
	offer4 := Offer{
		Image: "https://images2.bovpg.net/wsdv/media/1/1/1/2/7/112777.jpg",
		Title: "CAMBOYA / SIEM REAP",
		Text:  "Aprovecha la oportunidad de conocer un magnífico lugar: los templos de Angkor, únicos en el mundo.",
		Price: "Desde 242€ / pers",
	}
	offer5 := Offer{
		Image: "https://images2.bovpg.net/wsdv/media/1/1/5/9/4/159407.jpg",
		Title: "ESTADOS UNIDOS / NUEVA YORK",
		Text:  "Vive el vibrante espíritu de Nueva York y visita sus lugares más conocidos gracias a la excelente ubicación del hotel.",
		Price: "Desde 486€ / pers",
	}
	offer6 := Offer{
		Image: "https://images3.bovpg.net/wsdv/media/1/7/1/2/8/71289.jpg",
		Title: "MALDIVAS / RAA ATOLL",
		Text:  "Prepara las maletas en un hotel de 5* en las Maldivas, donde el lujo convive con un entorno natural y tropical único.",
		Price: "Desde 1759€ / pers",
	}

	offers.Offers = append(offers.Offers, offer1)
	offers.Offers = append(offers.Offers, offer2)
	offers.Offers = append(offers.Offers, offer3)
	offers.Offers = append(offers.Offers, offer4)
	offers.Offers = append(offers.Offers, offer5)
	offers.Offers = append(offers.Offers, offer6)

	data, _ := json.Marshal(offers.Offers)
	response.Write(data)
}