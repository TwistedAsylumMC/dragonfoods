package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RedDonut struct{}

// AlwaysConsumable ...
func (RedDonut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RedDonut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RedDonut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (RedDonut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RedDonut) Edible() bool {
	return true
}

// EncodeItem ...
func (RedDonut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:red_donut", 0
}

// Name ...
func (RedDonut) Name() string {
	return "Red Donut"
}

// Texture ...
func (RedDonut) Texture() image.Image {
	return textureFromName("35_donut")
}
