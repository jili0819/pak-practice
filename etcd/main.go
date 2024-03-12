package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	// 链接etcd
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("connect to etcd success")
	defer func() {
		_ = etcdCli.Close()
	}()
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	// put
	_, err = etcdCli.Put(ctx, "q1mi", "dsb")

	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("pus etcd success")
	// get
	resp, err := etcdCli.Get(ctx, "key1")
	if err != nil {
		log.Fatal("get err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		fmt.Println(ev.Lease)
	}

	// watch one key
	//ch1 := etcdCli.Watch(ctx, "k8s:")
	// watch all key
	/*ch1 := etcdCli.Watch(ctx, "k8s:", clientv3.WithPrefix())
	go func(ctx context.Context) {
		for {
			select {
			case ssss := <-ch1:
				for _, evv := range ssss.Events {
					fmt.Printf("Type :%s,key:%s,value:%s", evv.Type, evv.Kv.Key, evv.Kv.Value)
					fmt.Println()
				}
			case <-ctx.Done():
				fmt.Println("go func() close!!!")
				return
			}
		}
	}(ctx)
	time.Sleep(5 * time.Second)*/
}
