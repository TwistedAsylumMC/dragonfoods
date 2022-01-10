package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HotdogWithStuff struct{}

// AlwaysConsumable ...
func (HotdogWithStuff) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HotdogWithStuff) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HotdogWithStuff) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (HotdogWithStuff) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HotdogWithStuff) Edible() bool {
	return true
}

// EncodeItem ...
func (HotdogWithStuff) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:hotdog_with_stuff", 0
}

// Name ...
func (HotdogWithStuff) Name() string {
	return "HotdogWithStuff"
}

// Texture ...
func (HotdogWithStuff) Texture() image.Image {
	return textureFromName("55_hotdog_sauce")
}
