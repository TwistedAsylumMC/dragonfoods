package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Celery struct{}

// AlwaysConsumable ...
func (Celery) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Celery) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Celery) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Celery) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Celery) Edible() bool {
	return true
}

// EncodeItem ...
func (Celery) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:celery", 0
}

// Name ...
func (Celery) Name() string {
	return "Celery"
}

// Texture ...
func (Celery) Texture() image.Image {
	return textureFromName("celery_01")
}
