package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	CHANGELOG = "CHANGELOG.md"
)

func main() {
	// 获取当前在 Docker 中运行的服务及其版本
	dockerVersions := getDockerVersions()

	// 读取用户输入的新版本号
	var newVersionNumber string
	fmt.Print("请输入新的版本号（格式：x.y.z）: ")
	fmt.Scanln(&newVersionNumber)

	// 获取最新版本记录的行号范围
	startLine, endLine := getVersionLines()

	// 提取最新版本的服务和版本信息到map
	latestVersions := getLatestVersions(startLine, endLine)

	// 创建新的版本记录
	newVersion := createNewVersionRecord(newVersionNumber, dockerVersions, latestVersions)

	// 将新的版本记录插入到 CHANGELOG.md 文件的第三行
	insertNewVersionRecord(newVersion)

	fmt.Println("CHANGELOG.md 更新成功，添加了版本", newVersionNumber, "的信息。")

	// 获取版本号
	version := getCurrentVersion()
	fmt.Println("本次发版版本:", version)

	// 获取当前日期
	now := time.Now().Format("2006-01-02")
	fmt.Println("本次发版日期:", now)

	// 获取本次发版内容
	upd := getUpdateContent()
	fmt.Println("本次要更新的容器有:\n", upd)

	// 确认上传 CHANGELOG 到 master 分支
	var add string
	fmt.Print("即将上传本次发版的CHANGELOG到master,请确认本次更新内容是否正确,yes or no: ")
	fmt.Scanln(&add)
	if add == "yes" {
		execGitCommands(CHANGELOG, now)
	} else {
		fmt.Println("未确认，将退出")
		os.Exit(0)
	}

	// 确认创建标签
	var tag string
	fmt.Print("即将对本次发版", version, "打标记,请确认是否创建带标签的版本,yes or no: ")
	fmt.Scanln(&tag)
	if tag == "yes" {
		createGitTag(version, now)
	} else {
		fmt.Println("未确认，将退出")
		os.Exit(0)
	}
}

func getDockerVersions() []string {
	cmd := exec.Command("sh", "-c", "docker -H tcp://192.168.0.214:2375 ps --format '{{.Names}}' | grep online | awk -F '[-]' '{gsub(/^v/, \"\", $3); print $2 \"-\" $3}'")
	out, _ := cmd.Output()
	return strings.Split(strings.TrimSpace(string(out)), "\n")
}

func getVersionLines() (int, int) {
	file, _ := os.Open(CHANGELOG)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var startLine, endLine int
	count := 0
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if match, _ := regexp.MatchString(`### version [0-9]+\.[0-9]+\.[0-9]+`, line); match {
			count++
			if count == 1 {
				startLine = lineNum
			}
			if count == 2 {
				endLine = lineNum
				break
			}
		}
	}
	return startLine, endLine
}

func getLatestVersions(startLine, endLine int) map[string]string {
	file, _ := os.Open(CHANGELOG)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	latestVersions := make(map[string]string)

	for scanner.Scan() {
		lineNum++
		if lineNum >= startLine && lineNum < endLine {
			line := scanner.Text()
			if strings.Contains(line, "|") {
				fields := strings.FieldsFunc(line, func(c rune) bool {
					return c == '|'
				})
				if len(fields) > 2 {
					service := strings.TrimSpace(fields[1])
					version := strings.TrimSpace(fields[2])
					latestVersions[service] = version
				}
			}
		}
	}
	return latestVersions
}

func createNewVersionRecord(newVersionNumber string, dockerVersions []string, latestVersions map[string]string) string {
	newVersion := fmt.Sprintf("### version %s\n\n1. 线上服务版本更新(NOCHANGE将不更新)\n\n| 服务     | 版本   | 备注      |\n|:--------:|:------:|:---------:|", newVersionNumber)

	for _, line := range dockerVersions {
		parts := strings.Split(line, "-")
		service, version := parts[0], parts[1]
		remark := ""
		if latestVersions[service] == version {
			remark = "NOCHANGE"
		}
		newVersion += fmt.Sprintf("\n| %-9s | %-7s | %-8s |", service, version, remark)
	}

	return "\n" + newVersion + "\n"
}

func insertNewVersionRecord(newVersion string) {
	file, _ := os.Open(CHANGELOG)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	output := append([]string{lines[0], lines[1], newVersion}, lines[2:]...)
	_ = os.WriteFile(CHANGELOG, []byte(strings.Join(output, "\n")), 0644)
}

func getCurrentVersion() string {
	file, _ := os.Open(CHANGELOG)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if match := regexp.MustCompile(`version [0-9.]+`).FindString(line); match != "" {
			return "v" + strings.TrimPrefix(match, "version ")
		}
	}
	return ""
}

func getUpdateContent() string {
	file, _ := os.Open(CHANGELOG)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	var capture bool

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "### version") && capture {
			break
		}
		if capture {
			lines = append(lines, strings.TrimSpace(line))
		}
		if strings.Contains(line, "### version") {
			capture = true
		}
	}

	var result []string
	for _, line := range lines {
		if strings.HasPrefix(line, "|") && !strings.Contains(line, "NOCHANGE") {
			fields := strings.FieldsFunc(line, func(c rune) bool {
				return c == '|'
			})
			if len(fields) > 2 {
				result = append(result, fmt.Sprintf("%s %s", strings.TrimSpace(fields[1]), strings.TrimSpace(fields[2])))
			}
		}
	}
	sort.Strings(result) // Sort to maintain consistent output
	return strings.Join(result, "\n")
}

func execGitCommands(changelog, now string) {
	exec.Command("git", "add", changelog).Run()
	exec.Command("git", "commit", "-m", "[onlinegolive] "+now).Run()
	exec.Command("git", "push", "origin", "master").Run()
	fmt.Println("推送成功")
}

func createGitTag(version, now string) {
	exec.Command("git", "tag", "-a", version, "-m", "[onlinegolive] "+now).Run()
	exec.Command("git", "push", "origin", version).Run()
	fmt.Println("标签", version, "创建成功")
	fmt.Println("标记push成功,已触发jenkins的online cicd脚本,请去jenkins查看:\nhttp://192.168.2.254:8998/jenkins/job/bbx-tag")
}
