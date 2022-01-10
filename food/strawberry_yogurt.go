package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type StrawberryYogurt struct{}

// AlwaysConsumable ...
func (StrawberryYogurt) AlwaysConsumable() bool {
	return false
}

// Category ...
func (StrawberryYogurt) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (StrawberryYogurt) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (StrawberryYogurt) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (StrawberryYogurt) Edible() bool {
	return true
}

// EncodeItem ...
func (StrawberryYogurt) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:strawberry_yogurt", 0
}

// Name ...
func (StrawberryYogurt) Name() string {
	return "StrawberryYogurt"
}

// Texture ...
func (StrawberryYogurt) Texture() image.Image {
	return textureFromName("strawberry_ice_cream")
}
