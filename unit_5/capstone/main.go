package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Mover interface {
	Move() string
}

type Eater interface {
	Eat() string
}

type Sleeper interface {
	Sleep() string
	WakeUp() string
}
type Gender struct {
	male bool
}

// SubjectPronoun can be he/she
func (g Gender) SubjectPronoun() string {
	if g.male {
		return "he"
	}
	return "she"
}

// ObjectPronoun can be him/her
func (g Gender) ObjectPronoun() string {
	if g.male {
		return "him"
	}
	return "her"
}

// PossessivePronoun can be his/hers
func (g Gender) PossessivePronoun() string {
	if g.male {
		return "his"
	}
	return "hers"
}

// PossessiveAdjective can be his/her
func (g Gender) PossessiveAdjective() string {
	if g.male {
		return "his"
	}
	return "her"
}

type Animal interface {
	Mover
	Eater
	Sleeper
	fmt.Stringer
}

type Capybara struct {
	Gender
	name string
}

func (c Capybara) String() string { return "Capybara " + c.name }

func (c *Capybara) Eat() string {
	return fmt.Sprintf("%v ate a watermelon and very much enjoyed it.", c)
}
func (c *Capybara) Move() string {
	return fmt.Sprintf("%v sat and did nothing until %v jumped into the water and swam away very quickly.", c, c.SubjectPronoun())
}
func (c *Capybara) Sleep() string {
	return fmt.Sprintf("%v was too lazy to go to bed and fell asleep where %v was sitting.", c, c.SubjectPronoun())
}
func (c *Capybara) WakeUp() string {
	return fmt.Sprintf("%v woke up and %v was very groggy.", c, c.SubjectPronoun())
}

type Gopher struct {
	Gender
	name          string
	isUnderGround bool
}

func (g Gopher) String() string { return "Gopher " + g.name }

func (g *Gopher) Eat() string {
	return fmt.Sprintf("%v digs up a potato and %v eats it!", g, g.SubjectPronoun())
}

func (g *Gopher) Move() string {
	var description string
	if g.isUnderGround {
		// Moving starting under ground
		if rand.Intn(2) == 0 {
			// Moving above ground
			description = fmt.Sprintf(
				"%v pokes %v head out of %v tunnel and runs behind a bush.",
				g,
				g.PossessiveAdjective(),
				g.PossessiveAdjective(),
			)
			g.isUnderGround = false
		} else {
			// Staying under ground
			description = fmt.Sprintf(
				"%v crawls through the tunnels in %v hole home.",
				g,
				g.PossessiveAdjective(),
			)
		}
	} else {
		// Moving starting above ground
		if rand.Intn(2) == 0 {
			// Staying above ground
			description = fmt.Sprintf(
				"%v seems to be lost and can't find %v hole.",
				g,
				g.PossessiveAdjective(),
			)
		} else {
			// Moving under ground
			description = fmt.Sprintf(
				"%v is scared and scurries into %v hole to hide.",
				g,
				g.PossessiveAdjective(),
			)
			g.isUnderGround = true
		}
	}
	return description
}

func (g *Gopher) Sleep() string {
	if g.isUnderGround {
		return fmt.Sprintf("%v cozies up in %v tunnel home and falls fast asleep.", g, g.PossessiveAdjective())
	}
	return fmt.Sprintf("%v can't find %v tunnel home and sleeps in a bush.", g, g.PossessiveAdjective())
}
func (g *Gopher) WakeUp() string {
	return fmt.Sprintf("%v is woken up by footsteps above %v tunnel.", g, g.PossessiveAdjective())
}

type Cat struct {
	Gender
	name string
}

func (c Cat) String() string { return "Cat " + c.name }

func (c *Cat) Eat() string {
	switch rand.Intn(3) {
	case 0: // Fish
		return fmt.Sprintf("%v spotted a small fish swimming close to the surface of the water and nabbed it!", c)
	case 1: // Milk
		return fmt.Sprintf("Someone left out a warm bowl of milk and %v has found it.", c)
	default: // Mouse
		return fmt.Sprintf("%v spotted a mouse hiding behind the sofa and pounced on it and ate it.", c)
	}
}
func (c *Cat) Move() string {
	return fmt.Sprintf("%v slinks around the perimeter, pausing to make bird chirps as %v goes.", c, c.SubjectPronoun())
}
func (c *Cat) Sleep() string {
	return fmt.Sprintf("%v has curled up into a fluffy ball for the night.", c)
}
func (c *Cat) WakeUp() string {
	return fmt.Sprintf("%v lets out a big yawn and does a big stretch and is ready for the day.", c)
}

