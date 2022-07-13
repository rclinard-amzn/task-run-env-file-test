// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"log"
)

func main() {
	log.Println(os.Getenv("MAGIC_VERB") + " " + os.Getenv("MAGIC_WORD"))
	log.Println(os.Getenv("TEST_MESSAGE"))
}
