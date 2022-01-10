package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CinnamonBread struct{}

// AlwaysConsumable ...
func (CinnamonBread) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CinnamonBread) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CinnamonBread) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (CinnamonBread) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CinnamonBread) Edible() bool {
	return true
}

// EncodeItem ...
func (CinnamonBread) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cinnamon_bread", 0
}

// Name ...
func (CinnamonBread) Name() string {
	return "CinnamonBread"
}

// Texture ...
func (CinnamonBread) Texture() image.Image {
	return textureFromName("loaf5")
}
