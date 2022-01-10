package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChesseCake struct{}

// AlwaysConsumable ...
func (ChesseCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChesseCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChesseCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(17, 10.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChesseCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChesseCake) Edible() bool {
	return true
}

// EncodeItem ...
func (ChesseCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chesse_cake", 0
}

// Name ...
func (ChesseCake) Name() string {
	return "Chesse Cake"
}

// Texture ...
func (ChesseCake) Texture() image.Image {
	return textureFromName("white_cheese")
}
