package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PotatoChips struct{}

// AlwaysConsumable ...
func (PotatoChips) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PotatoChips) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PotatoChips) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (PotatoChips) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PotatoChips) Edible() bool {
	return true
}

// EncodeItem ...
func (PotatoChips) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:potato_chips", 0
}

// Name ...
func (PotatoChips) Name() string {
	return "PotatoChips"
}

// Texture ...
func (PotatoChips) Texture() image.Image {
	return textureFromName("potatochip_yellow")
}
