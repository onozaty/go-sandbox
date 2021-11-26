package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	src := "src"

	// シンボリックリンク(管理者権限必要)
	{
		err := os.Symlink(src, "symlink")
		if err != nil {
			log.Fatal(err)
		}
	}

	// ジャンクション
	{
		err := mklink("J", "junk", src)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func mklink(linktype string, link string, target string) error {

	output, err := exec.Command("cmd", "/c", "mklink", "/"+linktype, link, target).CombinedOutput()
	if err != nil {
		return fmt.Errorf("\"mklink /%s %v %v\" command failed: %v\n%v", linktype, link, target, err, string(output))
	}

	return nil
}
