package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Grub struct{}

// AlwaysConsumable ...
func (Grub) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Grub) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Grub) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Grub) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Grub) Edible() bool {
	return true
}

// EncodeItem ...
func (Grub) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:grub", 0
}

// Name ...
func (Grub) Name() string {
	return "Grub"
}

// Texture ...
func (Grub) Texture() image.Image {
	return textureFromName("grub")
}
