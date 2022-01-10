package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MexicanStreetCorn struct{}

// AlwaysConsumable ...
func (MexicanStreetCorn) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MexicanStreetCorn) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MexicanStreetCorn) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (MexicanStreetCorn) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MexicanStreetCorn) Edible() bool {
	return true
}

// EncodeItem ...
func (MexicanStreetCorn) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mexican_street_corn", 0
}

// Name ...
func (MexicanStreetCorn) Name() string {
	return "Mexican Street Corn"
}

// Texture ...
func (MexicanStreetCorn) Texture() image.Image {
	return textureFromName("corn3")
}
