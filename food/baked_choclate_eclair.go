package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedChoclateEclair struct{}

// AlwaysConsumable ...
func (BakedChoclateEclair) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedChoclateEclair) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedChoclateEclair) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedChoclateEclair) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedChoclateEclair) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedChoclateEclair) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_choclate_eclair", 0
}

// Name ...
func (BakedChoclateEclair) Name() string {
	return "Baked Choclate Eclair"
}

// Texture ...
func (BakedChoclateEclair) Texture() image.Image {
	return textureFromName("bakedchoceclair")
}
