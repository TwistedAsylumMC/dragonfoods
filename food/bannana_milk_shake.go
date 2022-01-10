package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BannanaMilkShake struct{}

// AlwaysConsumable ...
func (BannanaMilkShake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BannanaMilkShake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BannanaMilkShake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BannanaMilkShake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BannanaMilkShake) Edible() bool {
	return true
}

// EncodeItem ...
func (BannanaMilkShake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bannana_milk_shake", 0
}

// Name ...
func (BannanaMilkShake) Name() string {
	return "Bannana Milk Shake"
}

// Texture ...
func (BannanaMilkShake) Texture() image.Image {
	return textureFromName("shake2")
}
