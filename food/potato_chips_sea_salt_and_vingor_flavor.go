package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PotatoChipsSeaSaltAndVingorFlavor struct{}

// AlwaysConsumable ...
func (PotatoChipsSeaSaltAndVingorFlavor) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PotatoChipsSeaSaltAndVingorFlavor) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PotatoChipsSeaSaltAndVingorFlavor) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(9, 5.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (PotatoChipsSeaSaltAndVingorFlavor) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PotatoChipsSeaSaltAndVingorFlavor) Edible() bool {
	return true
}

// EncodeItem ...
func (PotatoChipsSeaSaltAndVingorFlavor) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:potato_chips_sea_salt_and_vingor_flavor", 0
}

// Name ...
func (PotatoChipsSeaSaltAndVingorFlavor) Name() string {
	return "Potato Chips Sea Salt And Vingor Flavor"
}

// Texture ...
func (PotatoChipsSeaSaltAndVingorFlavor) Texture() image.Image {
	return textureFromName("potatochip_blue")
}
