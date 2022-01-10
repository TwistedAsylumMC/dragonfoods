package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BoneMarrow struct{}

// AlwaysConsumable ...
func (BoneMarrow) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BoneMarrow) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BoneMarrow) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BoneMarrow) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BoneMarrow) Edible() bool {
	return true
}

// EncodeItem ...
func (BoneMarrow) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bone_marrow", 0
}

// Name ...
func (BoneMarrow) Name() string {
	return "BoneMarrow"
}

// Texture ...
func (BoneMarrow) Texture() image.Image {
	return textureFromName("bonemarrow")
}
