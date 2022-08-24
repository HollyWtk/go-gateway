package utils

import (
	"github.com/nacos-group/nacos-sdk-go/model"
	"math/rand"
)

func WeightRandom(instances []model.Instance) model.Instance {
	if len(instances) == 1 {
		return instances[0]
	}
	var sum = 0.0
	for _, i := range instances {
		sum += i.Weight
	}
	r := rand.Float64() * sum
	var t = 0.0
	for i, w := range instances {
		t += w.Weight
		if t > r {
			return instances[i]
		}
	}
	return instances[len(instances)-1]
}
