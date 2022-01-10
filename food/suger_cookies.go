package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SugerCookies struct{}

// AlwaysConsumable ...
func (SugerCookies) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SugerCookies) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SugerCookies) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (SugerCookies) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SugerCookies) Edible() bool {
	return true
}

// EncodeItem ...
func (SugerCookies) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:suger_cookies", 0
}

// Name ...
func (SugerCookies) Name() string {
	return "Suger Cookies"
}

// Texture ...
func (SugerCookies) Texture() image.Image {
	return textureFromName("sugercookie")
}
