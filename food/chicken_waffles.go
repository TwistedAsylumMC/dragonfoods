package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChickenWaffles struct{}

// AlwaysConsumable ...
func (ChickenWaffles) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChickenWaffles) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChickenWaffles) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChickenWaffles) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChickenWaffles) Edible() bool {
	return true
}

// EncodeItem ...
func (ChickenWaffles) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chicken_waffles", 0
}

// Name ...
func (ChickenWaffles) Name() string {
	return "ChickenWaffles"
}

// Texture ...
func (ChickenWaffles) Texture() image.Image {
	return textureFromName("102_waffle_dish")
}
