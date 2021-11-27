package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	tempDir, err := createTempDir()
	if err != nil {
		log.Fatal(err)
	}

	srcDir, err := createDir(tempDir, "src")
	if err != nil {
		log.Fatal(err)
	}

	// シンボリックリンク(管理者権限必要)
	{
		err := createSymlink(srcDir, filepath.Join(tempDir, "symlink"))
		if err != nil {
			log.Fatal(err)
		}
	}

	// ジャンクション
	{
		err := createJunction(srcDir, filepath.Join(tempDir, "junc"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createSymlink(src string, dest string) error {

	err := os.Symlink(src, dest)
	if err != nil {
		return err
	}

	return nil
}

func createJunction(src string, dest string) error {

	err := mklink("J", dest, src)
	if err != nil {
		return err
	}

	return nil
}

func mklink(linktype string, link string, target string) error {

	output, err := exec.Command("cmd", "/c", "mklink", "/"+linktype, link, target).CombinedOutput()
	if err != nil {
		return fmt.Errorf("\"mklink /%s %v %v\" command failed: %v\n%v", linktype, link, target, err, string(output))
	}

	return nil
}

func createTempDir() (string, error) {

	return os.MkdirTemp("", "win-link")
}

func createDir(base string, name string) (string, error) {

	dir := filepath.Join(base, name)

	err := os.Mkdir(dir, 0755)
	if err != nil {
		return "", err
	}

	return dir, nil
}
