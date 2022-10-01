package main

import (
	"context"
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
