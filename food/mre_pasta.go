package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MrePasta struct{}

// AlwaysConsumable ...
func (MrePasta) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MrePasta) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MrePasta) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (MrePasta) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MrePasta) Edible() bool {
	return true
}

// EncodeItem ...
func (MrePasta) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mre_pasta", 0
}

// Name ...
func (MrePasta) Name() string {
	return "MrePasta"
}

// Texture ...
func (MrePasta) Texture() image.Image {
	return textureFromName("mre")
}
