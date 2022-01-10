package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type TurkeyHotdog struct{}

// AlwaysConsumable ...
func (TurkeyHotdog) AlwaysConsumable() bool {
	return false
}

// Category ...
func (TurkeyHotdog) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (TurkeyHotdog) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (TurkeyHotdog) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (TurkeyHotdog) Edible() bool {
	return true
}

// EncodeItem ...
func (TurkeyHotdog) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:turkey_hotdog", 0
}

// Name ...
func (TurkeyHotdog) Name() string {
	return "TurkeyHotdog"
}

// Texture ...
func (TurkeyHotdog) Texture() image.Image {
	return textureFromName("turkeyhotdog")
}
