package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PotatoChipsSourCreamFlavor struct{}

// AlwaysConsumable ...
func (PotatoChipsSourCreamFlavor) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PotatoChipsSourCreamFlavor) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PotatoChipsSourCreamFlavor) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (PotatoChipsSourCreamFlavor) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PotatoChipsSourCreamFlavor) Edible() bool {
	return true
}

// EncodeItem ...
func (PotatoChipsSourCreamFlavor) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:potato_chips_sour_cream_flavor", 0
}

// Name ...
func (PotatoChipsSourCreamFlavor) Name() string {
	return "Potato Chips Sour Cream Flavor"
}

// Texture ...
func (PotatoChipsSourCreamFlavor) Texture() image.Image {
	return textureFromName("potatochip_green")
}
