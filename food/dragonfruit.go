package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Dragonfruit struct{}

// AlwaysConsumable ...
func (Dragonfruit) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Dragonfruit) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Dragonfruit) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Dragonfruit) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Dragonfruit) Edible() bool {
	return true
}

// EncodeItem ...
func (Dragonfruit) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:dragonfruit", 0
}

// Name ...
func (Dragonfruit) Name() string {
	return "Dragonfruit"
}

// Texture ...
func (Dragonfruit) Texture() image.Image {
	return textureFromName("dragonfruit2")
}
