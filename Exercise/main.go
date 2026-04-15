package main

// the file rescriptor is set non-blocking mode.
// if the file descriptor is not ready
// socket buffred is full socket buffer is empty
// go uses netpoller, is its abscration over syscall . let's file
//netpoller is like abscartion over syscall and it going to block the file descripiton into non blocking
//netpoller uses kquue in macos and epoll in lunx , iocp in windows
// lets say go routine is opening a call to netdial , netpoller is going to se the file descriptor into non blocking mode
//

func main() {

}
