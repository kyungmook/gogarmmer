package pkglib

import (
	"log"
)

func privateTest() {
	log.Printf("private function");
}

func PublicTest() {
	log.Printf("public function");
}

func init() {
	log.Printf("pkg lib initialize");
}


