package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MreTaco struct{}

// AlwaysConsumable ...
func (MreTaco) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MreTaco) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MreTaco) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (MreTaco) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MreTaco) Edible() bool {
	return true
}

// EncodeItem ...
func (MreTaco) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mre_taco", 0
}

// Name ...
func (MreTaco) Name() string {
	return "Mre Taco"
}

// Texture ...
func (MreTaco) Texture() image.Image {
	return textureFromName("mre3")
}
