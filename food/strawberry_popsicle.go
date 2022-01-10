package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type StrawberryPopsicle struct{}

// AlwaysConsumable ...
func (StrawberryPopsicle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (StrawberryPopsicle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (StrawberryPopsicle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (StrawberryPopsicle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (StrawberryPopsicle) Edible() bool {
	return true
}

// EncodeItem ...
func (StrawberryPopsicle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:strawberry_popsicle", 0
}

// Name ...
func (StrawberryPopsicle) Name() string {
	return "StrawberryPopsicle"
}

// Texture ...
func (StrawberryPopsicle) Texture() image.Image {
	return textureFromName("ice_cream_bar_01")
}
