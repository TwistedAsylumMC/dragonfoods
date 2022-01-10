package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedCinnamonBun struct{}

// AlwaysConsumable ...
func (BakedCinnamonBun) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedCinnamonBun) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedCinnamonBun) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedCinnamonBun) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedCinnamonBun) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedCinnamonBun) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_cinnamon_bun", 0
}

// Name ...
func (BakedCinnamonBun) Name() string {
	return "BakedCinnamonBun"
}

// Texture ...
func (BakedCinnamonBun) Texture() image.Image {
	return textureFromName("cinnamon_bun")
}
