package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CandyCane struct{}

// AlwaysConsumable ...
func (CandyCane) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CandyCane) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CandyCane) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (CandyCane) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CandyCane) Edible() bool {
	return true
}

// EncodeItem ...
func (CandyCane) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:candy_cane", 0
}

// Name ...
func (CandyCane) Name() string {
	return "CandyCane"
}

// Texture ...
func (CandyCane) Texture() image.Image {
	return textureFromName("candy_cane")
}
