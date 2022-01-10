package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Almond struct{}

// AlwaysConsumable ...
func (Almond) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Almond) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Almond) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Almond) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Almond) Edible() bool {
	return true
}

// EncodeItem ...
func (Almond) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:almond", 0
}

// Name ...
func (Almond) Name() string {
	return "Almond"
}

// Texture ...
func (Almond) Texture() image.Image {
	return textureFromName("almond")
}
