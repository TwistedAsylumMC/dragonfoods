package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateChipBagel struct{}

// AlwaysConsumable ...
func (ChoclateChipBagel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateChipBagel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateChipBagel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateChipBagel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateChipBagel) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateChipBagel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_chip_bagel", 0
}

// Name ...
func (ChoclateChipBagel) Name() string {
	return "Choclate Chip Bagel"
}

// Texture ...
func (ChoclateChipBagel) Texture() image.Image {
	return textureFromName("bagel3")
}
