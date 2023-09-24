package service

import (
	"bbs/internal/configs"
	"log/slog"
	"net"
	"os"
	"syscall"
)

func APIServer() {
	serverSocketFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		slog.Error("Error creating socket", "error", err)
		os.Exit(1)
	}
	defer syscall.Close(serverSocketFd)

	serverAddr := syscall.SockaddrInet4{Port: configs.Cfg.APIServer.Port}
	copy(serverAddr.Addr[:], net.ParseIP("0.0.0.0").To4())
	if err = syscall.Bind(serverSocketFd, &serverAddr); err != nil {
		slog.Error("Error binding socket", "error", err)
		os.Exit(1)
	}

	if err = syscall.Listen(serverSocketFd, 5); err != nil {
		slog.Error("Error listening", "error", err)
		os.Exit(1)
	}

	slog.Info("Socket created, bound and listening.", "Port", configs.Cfg.APIServer.Port)
	for {
		nfd, clientAddr, err := syscall.Accept(serverSocketFd)
		if err != nil {
			slog.Error("Connection accepted error", "Error:", err)
			continue
		}
		go handleConnection(nfd, clientAddr)
	}
}

func handleConnection(fd int, clientAddr syscall.Sockaddr) {
	defer syscall.Close(fd)
	slog.Debug("Connection accepted.", "fd", fd, "sa", clientAddr)
	var buf [32 * 1024]byte

	syscall.Write(fd, []byte("Welcome to bbs!\n"))

	for {
		nr, err := syscall.Read(fd, buf[:])
		if err != nil {
			slog.Error("Error reading from socket.", "fd", fd, "error", err)
			return
		}
		if nr == 0 {
			slog.Debug("EOF received, closing connection.", "fd", fd)
			return
		}
		slog.Debug("Read data from socket.", "fd", fd, "data", string(buf[:nr]))
	}
}
