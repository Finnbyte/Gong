package ball

import (
	"github.com/gdamore/tcell/v2"
)

type Ball struct {
	Pos          [2]int
	HasHitPlayer bool
    Style tcell.Style
    Direction int
}

func (b *Ball) SwapDirection() {
    switch b.Direction {
    case 1:
        b.Direction = -1
    case -1:
        b.Direction = 1
    }
}

func (b *Ball) Move(scr tcell.Screen, moveSpeed int) {
    // Flush old entity from screen
    oldPos := b.Pos
    b.Clear(oldPos, scr)

    // Handle bouncing
    if b.shouldBounce(scr) {
        b.SwapDirection()
    }

    // Change position values
    b.Pos[0], b.Pos[1] = b.Pos[0] + b.Direction, b.Pos[1]

    // Draw changes to movement
    b.Draw(scr)

    // Update screen manually
    scr.Show()
}

func (b *Ball) Reset(x, y int, scr tcell.Screen) {
    b.Clear([2]int{b.Pos[0], b.Pos[1]}, scr)
    b.Pos[0], b.Pos[1] = x, y
    b.Draw(scr)
    scr.Show()
}

func (b *Ball) WentIn(winWidth int) bool {
    if b.Pos[0] < 2 || b.Pos[0] > winWidth - 2 {
        return true
    }
    return false 
}

func (b *Ball) shouldBounce(scr tcell.Screen) bool {
    rightSide, _, _, _ := scr.GetContent(b.Pos[0] + 1, b.Pos[1])
    leftSide, _, _, _ := scr.GetContent(b.Pos[0] - 1, b.Pos[1])
    if rightSide == tcell.RuneVLine || leftSide == tcell.RuneVLine {
       return true 
    }

    return false
}

func (b Ball) Draw(scr tcell.Screen) {
    scr.SetContent(b.Pos[0], b.Pos[1], 'O', nil, b.Style)
}

func (b Ball) Clear(pos [2]int, scr tcell.Screen) {
    scr.SetContent(pos[0], pos[1], ' ', nil, tcell.StyleDefault)
}
