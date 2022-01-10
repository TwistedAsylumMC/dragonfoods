package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RaspberryLemonMeringuePie struct{}

// AlwaysConsumable ...
func (RaspberryLemonMeringuePie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RaspberryLemonMeringuePie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RaspberryLemonMeringuePie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (RaspberryLemonMeringuePie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RaspberryLemonMeringuePie) Edible() bool {
	return true
}

// EncodeItem ...
func (RaspberryLemonMeringuePie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raspberry_lemon_meringue_pie", 0
}

// Name ...
func (RaspberryLemonMeringuePie) Name() string {
	return "RaspberryLemonMeringuePie"
}

// Texture ...
func (RaspberryLemonMeringuePie) Texture() image.Image {
	return textureFromName("64_lemonpie_dish")
}
