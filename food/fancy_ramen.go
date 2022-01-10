package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FancyRamen struct{}

// AlwaysConsumable ...
func (FancyRamen) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FancyRamen) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FancyRamen) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (FancyRamen) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FancyRamen) Edible() bool {
	return true
}

// EncodeItem ...
func (FancyRamen) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fancy_ramen", 0
}

// Name ...
func (FancyRamen) Name() string {
	return "Fancy Ramen"
}

// Texture ...
func (FancyRamen) Texture() image.Image {
	return textureFromName("ramen")
}
