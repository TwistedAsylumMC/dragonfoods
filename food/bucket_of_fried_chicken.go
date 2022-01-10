package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BucketOfFriedChicken struct{}

// AlwaysConsumable ...
func (BucketOfFriedChicken) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BucketOfFriedChicken) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BucketOfFriedChicken) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BucketOfFriedChicken) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BucketOfFriedChicken) Edible() bool {
	return true
}

// EncodeItem ...
func (BucketOfFriedChicken) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bucket_of_fried_chicken", 0
}

// Name ...
func (BucketOfFriedChicken) Name() string {
	return "Bucket Of Fried Chicken"
}

// Texture ...
func (BucketOfFriedChicken) Texture() image.Image {
	return textureFromName("kfc_chicken")
}
