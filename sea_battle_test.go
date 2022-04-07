package main

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestGame(t *testing.T) {
	ctrl := gomock.NewController(t)
	p1 := NewMockPlayer(ctrl)
	p2 := NewMockPlayer(ctrl)
	g := NewGame(p1, p2, p1)

	p1.EXPECT().DoMove(0, 0).Return(MISS, nil)

	p2.EXPECT().DoMove(0, 0).Return(MISS, nil)

	g.HandleShoot("a0")

	g.HandleShoot("a0")

	//g.HandleShoot("b9")

	//g.HandleShoot("a0")

}

func TestSwitchPlayer(t *testing.T) { //проверка на сменяемость игрока

}
