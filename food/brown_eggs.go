package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BrownEggs struct{}

// AlwaysConsumable ...
func (BrownEggs) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BrownEggs) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BrownEggs) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BrownEggs) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BrownEggs) Edible() bool {
	return true
}

// EncodeItem ...
func (BrownEggs) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:brown_eggs", 0
}

// Name ...
func (BrownEggs) Name() string {
	return "BrownEggs"
}

// Texture ...
func (BrownEggs) Texture() image.Image {
	return textureFromName("egg_white")
}
