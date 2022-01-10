package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CosmicBrownie struct{}

// AlwaysConsumable ...
func (CosmicBrownie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CosmicBrownie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CosmicBrownie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (CosmicBrownie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CosmicBrownie) Edible() bool {
	return true
}

// EncodeItem ...
func (CosmicBrownie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cosmic_brownie", 0
}

// Name ...
func (CosmicBrownie) Name() string {
	return "CosmicBrownie"
}

// Texture ...
func (CosmicBrownie) Texture() image.Image {
	return textureFromName("cosmicbrownie")
}
