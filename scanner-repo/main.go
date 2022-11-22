package main

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/name"
)

func main() {
	image := "docker.io/library/nats:2.9.3-alpine"
	imageID := "docker.io/library/nats@sha256:e0615d6d59dad37b1eee587e2e89d2feee1d93537c4f5d68a6550685f0d1ab48"
	r1, err := name.ParseReference(image)
	if err != nil {
		return
	}
	r2, err := name.ParseReference(imageID)
	if err != nil {
		return
	}
	fmt.Printf("r1 = %+v \n", r1)
	fmt.Printf("r2 = %+v \n", r2)
	fmt.Println(r1.Context(), " ", r2.Context())
}
