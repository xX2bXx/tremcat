package main

import (
	"log"
	"math/rand/v2"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	// check if chafa is installed on a host machine
	if _, err := exec.LookPath("chafa"); err != nil {
		log.Fatalf("error: chafa not found")
	}
	
	r := gin.Default()

	// get all the filenames in pics/
	entries, err := os.ReadDir("pics")
	if err != nil {
	    log.Fatal(err)
	}

	files := make([]string, 0, len(entries))
	
	for _, entry := range entries {
	    if !entry.IsDir() {
	        files = append(files, "pics/" + entry.Name())
	    }
	}
	
	r.GET("/", func(c *gin.Context) {
		// get user's terminal size
		w := c.DefaultQuery("w", "80")
		h := c.DefaultQuery("h", "45")

		randPic := files[rand.IntN(len(files))]
		out, err := exec.Command("chafa", "--size", w+"x"+h, "--symbols=block", "--colors=full", randPic).Output()
		if err != nil {
			log.Printf("chafa failed: %v", err)
			c.String(500, "failed to render a cat")
			return
		}
		
		c.Data(200, "text/plain; charset=utf-8", out)
	})
	r.Run(":8080")
}
