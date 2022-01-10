package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ZucchiniFlowers struct{}

// AlwaysConsumable ...
func (ZucchiniFlowers) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ZucchiniFlowers) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ZucchiniFlowers) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (ZucchiniFlowers) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ZucchiniFlowers) Edible() bool {
	return true
}

// EncodeItem ...
func (ZucchiniFlowers) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:zucchini_flowers", 0
}

// Name ...
func (ZucchiniFlowers) Name() string {
	return "ZucchiniFlowers"
}

// Texture ...
func (ZucchiniFlowers) Texture() image.Image {
	return textureFromName("zucchiniflowers")
}
