package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type LemonChicken struct{}

// AlwaysConsumable ...
func (LemonChicken) AlwaysConsumable() bool {
	return false
}

// Category ...
func (LemonChicken) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (LemonChicken) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (LemonChicken) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (LemonChicken) Edible() bool {
	return true
}

// EncodeItem ...
func (LemonChicken) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:lemon_chicken", 0
}

// Name ...
func (LemonChicken) Name() string {
	return "Lemon Chicken"
}

// Texture ...
func (LemonChicken) Texture() image.Image {
	return textureFromName("lemonchicken")
}
