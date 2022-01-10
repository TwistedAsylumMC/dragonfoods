package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Beandip struct{}

// AlwaysConsumable ...
func (Beandip) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Beandip) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Beandip) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Beandip) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Beandip) Edible() bool {
	return true
}

// EncodeItem ...
func (Beandip) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:beandip", 0
}

// Name ...
func (Beandip) Name() string {
	return "Beandip"
}

// Texture ...
func (Beandip) Texture() image.Image {
	return textureFromName("beandip")
}
