package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedVolcanoCake struct{}

// AlwaysConsumable ...
func (BakedVolcanoCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedVolcanoCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedVolcanoCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedVolcanoCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedVolcanoCake) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedVolcanoCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_volcano_cake", 0
}

// Name ...
func (BakedVolcanoCake) Name() string {
	return "BakedVolcanoCake"
}

// Texture ...
func (BakedVolcanoCake) Texture() image.Image {
	return textureFromName("backedvolcake")
}
