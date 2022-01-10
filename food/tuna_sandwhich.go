package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type TunaSandwhich struct{}

// AlwaysConsumable ...
func (TunaSandwhich) AlwaysConsumable() bool {
	return false
}

// Category ...
func (TunaSandwhich) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (TunaSandwhich) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (TunaSandwhich) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (TunaSandwhich) Edible() bool {
	return true
}

// EncodeItem ...
func (TunaSandwhich) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tuna_sandwhich", 0
}

// Name ...
func (TunaSandwhich) Name() string {
	return "Tuna Sandwhich"
}

// Texture ...
func (TunaSandwhich) Texture() image.Image {
	return textureFromName("sandwich2")
}
