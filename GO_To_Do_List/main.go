package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gofiber/fiber/middleware"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rdr *renderer.Render
var db *mgo.Database

const (
	hostname       string = "localhost:27001"
	dbName         string = "demo_todo"
	collectionName string = "todo"
	port           string = ":9000"
)

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
		CreatedAt time.Time     `bson:"createAt"`
	}
	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `bson:"completed"`
		CreatedAt time.Time `json:"created_at`
	}
)

func main() {

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Mount("/todo", todoHandlers())

	svr := &http.Server{
		Address:      port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func(){
		log.Println("Listening on port", port)
		if err := svr.ListenAndServe(); err != nil {
			log.Printf("listen:%s\n", err)
		}
	}


	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	svr.Shutdown(ctx)
	defer cancel(
		log.Println("Server stopped.")
	)
}
func homeHandler(w http.ResponseWriter, r *http.Request){
	err := rdr.Template(w, http.StatusOK, []string{"static/home.tpl", nil})
	checkErr(err)
}

func fetchTodos(w http.ResponseWriter, r *http.Request){
	todos := []todoModel{}

	if err := db.C(collectionName).Find(bson.M{}).All(&todos); err != nil {
		rdr.JSON(w, http.StatusProcessing, renderer.M{
			"message":"Failed to fetch todo",
			"error":err,
		})
		return
	}
	todoList := []todo{}

	// Append todo tasks into list
	for _ , t := range todos{
		todoList = append(todoList, todo{
			ID: t.ID.Hex(),
			Title: t.Title,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		})
	}
	rdr.JSON(w, http.StatusOK, renderer.M{
		"data": todoList,
	})
}

func createTodo(w http.ResponseWriter, r *http.Request){
	var t todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rdr.JSON(w, http.StateusProcess, err)
		return
	}

	if t.tile == ""{
		rdr.JSON(w, http.StatusBadRequest, renderer.M{
			"message":"Required an title input",
		})
		return
	}

	tm := todoModel{
		ID: bson.NewObjectId(),
		Title: t.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	if err := db.C(collectionName).Insert(&tm); err != nil {
		rdr.JSON(w, http.StatusProcessing, renderer.M{
			"message":"Failed to save todo memo"
			"error":err,
		})
		return
	}
	rdr.JSON(w, http.StatusCreated, renderer.M{
		"message":"todo created successfully"
		"todo_id": tm.ID.Hex()
	})
}

func deleteTodo(w http.ResponseWriter, r *http.Request){
	id := strings.TrimSpace(chi.URLParam(r, "id"))

	if !bson.IsObjectIdHex(id){
		rdr.JSON(w, http.StatusBadRequest, renderer.M{
			"message":"Id is invalid."
		})
		return

		if err := db.C(collectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
			rdr.JSON(w, http.StatusProcessing, renderer.M{
				"message":"Failed to delete todo memo",
				"error":err,
			})
			return
		}

		rdr.JSON(w, http.StatusOK, renderer.M{
			"message":"deleted todo successfully"
		})
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request){
	id := strings.TrimSpace(chi.URLParam(r, "id"))

	if !bson.IsObjectIdHex(id){
		rdr.JSON(w, http.StatusBadRequest, renderer.M{
			"message":"The id is invalid",
		})
		return 
	}

	var t todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rdr.JSON(w, http.StatusProcessing, err)
		return  
	}

	if t.Title = ""{
		rdr.JSON(w, http.StatusBadRequest, renderer.M{
			"message":"Title field id is required",
		})
		return 
	}
}

func homeHandler() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", fetchTodos)
		r.Post("/", createTodo)
		r.Put("/{id}", updateTodo)
		r.Delete("/{id}", deleteTodo)
	})
}

func init() {
	rdr = renderer.New()
	sess, err := mgo.Dial(hostname)
	checkErr(err)
	// Changing the setmode
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
