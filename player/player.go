package player

import (
        "os/exec"
        "fmt"
)

func RunVideo(url, title string) (string, error) {
        video_title := fmt.Sprintf("Title:%0s\n", title)
        fmt.Println(video_title)
        cmd := exec.Command("termux-open-url", url)
        cmd.Run()
        return video_title, nil
        return url, nil
}
