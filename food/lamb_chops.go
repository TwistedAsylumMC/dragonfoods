package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type LambChops struct{}

// AlwaysConsumable ...
func (LambChops) AlwaysConsumable() bool {
	return false
}

// Category ...
func (LambChops) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (LambChops) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(54, 32.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (LambChops) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (LambChops) Edible() bool {
	return true
}

// EncodeItem ...
func (LambChops) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:lamb_chops", 0
}

// Name ...
func (LambChops) Name() string {
	return "Lamb Chops"
}

// Texture ...
func (LambChops) Texture() image.Image {
	return textureFromName("lambchops")
}
