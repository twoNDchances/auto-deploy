package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh"
)

var (
	finalReport []map[string]string
	success int
	fail int
)

type AutoDeploy struct {
	Hostname     string
	Username     string
	Password     string
	Port         uint
	SiteFile     string
	ListCommands []string
}

func (ad *AutoDeploy) NewClient() (client *ssh.Client, err error) {
	sshConfig := &ssh.ClientConfig{
		User: ad.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(ad.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err = ssh.Dial("tcp", fmt.Sprintf("%s:%v", ad.Hostname, ad.Port), sshConfig)

	if err != nil {
		return nil, err
	}
	return
}

func (ad *AutoDeploy) Execute(client *ssh.Client, wg *sync.WaitGroup, site string) {
	defer wg.Done()
	for _, command := range ad.ListCommands {
		session, err := client.NewSession()
		if err != nil {
			log.Printf("Fail to create session for %s: %s", site, err)
			finalReport = append(finalReport, map[string]string{
				setColor(site, "blue"): setColor("failed", "red"),
			})
			fail++
			return
		}
		filterCmd := strings.ReplaceAll(command, "$site", site)
		output, err := session.CombinedOutput(filterCmd)
		if err != nil {
			log.Printf("Command fail at %s: %s", site, err)
			finalReport = append(finalReport, map[string]string{
				setColor(site, "blue"): setColor("failed", "red"),
			})
			fail++
			return
		}

		fmt.Printf("Output of '%s' at '%s':\n%s\n", filterCmd, site, output)
		session.Close()
		success++
	}
}

func (ad *AutoDeploy) LoadSiteFile() (*[]string, error) {
	var lines []string
	file, err := os.Open(ad.SiteFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &lines, nil
}
