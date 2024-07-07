## MyCheck

Представьте, что вы разрабатываете многопользовательскую игру. Игровая логика обсчитывается на сервере, реализованном на Go. В какой-то момент от дизайнеров поступает запрос на внедрение для игроков массовых заклинаний. То есть заклинание должно действовать на множество объектов, имеющих разнообразные структуры и типы. Тут вы понимаете, что переписывать все типы и реализовывать для каждого интерфейс CastReceiver — слишком сложная задача.
Реализуйте применение заклинаний с помощью рефлексии. Каждое заклинание удовлетворяет интерфейсу Spell — можно узнать, на какую характеристику объекта и на какую величину оно влияет.
См. примеры в тестах.

#### Пример main.go для данной функции:

```
package main

import (
    "fmt"
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
    // реализуйте эту функцию.
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

func main() {

    player := &Player{
        name:   "Player_1",
        health: 100,
    }

    enemies := []interface{}{
        &Zombie{Health: 1000},
        &Zombie{Health: 1000},
        &Orc{Health: 500},
        &Orc{Health: 500},
        &Orc{Health: 500},
        &Daemon{Health: 1000},
        &Daemon{Health: 1000},
        &Wall{Durability: 100},
    }

    CastToAll(newSpell("fire", "Health", -50), append(enemies, player))
    CastToAll(newSpell("heal", "Health", 190), append(enemies, player))

    fmt.Println(player)
}
```

### Проверка решения

Для запуска тестов нужно выполнить следующую команду:

```
go test -v ./sprint-2/reflect-spell/...
```

### Линтер

Если вы ещё этого не сделали, установите [golangci-lint](https://github.com/golangci/golangci-lint) и проверьте решение перед отправкой!
```
golangci-lint -v run ./sprint-2/reflect-spell/...
```
