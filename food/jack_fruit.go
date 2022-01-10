package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type JackFruit struct{}

// AlwaysConsumable ...
func (JackFruit) AlwaysConsumable() bool {
	return false
}

// Category ...
func (JackFruit) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (JackFruit) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (JackFruit) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (JackFruit) Edible() bool {
	return true
}

// EncodeItem ...
func (JackFruit) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:jack_fruit", 0
}

// Name ...
func (JackFruit) Name() string {
	return "Jack Fruit"
}

// Texture ...
func (JackFruit) Texture() image.Image {
	return textureFromName("jackfruit_01")
}
