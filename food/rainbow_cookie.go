package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RainbowCookie struct{}

// AlwaysConsumable ...
func (RainbowCookie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RainbowCookie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RainbowCookie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (RainbowCookie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RainbowCookie) Edible() bool {
	return true
}

// EncodeItem ...
func (RainbowCookie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:rainbow_cookie", 0
}

// Name ...
func (RainbowCookie) Name() string {
	return "Rainbow Cookie"
}

// Texture ...
func (RainbowCookie) Texture() image.Image {
	return textureFromName("rainbow_cookie")
}
