package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"./parser"
	"./win32"
)

type WindowOptions struct {
	Fullscreen, Sound   bool
	X, Y, Width, Height int
	VideoDirectory      string
}

type Configuration struct {
	VideoPlayerLocation   string
	VideoPlayerExecutable string
	WindowClassName       string
	Windows               []WindowOptions
}

func fetch_user_config() (config Configuration) {
	data, _ := ioutil.ReadFile("./config.json")
	parser.ParseJsonData(string(data), &config)
	return
}

func start_media_player(video_directory string, config Configuration, window WindowOptions) {
	mpv_args := []string{
		"--player-operation-mode=pseudo-gui",
		"--force-window=no",
		"--terminal=no",
		"--loop-playlist=inf",
	}

	if !window.Sound {
		mpv_args = append(mpv_args, "--no-audio")
	}

	mpv_args = append(mpv_args, video_directory)

	executable := fmt.Sprintf("%s/%s", config.VideoPlayerLocation, config.VideoPlayerExecutable)
	mpv := exec.Command(executable, mpv_args...)
	mpv.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	err := mpv.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func start_wallpaper(pid int, config Configuration, window WindowOptions) {
	wp_path := "wp"

	mv_args := fmt.Sprintf("mv --pid %d -x %d -y %d -w %d -h %d", pid, window.X, window.Y, window.Width, window.Height)
	add_args := fmt.Sprintf("add --pid %d ", pid)

	if window.Fullscreen {
		add_args += "--fullscreen "
	}

	cmd := exec.Command(wp_path, strings.Split(mv_args, " ")...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()

	cmd = exec.Command(wp_path, strings.Split(add_args, " ")...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}

func kill_previous_mpvs(config Configuration) {
	cmd := exec.Command("TASKKILL", "-f", "-IM", config.VideoPlayerExecutable)
	cmd.CombinedOutput()
}

func main() {
	config := fetch_user_config()

	kill_previous_mpvs(config)

	// start video players
	for _, window := range config.Windows {
		start_media_player(window.VideoDirectory, config, window)
	}

	// give all the media players proper time to startup
	time.Sleep(time.Second * 2)

	// fetch the pids
	mpv_pids := []uintptr{}
	handles := win32.GetAllProcessHandles()
	for _, handle := range handles {
		if win32.GetWindowClassName(handle) == config.WindowClassName {
			mpv_pids = append(mpv_pids, win32.GetWindowProcessId(handle))
		}
	}

	// move the wallpaper
	for i, window := range config.Windows {
		start_wallpaper((int)(mpv_pids[i]), config, window)
	}
}
