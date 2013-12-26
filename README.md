lazors
======
package main

import "github.com/captncraig/lazors"


func main() {
    b := lazors.ClassicSetup()
	lazors.GetFullPath(&b, 0, lazors.South)
}
