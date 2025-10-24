package main

import (
	"fmt"
	"math/rand"
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

// Eat implements Animal.
func (c Capybara) Eat() string {
	return fmt.Sprintf("%v ate a watermelon and very much enjoyed it.", c)
}

// Move implements Animal.
func (c Capybara) Move() string {
	return fmt.Sprintf("%v sat and did nothing until %v jumped into the water and swam away very quickly.", c, c.SubjectPronoun())
}

// Sleep implements Animal.
func (c Capybara) Sleep() string {
	return fmt.Sprintf("%v was too lazy to go to bed and fell asleep where %v was sitting.", c, c.SubjectPronoun())
}

// WakeUp implements Animal.
func (c Capybara) WakeUp() string {
	return fmt.Sprintf("%v woke up and %v was very groggy", c, c.SubjectPronoun())
}

type Gopher struct {
	Gender
	name          string
	isUnderGround bool
}

// String implements Animal.
func (g Gopher) String() string { return "Gopher " + g.name }

// Eat implements Animal.
func (g Gopher) Eat() string {
	return fmt.Sprintf("%v digs up a potato and %v eats it!", g, g.SubjectPronoun())
}

// Move implements Animal.
func (g Gopher) Move() string {
	var description string
	if g.isUnderGround {
		// Moving starting under ground
		if rand.Intn(2) == 0 {
			// Moving above ground
			description = fmt.Sprintf(
				"%v pokes %v head out of %v tunnel and runs behind a bush.",
				g,
				g.Gender.PossessiveAdjective(),
				g.Gender.PossessiveAdjective(),
			)
			g.isUnderGround = false
		} else {
			// Staying under ground
			description = fmt.Sprintf(
				"%v crawls through the tunnels in %v hole home.",
				g,
				g.Gender.PossessiveAdjective(),
			)
		}
	} else {
		// Moving starting above ground
		if rand.Intn(2) == 0 {
			// Staying above ground
			description = fmt.Sprintf(
				"%v seems to be lost and can't find %v hole.",
				g,
				g.Gender.PossessiveAdjective(),
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

// Sleep implements Animal.
func (g Gopher) Sleep() string {
	return fmt.Sprintf("%v cozies up in %v tunnel home and falls fast asleep.", g, g.PossessiveAdjective())
}

// WakeUp implements Animal.
func (g Gopher) WakeUp() string {
	return fmt.Sprintf(
		"%v is woken up by footsteps above %v tunnel.",
		g,
		g.Gender.PossessiveAdjective(),
	)
}

type Cat struct {
	Gender
	name string
}

// String implements Animal.
func (c Cat) String() string {
}

// Eat implements Animal.
func (c Cat) Eat() string {
}

// Move implements Animal.
func (c Cat) Move() string {
}

// Sleep implements Animal.
func (c Cat) Sleep() string {
}

// WakeUp implements Animal.
func (c Cat) WakeUp() string {
}

type BushBaby struct {
	Gender
	name string
}

// String implements Animal.
func (b BushBaby) String() string { return "Bush baby " + b.name }

// Eat implements Animal.
func (b BushBaby) Eat() string {
}

// Move implements Animal.
func (b BushBaby) Move() string {
}

// Sleep implements Animal.
func (b BushBaby) Sleep() string {
}

// WakeUp implements Animal.
func (b BushBaby) WakeUp() string {
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
	var sleepDescriptions []string = make([]string, len(animals))
	for _, animal := range animals {
		sleepDescriptions = append(sleepDescriptions, animal.Sleep())
	}
	return sleepDescriptions
}
func (sanctuary AnimalSanctuary) alarmClock() []string {
	animals := sanctuary.getAllAnimals()
	var wakingUpDescriptions []string = make([]string, len(animals))
	for _, animal := range animals {
		wakingUpDescriptions = append(wakingUpDescriptions, animal.WakeUp())
	}
	return wakingUpDescriptions
}
func (sanctuary AnimalSanctuary) getAllAnimals() []Animal {
	var animals []Animal = make(
		[]Animal,
		len(sanctuary.capybaras)+len(sanctuary.gophers)+len(sanctuary.cats)+len(sanctuary.bushBabies),
	)
	for _, capybara := range sanctuary.capybaras {
		animals = append(animals, capybara)
	}
	for _, gopher := range sanctuary.gophers {
		animals = append(animals, gopher)
	}
	for _, cat := range sanctuary.cats {
		animals = append(animals, cat)
	}
	for _, bushBaby := range sanctuary.bushBabies {
		animals = append(animals, bushBaby)
	}
	return animals
}

func main() {
	const (
		sunset  = 0
		sunrise = 8
	)
	sanctuary := AnimalSanctuary{}
	for range 3 {
		for hour := range 24 {
			switch {
			case hour == sunset:
				sleepDescriptions := sanctuary.bedTime()
				for _, description := range sleepDescriptions {
					fmt.Println(description)
					time.Sleep(time.Millisecond * 100)
				}
			case hour == sunrise:
				wakeUpDescription := sanctuary.alarmClock()
				for _, description := range wakeUpDescription {
					fmt.Println(description)
					time.Sleep(time.Millisecond * 100)
				}
			case hour > sunset && hour < sunrise:
			default:
				randomAnimal := sanctuary.pickRandomAnimal()
				if rand.Intn(2) == 1 {
					fmt.Println(randomAnimal.Eat())
				} else {
					fmt.Println(randomAnimal.Move())
				}
			}
			time.Sleep(1)
		}
	}
}
