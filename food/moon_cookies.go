package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MoonCookies struct{}

// AlwaysConsumable ...
func (MoonCookies) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MoonCookies) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MoonCookies) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (MoonCookies) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MoonCookies) Edible() bool {
	return true
}

// EncodeItem ...
func (MoonCookies) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:moon_cookies", 0
}

// Name ...
func (MoonCookies) Name() string {
	return "Moon Cookies"
}

// Texture ...
func (MoonCookies) Texture() image.Image {
	return textureFromName("cookie_png")
}
