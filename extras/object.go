package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

const (
	MyObjectStore = "nice"
	//MyObjectStore = "trivy"
	ObjectName = "db.tar.gz"
)

func main() {
	store := getStore()
	fmt.Println("Got the Store")
	listObjects(store)

	fmt.Println("Lets Put something into an object")
	//putString, err := store.PutString(ObjectName, "hello world")
	//if err != nil {
	//	return
	//}
	//fmt.Printf("%+v \n", putString)

	info, err := store.PutFile("/home/arnob/something.txt")
	// `nats object get nice /home/arnob/something.txt -O hi.txt`
	if err != nil {
		return
	}
	fmt.Printf("Put command output: %+v \n", info)
}

func getStore() nats.ObjectStore {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println(err)
	}
	jsContext, err := nc.JetStream()
	if err != nil {
		panic(err)
	}

	store, err := jsContext.ObjectStore(MyObjectStore)
	if err != nil { // stream not found
		if store, err = jsContext.CreateObjectStore(&nats.ObjectStoreConfig{
			Bucket: MyObjectStore,
		}); err != nil {
			panic(err)
		}
	}
	return store
}

func listObjects(store nats.ObjectStore) {
	list, err := store.List()
	if err != nil {
		return
	}
	for _, info := range list {
		fmt.Printf("%+v %+v %+v %+v \n", info.ObjectMeta, info.Bucket, info.Size, info.Digest)
	}
}

func getObject(store nats.ObjectStore) {
	info, err := store.GetInfo(ObjectName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetInfo output: %+v \n", info)
}
