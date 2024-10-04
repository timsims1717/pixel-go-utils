package typeface

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/pixel-go-utils/object"
)

var (
	theSymbols = map[string]symbol{}
)

type symbol struct {
	spr *pixel.Sprite
	sca float64
}

type symbolHandle struct {
	symbol symbol
	trans  *object.Object
}

func RegisterSymbol(key string, spr *pixel.Sprite, scalar float64) {
	theSymbols[key] = symbol{
		spr: spr,
		sca: scalar,
	}
}
