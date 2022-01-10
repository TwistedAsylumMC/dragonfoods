package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChesseeBruger struct{}

// AlwaysConsumable ...
func (ChesseeBruger) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChesseeBruger) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChesseeBruger) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChesseeBruger) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChesseeBruger) Edible() bool {
	return true
}

// EncodeItem ...
func (ChesseeBruger) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chessee_bruger", 0
}

// Name ...
func (ChesseeBruger) Name() string {
	return "ChesseeBruger"
}

// Texture ...
func (ChesseeBruger) Texture() image.Image {
	return textureFromName("16burger")
}
