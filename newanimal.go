// User creates animals and get info about these animals
// animas have a name and can be either cow, bird, or snake
// user either creates new animal or requests information abput created animals

// newanimal commands must be a single line containing 3 strings - newanimal name Type
// Prints 'Created it!'

// query commands must be a single line containig 3 strings - query name requested-information
// requests are eat, move, Speak

// define an interface type Animal - contains EAT, Move, Speak methods

// define 3 types - cow bird snake, define their methods
// when user creates an animal, create corresponding object

package main

import (
  "fmt"
)

type Animal interface {
  Eat()
  Move()
  Speak()
}

type Cow struct {
  name string
  food string
  locomotion string
  sound string
}

type Bird struct {
  name string
  food string
  locomotion string
  sound string
}

type Snake struct {
  name string
  food string
  locomotion string
  sound string
}

func (a Cow) Eat() {
  fmt.Println(a.food, "\n")
}

func (a Bird) Eat() {
  fmt.Println(a.food, "\n")
}

func (a Snake) Eat() {
  fmt.Println(a.food, "\n")
}

func (a Cow) Move() {
  fmt.Println(a.locomotion, "\n")
}

func (a Bird) Move() {
  fmt.Println(a.locomotion, "\n")
}

func (a Snake) Move() {
  fmt.Println(a.locomotion, "\n")
}

func (a Cow) Speak() {
  fmt.Println(a.sound, "\n")
}

func (a Bird) Speak() {
  fmt.Println(a.sound, "\n")
}

func (a Snake) Speak() {
  fmt.Println(a.sound, "\n")
}

func newanimal(mapanimal map[string]Animal, name string, animal_type string){
  switch animal_type {
  case "cow":
    mapanimal[name] = Cow{food:"grass", locomotion:"walk", sound:"moo"}
    fmt.Println("Created it!")
  case "bird":
    mapanimal[name] = Bird{food:"worms", locomotion:"fly", sound:"peep"}
    fmt.Println("Created it!")
  case "snake":
    mapanimal[name] = Snake{food:"mice", locomotion:"slither", sound:"hss"}
    fmt.Println("Created it!")
  }
}

func info_query(animal Animal, info string) {
  switch info {
  case "eat":
    animal.Eat()
  case "move":
    animal.Move()
  case "speak":
    animal.Speak()
  default:
    fmt.Println("Not found.")
}
}
func main() {
  var mapAnimal map[string]Animal
  mapAnimal = make(map[string]Animal) 
  fmt.Println("To create a new animal follow instructions >newanimal name animal-type(cow,bird,snake)\n")
  fmt.Println("To query animals follow instructions >query name query-type(eat, move,speak)\n")
  var user_input_rqst, user_input_2, user_input_3 string
  for {
    fmt.Printf(">")
    fmt.Scan(&user_input_rqst)
    fmt.Scan(&user_input_2)
    fmt.Scan(&user_input_3)

    switch user_input_rqst {
    case "newanimal":
      newanimal(mapAnimal,user_input_2,user_input_3)
    case "query":
      var query_animal Animal 
      query_animal = mapAnimal[user_input_2]
      info_query(query_animal, user_input_3)
    }
  }
}
