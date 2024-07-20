package reflect_spell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name                  string
	charactersBeforeSpell []interface{}
	charactersAfterSpell  []interface{}
}

func TestCastToAll(t *testing.T) {
	for _, tc := range []testCase{
		{
			name: "one spell - fire",
			charactersBeforeSpell: []interface{}{
				&Zombie{Health: 1000},
				&Zombie{Health: 1000},
				&Orc{Health: 500},
				&Orc{Health: 500},
				&Orc{Health: 500},
				&Daemon{Health: 1000},
				&Daemon{Health: 1000},
			},
			charactersAfterSpell: []interface{}{
				&Zombie{Health: 950},
				&Zombie{Health: 950},
				&Orc{Health: 450},
				&Orc{Health: 450},
				&Orc{Health: 450},
				&Daemon{Health: 950},
				&Daemon{Health: 950},
			},
		},
		{
			name: "two spells - fire for splash health damage, protect to increase durability",
			charactersBeforeSpell: []interface{}{
				&Zombie{Health: 1000},
				&Zombie{Health: 1000},
				&Orc{Health: 500},
				&Orc{Health: 500},
				&Orc{Health: 500},
				&Daemon{Health: 1000},
				&Daemon{Health: 1000},
				&Wall{Durability: 100},
			},
			charactersAfterSpell: []interface{}{
				&Zombie{Health: 950},
				&Zombie{Health: 950},
				&Orc{Health: 450},
				&Orc{Health: 450},
				&Orc{Health: 450},
				&Daemon{Health: 950},
				&Daemon{Health: 950},
				&Wall{Durability: 250},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			CastToAll(newSpell("fire", "Health", -50), tc.charactersBeforeSpell)
			CastToAll(newSpell("protect", "Durability", 150), tc.charactersBeforeSpell)

			assert.Equal(t, tc.charactersBeforeSpell, tc.charactersAfterSpell)
		})
	}
}
