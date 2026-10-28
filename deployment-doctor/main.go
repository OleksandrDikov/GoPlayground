package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type ClientVersion struct {
	GitVersion string `json:"gitVersion"`
}

type VersionInfo struct {
	ClientVersion ClientVersion `json:"clientVersion"`
}

type Metadata struct {
	Metadata string `json:"metadata"`
}

func main() {
	fmt.Println("Deployment Doctor")
	namespace := "default"
	path, err := exec.LookPath("kubectl")
	if err != nil {
		fmt.Fprintln(os.Stderr, "kubectl not found:", err)
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		namespace = os.Args[1]
	}
	fmt.Println("Namespace:", namespace)

	fmt.Println("Path:", path)
	output, err := getClientVersion(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	clientVersion, err := parseClientVersion(output)
	if err != nil {
		fmt.Errorf("client version error: %v", err)
	}
	fmt.Println("Client Version:", clientVersion)
}

func getClientVersion(path string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, path, "version", "--client", "--output=json")
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("get kubectl client version: %w", err)
	}
	return output, nil
}

func parseClientVersion(data []byte) (string, error) {
	var version VersionInfo
	err := json.Unmarshal(data, &version)
	if err != nil {
		return "", fmt.Errorf("decode kubectl version: %w", err)
	}
	if version.ClientVersion.GitVersion == "" {
		return "", errors.New("kubectl version is empty")
	}
	return version.ClientVersion.GitVersion, nil
}
