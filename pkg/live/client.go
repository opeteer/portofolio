package live

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/opeteer/opeteer-cli/pkg/data"
)

type NetworkProbe struct {
	Endpoint    string        `json:"endpoint"`
	LatencyMs   float64       `json:"latency_ms"`
	Duration    time.Duration `json:"duration"`
	Connected   bool          `json:"connected"`
	StatusCode  int           `json:"status_code"`
	RateLimit   int           `json:"rate_limit"`
	RateRemain  int           `json:"rate_remaining"`
	RateReset   time.Time     `json:"rate_reset"`
	Protocol    string        `json:"protocol"`
	ErrorMsg    string        `json:"error,omitempty"`
}

type LiveProfile struct {
	Login       string    `json:"login"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	PublicRepos int       `json:"public_repos"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Location    string    `json:"location"`
	AvatarURL   string    `json:"avatar_url"`
}

type LiveRepo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	HTMLURL     string    `json:"html_url"`
	PushedAt    time.Time `json:"pushed_at"`
	Stars       int       `json:"stargazers_count"`
	Forks       int       `json:"forks_count"`
	OpenIssues  int       `json:"open_issues_count"`
}

type LiveReport struct {
	Timestamp   time.Time    `json:"timestamp"`
	IsFallback  bool         `json:"is_fallback"`
	Probe       NetworkProbe `json:"probe"`
	Profile     LiveProfile  `json:"profile"`
	RecentRepos []LiveRepo   `json:"recent_repos"`
}

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	if timeout <= 0 {
		timeout = 5 * time.Second
	}
	return &Client{
		BaseURL: "https://api.github.com",
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) FetchLiveTelemetry(limit int) (*LiveReport, error) {
	if limit <= 0 {
		limit = 5
	}

	report := &LiveReport{
		Timestamp:   time.Now(),
		IsFallback:  false,
		RecentRepos: []LiveRepo{},
	}

	// 1. Probe & Fetch Profile
	profileURL := fmt.Sprintf("%s/users/opeteer", c.BaseURL)
	report.Probe.Endpoint = profileURL

	req, err := http.NewRequest("GET", profileURL, nil)
	if err != nil {
		return c.fallbackReport(err)
	}
	req.Header.Set("User-Agent", "opeteer-cli/1.0 (+https://github.com/opeteer/opeteer)")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	start := time.Now()
	resp, err := c.HTTPClient.Do(req)
	duration := time.Since(start)

	report.Probe.Duration = duration
	report.Probe.LatencyMs = float64(duration.Microseconds()) / 1000.0

	if err != nil {
		report.Probe.Connected = false
		report.Probe.ErrorMsg = err.Error()
		return c.fallbackReport(err)
	}
	defer resp.Body.Close()

	report.Probe.Connected = true
	report.Probe.StatusCode = resp.StatusCode
	report.Probe.Protocol = resp.Proto

	// Parse RateLimit headers
	if rem := resp.Header.Get("X-RateLimit-Remaining"); rem != "" {
		report.Probe.RateRemain, _ = strconv.Atoi(rem)
	}
	if lim := resp.Header.Get("X-RateLimit-Limit"); lim != "" {
		report.Probe.RateLimit, _ = strconv.Atoi(lim)
	}
	if rst := resp.Header.Get("X-RateLimit-Reset"); rst != "" {
		if sec, err := strconv.ParseInt(rst, 10, 64); err == nil {
			report.Probe.RateReset = time.Unix(sec, 0)
		}
	}

	if resp.StatusCode != http.StatusOK {
		report.Probe.ErrorMsg = fmt.Sprintf("GitHub API returned status %d", resp.StatusCode)
		return c.fallbackReport(fmt.Errorf("%s", report.Probe.ErrorMsg))
	}

	if err := json.NewDecoder(resp.Body).Decode(&report.Profile); err != nil {
		return c.fallbackReport(err)
	}

	if report.Profile.Username == "" && report.Profile.Login != "" {
		report.Profile.Username = report.Profile.Login
	}
	if report.Profile.Login == "" && report.Profile.Username != "" {
		report.Profile.Login = report.Profile.Username
	}

	// 2. Fetch Recent Pushed Repos
	reposURL := fmt.Sprintf("%s/users/opeteer/repos?sort=pushed&per_page=%d", c.BaseURL, limit)
	repoReq, err := http.NewRequest("GET", reposURL, nil)
	if err == nil {
		repoReq.Header.Set("User-Agent", "opeteer-cli/1.0")
		repoReq.Header.Set("Accept", "application/vnd.github.v3+json")

		if repoResp, err := c.HTTPClient.Do(repoReq); err == nil {
			defer repoResp.Body.Close()
			if repoResp.StatusCode == http.StatusOK {
				var repos []LiveRepo
				if err := json.NewDecoder(repoResp.Body).Decode(&repos); err == nil {
					report.RecentRepos = repos
				}
			}
		}
	}

	return report, nil
}

func (c *Client) fallbackReport(cause error) (*LiveReport, error) {
	report := &LiveReport{
		Timestamp:  time.Now(),
		IsFallback: true,
		Probe: NetworkProbe{
			Endpoint:   fmt.Sprintf("%s/users/opeteer", c.BaseURL),
			Connected:  false,
			ErrorMsg:   cause.Error(),
			LatencyMs:  0,
			RateLimit:  60,
			RateRemain: 0,
		},
		Profile: LiveProfile{
			Login:       data.MyProfile.Username,
			Username:    data.MyProfile.Username,
			Name:        data.MyProfile.Name,
			PublicRepos: len(data.AllProjects),
			Followers:   4,
			Following:   12,
			Location:    data.MyProfile.Location,
		},
	}

	// Fill recent repos from local dataset
	for i, p := range data.AllProjects {
		if i >= 5 {
			break
		}
		report.RecentRepos = append(report.RecentRepos, LiveRepo{
			Name:        p.Name,
			Description: p.Description,
			Language:    p.Language,
			HTMLURL:     p.URL,
			PushedAt:    time.Now().Add(-time.Duration(i*12) * time.Hour),
		})
	}

	return report, nil
}
