package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawTBoneSteak struct{}

// AlwaysConsumable ...
func (RawTBoneSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawTBoneSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawTBoneSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawTBoneSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawTBoneSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (RawTBoneSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_t_bone_steak", 0
}

// Name ...
func (RawTBoneSteak) Name() string {
	return "RawTBoneSteak"
}

// Texture ...
func (RawTBoneSteak) Texture() image.Image {
	return textureFromName("steak_t-bone")
}