type BushBaby struct {
	Gender
	name string
}

func (b BushBaby) String() string { return "Bush baby " + b.name }

func (b *BushBaby) Eat() string {
	return fmt.Sprintf("%v found a juicy bug and %v ate it.", b, b.SubjectPronoun())
}

func (b *BushBaby) Move() string {
	return fmt.Sprintf("%v jumps from tree to tree.", b)
}

func (b *BushBaby) Sleep() string {
	return fmt.Sprintf("%v cuddles up with the other bush babies in their nest and falls asleep.", b)
}
func (b *BushBaby) WakeUp() string {
	return fmt.Sprintf("%v wakes up with the sunrise as soon as it hits the canopy of leaves.", b)
}

type AnimalSanctuary struct {
	capybaras  []Capybara
	gophers    []Gopher
	cats       []Cat
	bushBabies []BushBaby
}

func (sanctuary AnimalSanctuary) pickRandomAnimal() Animal {
	animals := sanctuary.getAllAnimals()
	animalIndex := rand.Intn(len(animals))
	return animals[animalIndex]
}

func (sanctuary AnimalSanctuary) bedTime() []string {
	animals := sanctuary.getAllAnimals()
	desc := make([]string, 0, len(animals))
	for _, a := range animals {
		desc = append(desc, a.Sleep())
	}
	return desc
}

func (sanctuary AnimalSanctuary) alarmClock() []string {
	animals := sanctuary.getAllAnimals()
	desc := make([]string, 0, len(animals))
	for _, a := range animals {
		desc = append(desc, a.WakeUp())
	}
	return desc
}

func (sanctuary AnimalSanctuary) getAllAnimals() []Animal {
	animals := make([]Animal, 0,
		len(sanctuary.capybaras)+len(sanctuary.gophers)+len(sanctuary.cats)+len(sanctuary.bushBabies),
	)

	// IMPORTANT: take addresses of slice elements by index (not the range variable)
	for i := range sanctuary.capybaras {
		animals = append(animals, &sanctuary.capybaras[i])
	}
	for i := range sanctuary.gophers {
		animals = append(animals, &sanctuary.gophers[i])
	}
	for i := range sanctuary.cats {
		animals = append(animals, &sanctuary.cats[i])
	}
	for i := range sanctuary.bushBabies {
		animals = append(animals, &sanctuary.bushBabies[i])
	}
	return animals
}

func main() {
	const (
		sunset  = 0
		sunrise = 8
	)
	male := Gender{true}
	female := Gender{false}
	sanctuary := AnimalSanctuary{
		capybaras: []Capybara{
			{Gender: male, name: "Pickles"},
			{Gender: female, name: "Dumpling"},
		},
		gophers: []Gopher{
			{Gender: female, name: "Daisy", isUnderGround: true},
		},
		cats: []Cat{
			{Gender: male, name: "Milo"},
			{Gender: female, name: "Mia"},
		},
		bushBabies: []BushBaby{
			{Gender: male, name: "Pip"},
			{Gender: female, name: "Dottie"},
		},
	}

	for range 3 {
		for hour := range 24 {
			switch {
			case hour == sunset:
				fmt.Println("Bedtime!")
				for _, d := range sanctuary.bedTime() {
					fmt.Println(d)
					time.Sleep(500 * time.Millisecond)
				}
				fmt.Println("Everyone is asleep.")
			case hour == sunrise:
				fmt.Println("Sunrise!")
				for _, d := range sanctuary.alarmClock() {
					fmt.Println(d)
					time.Sleep(500 * time.Millisecond)
				}
			case hour > sunset && hour < sunrise:
				fmt.Println(strings.Repeat(".", hour))
			default:
				a := sanctuary.pickRandomAnimal()
				if rand.Intn(2) == 1 {
					fmt.Println(a.Eat())
				} else {
					fmt.Println(a.Move())
				}
			}
			time.Sleep(2 * time.Second)
		}
	}
}
