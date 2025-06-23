package reflect_spell

import (
	"fmt"
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

//================================================================================================================================================

// reflect.TypeOf(): возвращает reflect.Type, представляющий тип переменной.
//
// reflect.ValueOf(): возвращает reflect.Value, представляющий значение переменной.
//
// Interface(): возвращает interface{}, представляющий текущее значение. Это позволяет извлекать оригинальные данные из reflect.Value.
//
// Kind(): возвращает reflect.Kind, представляющий конкретный тип данных, например, Int, Float64, Struct, и т.д.
//
// NumField() и Field(i int): используются для работы со структурами, позволяют получить количество полей и доступ к каждому полю соответственно.
//
// NumMethod() и Method(i int): используются для работы с методами, предоставляя доступ к методу и его характеристикам.

////================================================================================================================================================

func CastTo(spell Spell, object interface{}) {
	if receiver, ok := object.(CastReceiver); ok {
		receiver.ReceiveSpell(spell)
		return
	}

	val := reflect.ValueOf(object)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		if val.Kind() != reflect.Ptr {
			fmt.Printf("object %T not ptr", object)
			return
		} else if val.IsNil() {
			fmt.Printf("object is a nil pointer of type %T", object)
		} else {
			fmt.Printf("object not struct")
		}

		return
	}

	val = val.Elem()

	for indexOfField := 0; indexOfField < val.NumField(); indexOfField++ {
		field := val.Field(indexOfField)
		fieldName := val.Type().Field(indexOfField).Name

		if fieldName == spell.Char() {
			if !field.CanSet() {
				return
			}

			field.SetInt(field.Int() + int64(spell.Value()))
		}
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
