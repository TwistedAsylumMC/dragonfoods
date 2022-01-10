package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BeefJerky struct{}

// AlwaysConsumable ...
func (BeefJerky) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BeefJerky) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BeefJerky) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BeefJerky) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BeefJerky) Edible() bool {
	return true
}

// EncodeItem ...
func (BeefJerky) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:beef_jerky", 0
}

// Name ...
func (BeefJerky) Name() string {
	return "Beef Jerky"
}

// Texture ...
func (BeefJerky) Texture() image.Image {
	return textureFromName("beefjerkey")
}
