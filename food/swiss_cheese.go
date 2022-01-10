package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SwissCheese struct{}

// AlwaysConsumable ...
func (SwissCheese) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SwissCheese) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SwissCheese) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (SwissCheese) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SwissCheese) Edible() bool {
	return true
}

// EncodeItem ...
func (SwissCheese) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:swiss_cheese", 0
}

// Name ...
func (SwissCheese) Name() string {
	return "Swiss Cheese"
}

// Texture ...
func (SwissCheese) Texture() image.Image {
	return textureFromName("cheese3")
}
