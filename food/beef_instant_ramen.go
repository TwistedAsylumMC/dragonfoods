package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BeefInstantRamen struct{}

// AlwaysConsumable ...
func (BeefInstantRamen) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BeefInstantRamen) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BeefInstantRamen) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BeefInstantRamen) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BeefInstantRamen) Edible() bool {
	return true
}

// EncodeItem ...
func (BeefInstantRamen) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:beef_instant_ramen", 0
}

// Name ...
func (BeefInstantRamen) Name() string {
	return "BeefInstantRamen"
}

// Texture ...
func (BeefInstantRamen) Texture() image.Image {
	return textureFromName("instantramen4")
}
