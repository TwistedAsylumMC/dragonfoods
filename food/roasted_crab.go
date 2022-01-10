package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RoastedCrab struct{}

// AlwaysConsumable ...
func (RoastedCrab) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RoastedCrab) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RoastedCrab) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (RoastedCrab) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RoastedCrab) Edible() bool {
	return true
}

// EncodeItem ...
func (RoastedCrab) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:roasted_crab", 0
}

// Name ...
func (RoastedCrab) Name() string {
	return "Roasted Crab"
}

// Texture ...
func (RoastedCrab) Texture() image.Image {
	return textureFromName("crab2")
}
