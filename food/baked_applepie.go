package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedApplepie struct{}

// AlwaysConsumable ...
func (BakedApplepie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedApplepie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedApplepie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedApplepie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedApplepie) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedApplepie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_applepie", 0
}

// Name ...
func (BakedApplepie) Name() string {
	return "BakedApplepie"
}

// Texture ...
func (BakedApplepie) Texture() image.Image {
	return textureFromName("bakedapplepie2")
}
