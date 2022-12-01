package main

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

const (
	PublicImage  = "arnobkumarsaha/my-mongo:01"
	PrivateImage = "arnobkumarsaha/private-test:01"
)

func main() {
	var err error
	//TryPublicImage()
	err = TryPrivateImage()
	if err != nil {
		fmt.Println("error occurred : ", err)
	}
}

func TryPublicImage() error {
	pubRef, err := name.ParseReference(PublicImage)
	if err != nil {
		return err
	}
	pubDes, err := remote.Get(pubRef, remote.WithAuth(authn.Anonymous))
	if err != nil {
		return err
	}
	print(pubDes, "Public")
	return nil
}

func TryPrivateImage() error {
	priRef, err := name.ParseReference(PrivateImage)
	if err != nil {
		return err
	}
	PriDes, err := remote.Get(priRef, remote.WithAuth(authn.FromConfig(authn.AuthConfig{
		Username: "arnobkumarsaha",
		Password: "sustcse16",
	})))
	if err != nil {
		return err
	}
	print(PriDes, "Private")
	return nil
}

func print(d *remote.Descriptor, typ string) {
	fmt.Printf("%s: %+v %+v %s \n", typ, d.Ref, d.Descriptor, string(d.Manifest))
}
