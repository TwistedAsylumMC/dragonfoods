package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BagelAndLox struct{}

// AlwaysConsumable ...
func (BagelAndLox) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BagelAndLox) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BagelAndLox) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BagelAndLox) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BagelAndLox) Edible() bool {
	return true
}

// EncodeItem ...
func (BagelAndLox) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bagel_and_lox", 0
}

// Name ...
func (BagelAndLox) Name() string {
	return "BagelAndLox"
}

// Texture ...
func (BagelAndLox) Texture() image.Image {
	return textureFromName("bagel2")
}
