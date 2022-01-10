package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SpaghettiSquash struct{}

// AlwaysConsumable ...
func (SpaghettiSquash) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SpaghettiSquash) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SpaghettiSquash) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (SpaghettiSquash) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SpaghettiSquash) Edible() bool {
	return true
}

// EncodeItem ...
func (SpaghettiSquash) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spaghetti_squash", 0
}

// Name ...
func (SpaghettiSquash) Name() string {
	return "SpaghettiSquash"
}

// Texture ...
func (SpaghettiSquash) Texture() image.Image {
	return textureFromName("spaghettisquash")
}
