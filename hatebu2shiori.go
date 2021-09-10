package main
import (
  "os"
  "path/filepath"
  "fmt"
  "log"

  "github.com/mmcdole/gofeed/atom"
)

type Hatebu struct {
  Title       string  `json:"title"`
  Link        string  `json:"link"`
  Tag         string  `json:"tag"`
  Comment     string  `json:"comment,omitempty"`
  Published   string  `json:"published"`
}

func main() {
  // Put file path
  fmt.Print("  plz hatena your.bookmarks.atom file FULL path\n  > ")
  var file_path string
  fmt.Scan(&file_path)

  // Read file
  file_dir := filepath.Dir(file_path)
  file, err := os.Open(file_path)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  // Parse atom
  fp := atom.Parser{}
  feed, _ := fp.Parse(file)
  fmt.Println(feed.Title)

  // Add hatebu
  var Hatebus [] Hatebu
  for _, entry := range feed.Entries {
    if entry == nil {
      break
    }

    hatebu := Hatebu {
      Title: entry.Title,
      Link: entry.Links[0].Href,
      Tag: "hatebu",
      Published: entry.Published,
    }

    if "" != entry.Summary {
      hatebu.Comment = entry.Summary
    }
    Hatebus = append(Hatebus, hatebu)
  }

  // Write shiori cmd
  sh_path := filepath.Join(file_dir, "hatebu2shiori.sh")
  if Exists(sh_path) {
    fmt.Println("ERROR: shell file exist " + sh_path)
    os.Exit(1)
  }
  for _, hatebu := range Hatebus {
    shiori_cmd(hatebu, sh_path)
  }

  fmt.Println("Success: " + sh_path)
}

func shiori_cmd(hatebu Hatebu, sh_path string)  {
  cmd := "shiori add " + quot(hatebu.Link) + " -t " + quot(hatebu.Tag)

  if "" != hatebu.Comment {
    cmd = cmd + " -e " + quot(hatebu.Comment)
  }

  file_write(sh_path, cmd)
}

func quot(s string) (qs string) {
  qs = "'" + s + "'"
  return
}

func file_write(file_path string, msg string) {
  file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0766)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  fmt.Fprintln(file, msg)
}

func Exists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}
