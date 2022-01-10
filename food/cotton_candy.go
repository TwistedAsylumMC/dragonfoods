package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CottonCandy struct{}

// AlwaysConsumable ...
func (CottonCandy) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CottonCandy) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CottonCandy) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CottonCandy) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CottonCandy) Edible() bool {
	return true
}

// EncodeItem ...
func (CottonCandy) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cotton_candy", 0
}

// Name ...
func (CottonCandy) Name() string {
	return "Cotton Candy"
}

// Texture ...
func (CottonCandy) Texture() image.Image {
	return textureFromName("cotton_candy")
}
