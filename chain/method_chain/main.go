package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{
		Name:    name,
		Attack:  attack,
		Defense: defense,
	}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature, nil}
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{
		CreatureModifier: CreatureModifier{
			c, nil,
		},
	}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Printf("Doubling %s \b's attack %d \n", d.creature.Name, d.creature.Attack*2)

	d.creature.Attack *= 2
	d.CreatureModifier.Handle()

}

type IncreasedDefenseModifier struct {
	CreatureModifier
}

func NewIncreasedDefenseModifier(
	c *Creature) *IncreasedDefenseModifier {
	return &IncreasedDefenseModifier{CreatureModifier{
		creature: c}}
}

func (i *IncreasedDefenseModifier) Handle() {
	if i.creature.Attack <= 2 {
		fmt.Printf("Increasing %s \b's defense %d \n", i.creature.Name, i.creature.Defense+1)
		i.creature.Defense++
	}
	i.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(
	c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{
		creature: c}}
}

func (n *NoBonusesModifier) Handle() {
	// nothing here!
}

func main() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)

	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreasedDefenseModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))

	root.Handle()
	fmt.Println(goblin.String())
}
