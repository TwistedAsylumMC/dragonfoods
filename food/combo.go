package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Combo struct{}

// AlwaysConsumable ...
func (Combo) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Combo) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Combo) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Combo) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Combo) Edible() bool {
	return true
}

// EncodeItem ...
func (Combo) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:combo", 0
}

// Name ...
func (Combo) Name() string {
	return "Combo"
}

// Texture ...
func (Combo) Texture() image.Image {
	return textureFromName("combo")
}
