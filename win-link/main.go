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
	defer os.RemoveAll(tempDir)

	srcDir, err := createDir(tempDir, "src")
	if err != nil {
		log.Fatal(err)
	}

	// 適当なファイル作っておく
	f, err := os.Create(filepath.Join(srcDir, "a.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString("a")
	if err != nil {
		log.Fatal(err)
	}

	{
		err := createSymlink(srcDir, filepath.Join(tempDir, "symlink"))
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		err := createJunction(srcDir, filepath.Join(tempDir, "junc"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createSymlink(src string, link string) error {

	// シンボリックリンク作成(要管理者権限)
	err := os.Symlink(src, link)
	if err != nil {
		return err
	}

	// リンク元の情報取得
	linkSrc, err := os.Readlink(link)
	if err != nil {
		return err
	}

	fmt.Printf("シンボリックリンク リンク元: %s\n", linkSrc)

	err = os.Remove(link)
	if err != nil {
		return err
	}

	return nil
}

func createJunction(src string, link string) error {

	// ジャンクション作成
	err := mklink("J", link, src)
	if err != nil {
		return err
	}

	// リンク元の情報取得
	linkSrc, err := os.Readlink(link)
	if err != nil {
		return err
	}

	fmt.Printf("ジャンクション リンク元: %s\n", linkSrc)

	err = os.Remove(link)
	if err != nil {
		return err
	}

	return nil
}

func mklink(linktype string, link string, target string) error {

	output, err := exec.Command("cmd", "/c", "mklink", "/"+linktype, link, target).CombinedOutput()
	if err != nil {
		return fmt.Errorf("\"mklink /%s %s %s\" command failed: %v\n%s", linktype, link, target, err, string(output))
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
