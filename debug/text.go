package debug

import (
	"fmt"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"github.com/timsims1717/pixel-go-utils/typeface"
	"strings"
)

var (
	debugText *typeface.Text
	lines     []string
)

func InitializeText() {
	debugText = typeface.New("basic", typeface.NewAlign(typeface.Left, typeface.Top), 1.0, 2.0, 0., 0.)
}

func DrawText(win *pixelgl.Window) {
	if ShowText {
		var sb strings.Builder
		for i, line := range lines {
			if i != 0 {
				sb.WriteString("\n")
			}
			sb.WriteString(line)
		}
		debugText.SetText(sb.String())
		debugText.Obj.Pos = winV.Add(pixel.V(win.Bounds().W()*-0.5+2., win.Bounds().H()*0.5-2))
		debugText.Obj.Update()
		debugText.Draw(win)
	}
}

func AddText(s string) {
	lines = append(lines, s)
}

func AddIntCoords(label string, x, y int) {
	lines = append(lines, fmt.Sprintf("%s: (%d,%d)", label, x, y))
}

func AddTruthText(label string, b bool) {
	lines = append(lines, fmt.Sprintf("%s: %t", label, b))
}

func InsertText(s string, i int) {
	if i < 0 || len(lines) <= i || len(lines) == 0 {
		AddText(s)
	} else {
		tmp := append(lines[:i], s)
		tmp = append(tmp, lines[i:]...)
		lines = tmp
	}
}

func SetTextColor(col pixel.RGBA) {
	fpsText.SetColor(col)
	versionText.SetColor(col)
	debugText.SetColor(col)
}
