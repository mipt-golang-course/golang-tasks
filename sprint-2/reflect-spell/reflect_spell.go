package reflect_spell

import (
	"log"
	"reflect"
)

type Spell interface {
	// название заклинания
	Name() string
	// характеристика, на которую воздействует
	Char() string
	// количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу, то заклинание применяется через него
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object interface{}) {
	if player, ok := object.(CastReceiver); ok {
		player.ReceiveSpell(spell)
		return
	}

	// Player's not connected to interface. Apply spell with reflection
	val := reflect.ValueOf(object)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field := val.FieldByName(spell.Char())

	if field.IsValid() {
		if field.CanSet() {
			if field.Kind() == reflect.Int {
				field.SetInt(field.Int() + int64(spell.Value()))
			} else {
				log.Printf("Field %s in object %v has non-iteger type. Cannot apply", spell.Char(), object)
			}
		} else {
			log.Printf("Field %s cannot be set in object %v", spell.Char(), object)
		}
	} else {
		log.Printf("Field %s not found in object %v", spell.Char(), object)
	}
}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	// nolint: unused
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}
