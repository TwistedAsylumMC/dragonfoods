package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MacNCheese struct{}

// AlwaysConsumable ...
func (MacNCheese) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MacNCheese) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MacNCheese) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (MacNCheese) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MacNCheese) Edible() bool {
	return true
}

// EncodeItem ...
func (MacNCheese) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mac_n_cheese", 0
}

// Name ...
func (MacNCheese) Name() string {
	return "Mac N Cheese"
}

// Texture ...
func (MacNCheese) Texture() image.Image {
	return textureFromName("41_eggsalad_bowl")
}
