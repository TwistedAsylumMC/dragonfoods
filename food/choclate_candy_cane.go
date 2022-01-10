package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateCandyCane struct{}

// AlwaysConsumable ...
func (ChoclateCandyCane) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateCandyCane) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateCandyCane) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateCandyCane) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateCandyCane) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateCandyCane) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_candy_cane", 0
}

// Name ...
func (ChoclateCandyCane) Name() string {
	return "ChoclateCandyCane"
}

// Texture ...
func (ChoclateCandyCane) Texture() image.Image {
	return textureFromName("candycane")
}
