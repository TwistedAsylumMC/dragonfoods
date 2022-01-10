package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateEclair struct{}

// AlwaysConsumable ...
func (ChoclateEclair) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateEclair) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateEclair) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateEclair) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateEclair) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateEclair) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_eclair", 0
}

// Name ...
func (ChoclateEclair) Name() string {
	return "ChoclateEclair"
}

// Texture ...
func (ChoclateEclair) Texture() image.Image {
	return textureFromName("choclateeclair")
}
