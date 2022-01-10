package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChessePizza struct{}

// AlwaysConsumable ...
func (ChessePizza) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChessePizza) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChessePizza) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChessePizza) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChessePizza) Edible() bool {
	return true
}

// EncodeItem ...
func (ChessePizza) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chesse_pizza", 0
}

// Name ...
func (ChessePizza) Name() string {
	return "ChessePizza"
}

// Texture ...
func (ChessePizza) Texture() image.Image {
	return textureFromName("pizza3")
}
