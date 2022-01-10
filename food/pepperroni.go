package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pepperroni struct{}

// AlwaysConsumable ...
func (Pepperroni) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pepperroni) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pepperroni) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pepperroni) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pepperroni) Edible() bool {
	return true
}

// EncodeItem ...
func (Pepperroni) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pepperroni", 0
}

// Name ...
func (Pepperroni) Name() string {
	return "Pepperroni"
}

// Texture ...
func (Pepperroni) Texture() image.Image {
	return textureFromName("pepperoni")
}
