package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedTBoneSteak struct{}

// AlwaysConsumable ...
func (CookedTBoneSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedTBoneSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedTBoneSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(16, 9.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedTBoneSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedTBoneSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedTBoneSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_t_bone_steak", 0
}

// Name ...
func (CookedTBoneSteak) Name() string {
	return "Cooked T Bone Steak"
}

// Texture ...
func (CookedTBoneSteak) Texture() image.Image {
	return textureFromName("steak_t_bone2")
}
