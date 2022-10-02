package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl", nil})
	checkErr(err)
}

func fetchTodos(w http.ResponseWriter, r *http.Request){
	todos := []todoModel{}

	if err := db.C(collectionName).Find(bson.M{}).All(&todos); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
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
	rnd.JSON(w, http.StatusOK, renderer.M{
		"data": todoList,
	})
}

func createTodo(w http.ResponseWriter, r *http.Request){
	var t todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rdr.JSON(w, http.StateusProcess, err)
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
