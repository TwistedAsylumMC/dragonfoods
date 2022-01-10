package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ButterChicken struct{}

// AlwaysConsumable ...
func (ButterChicken) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ButterChicken) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ButterChicken) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (ButterChicken) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ButterChicken) Edible() bool {
	return true
}

// EncodeItem ...
func (ButterChicken) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:butter_chicken", 0
}

// Name ...
func (ButterChicken) Name() string {
	return "Butter Chicken"
}

// Texture ...
func (ButterChicken) Texture() image.Image {
	return textureFromName("butterchicken")
}
