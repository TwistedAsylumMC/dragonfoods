package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateCookies struct{}

// AlwaysConsumable ...
func (ChoclateCookies) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateCookies) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateCookies) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateCookies) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateCookies) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateCookies) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_cookies", 0
}

// Name ...
func (ChoclateCookies) Name() string {
	return "ChoclateCookies"
}

// Texture ...
func (ChoclateCookies) Texture() image.Image {
	return textureFromName("cookie_1")
}
