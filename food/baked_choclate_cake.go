package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedChoclateCake struct{}

// AlwaysConsumable ...
func (BakedChoclateCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedChoclateCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedChoclateCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedChoclateCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedChoclateCake) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedChoclateCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_choclate_cake", 0
}

// Name ...
func (BakedChoclateCake) Name() string {
	return "Baked Choclate Cake"
}

// Texture ...
func (BakedChoclateCake) Texture() image.Image {
	return textureFromName("cake_chocolate")
}
