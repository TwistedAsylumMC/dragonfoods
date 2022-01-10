package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type IcecreamSandwich struct{}

// AlwaysConsumable ...
func (IcecreamSandwich) AlwaysConsumable() bool {
	return false
}

// Category ...
func (IcecreamSandwich) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (IcecreamSandwich) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (IcecreamSandwich) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (IcecreamSandwich) Edible() bool {
	return true
}

// EncodeItem ...
func (IcecreamSandwich) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:icecream_sandwich", 0
}

// Name ...
func (IcecreamSandwich) Name() string {
	return "Icecream Sandwich"
}

// Texture ...
func (IcecreamSandwich) Texture() image.Image {
	return textureFromName("icecreamsandwich")
}
