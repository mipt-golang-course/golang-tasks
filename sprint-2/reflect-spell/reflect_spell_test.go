package reflectspell_test

import (
	"testing"

	spells "github.com/mipt-golang-course/golang-tasks/sprint-2/reflect-spell"
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
				&spells.Zombie{Health: 1000},
				&spells.Zombie{Health: 1000},
				&spells.Orc{Health: 500},
				&spells.Orc{Health: 500},
				&spells.Orc{Health: 500},
				&spells.Daemon{Health: 1000},
				&spells.Daemon{Health: 1000},
			},
			charactersAfterSpell: []interface{}{
				&spells.Zombie{Health: 950},
				&spells.Zombie{Health: 950},
				&spells.Orc{Health: 450},
				&spells.Orc{Health: 450},
				&spells.Orc{Health: 450},
				&spells.Daemon{Health: 950},
				&spells.Daemon{Health: 950},
			},
		},
		{
			name: "two spells - fire for splash health damage, protect to increase durability",
			charactersBeforeSpell: []interface{}{
				&spells.Zombie{Health: 1000},
				&spells.Zombie{Health: 1000},
				&spells.Orc{Health: 500},
				&spells.Orc{Health: 500},
				&spells.Orc{Health: 500},
				&spells.Daemon{Health: 1000},
				&spells.Daemon{Health: 1000},
				&spells.Wall{Durability: 100},
			},
			charactersAfterSpell: []interface{}{
				&spells.Zombie{Health: 950},
				&spells.Zombie{Health: 950},
				&spells.Orc{Health: 450},
				&spells.Orc{Health: 450},
				&spells.Orc{Health: 450},
				&spells.Daemon{Health: 950},
				&spells.Daemon{Health: 950},
				&spells.Wall{Durability: 250},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			spells.CastToAll(spells.NewSpell("fire", "Health", -50), tc.charactersBeforeSpell)
			spells.CastToAll(spells.NewSpell("protect", "Durability", 150), tc.charactersBeforeSpell)

			assert.Equal(t, tc.charactersBeforeSpell, tc.charactersAfterSpell)
		})
	}
}
