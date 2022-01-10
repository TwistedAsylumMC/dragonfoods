package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Bannana struct{}

// AlwaysConsumable ...
func (Bannana) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Bannana) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Bannana) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Bannana) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Bannana) Edible() bool {
	return true
}

// EncodeItem ...
func (Bannana) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bannana", 0
}

// Name ...
func (Bannana) Name() string {
	return "Bannana"
}

// Texture ...
func (Bannana) Texture() image.Image {
	return textureFromName("banana")
}
