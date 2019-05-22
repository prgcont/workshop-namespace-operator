package main

import (
	"fmt"

	wnov1alpha1 "github.com/prgcont/workshop-namespace-operator/pkg/apis/operator/v1alpha1"
)

func main() {
	fmt.Println(wnov1alpha1.WorkshopNamespace{})
}
