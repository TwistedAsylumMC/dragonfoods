package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SweetCookies struct{}

// AlwaysConsumable ...
func (SweetCookies) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SweetCookies) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SweetCookies) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (SweetCookies) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SweetCookies) Edible() bool {
	return true
}

// EncodeItem ...
func (SweetCookies) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sweet_cookies", 0
}

// Name ...
func (SweetCookies) Name() string {
	return "Sweet Cookies"
}

// Texture ...
func (SweetCookies) Texture() image.Image {
	return textureFromName("sweetcookie")
}
