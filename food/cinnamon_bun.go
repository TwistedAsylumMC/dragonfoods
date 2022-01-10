package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CinnamonBun struct{}

// AlwaysConsumable ...
func (CinnamonBun) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CinnamonBun) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CinnamonBun) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (CinnamonBun) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CinnamonBun) Edible() bool {
	return true
}

// EncodeItem ...
func (CinnamonBun) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cinnamon_bun", 0
}

// Name ...
func (CinnamonBun) Name() string {
	return "Cinnamon Bun"
}

// Texture ...
func (CinnamonBun) Texture() image.Image {
	return textureFromName("baked_cinnamon_bun")
}
