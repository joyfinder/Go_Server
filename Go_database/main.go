package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}
	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}
}

func (d *Driver) Write() error {

}

func (d *Driver) Read() error {

}

func (d *Driver) ReadAll() {

}

func (d *Driver) Delete() error {

}

func (d *Driver) getOrCreateMutex() *sync.Mutex {

}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}
type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "23", "2334333", "Real", Address{"indo", "asd", "thai", "120381"}},
		{"Alice", "23", "2334333", "Fake comp", Address{"indo", "asd", "thai", "112804"}},
		{"Annie", "23", "2334333", "MSFT", Address{"indo", "asd", "thai", "1201924"}},
		{"Charlie", "23", "2334333", "Meta", Address{"indo", "asd", "thai", "1212412"}},
		{"Steph", "23", "2334333", "Google", Address{"indo", "asd", "thai", "15235524"}},
		{"Even", "23", "2334333", "Unicorn", Address{"indo", "asd", "thai", "12124"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Contact: value.Contact,
			Company: value.Company,
			Age:     value.Age,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	all_users := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		all_users = append(all_users, employeeFound)
	}
	fmt.Println((all_users))

	// db.Delete("user", "John"); err != nil{
	// 	fmt.Println("Error",err)
	// }
}
