package win32

import (
	"bytes"
	"syscall"
	"unsafe"
)

var (
	user32DLL                    = syscall.NewLazyDLL("user32.dll")
	procGetTopWindow             = user32DLL.NewProc("GetTopWindow")
	procGetWindowTextW           = user32DLL.NewProc("GetWindowTextW")
	procGetWindow                = user32DLL.NewProc("GetWindow")
	procGetWindowThreadProcessId = user32DLL.NewProc("GetWindowThreadProcessId")
	procGetClassNameW            = user32DLL.NewProc("GetClassNameW")
)

func GetWindowText(hWnd uintptr) string {
	buffer := make([]byte, 1024)
	procGetWindowTextW.Call(hWnd, uintptr(unsafe.Pointer(&buffer[0])), uintptr(1024))

	buffer = bytes.Replace(buffer, []byte("\x00"), []byte(""), -1)
	return string(buffer)
}

func GetWindowClassName(hWnd uintptr) string {
	buffer := make([]byte, 1024)
	procGetClassNameW.Call(hWnd, uintptr(unsafe.Pointer(&buffer[0])), uintptr(1024))

	buffer = bytes.Replace(buffer, []byte("\x00"), []byte(""), -1)
	return string(buffer)
}

func GetWindowProcessId(hWnd uintptr) uintptr {
	pid := uintptr(0)
	procGetWindowThreadProcessId.Call(hWnd, uintptr(unsafe.Pointer(&pid)))
	return pid
}

func GetAllProcessHandles() (handles []uintptr) {
	// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-gettopwindow
	hWnd, _, err := procGetTopWindow.Call(uintptr(0))

	for err == syscall.Errno(0) {
		handles = append(handles, hWnd)

		// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getnextwindow
		hWnd, _, err = procGetWindow.Call(uintptr(hWnd), uintptr(2))
	}

	return
}
