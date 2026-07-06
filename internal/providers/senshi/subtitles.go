package senshi

import (
	"net/http"
	"net/url"
	"strings"
)

type senshiSubtitleTrack struct {
	Src     string `json:"src"`
	Label   string `json:"label"`
	Default bool   `json:"default"`
}

func senshiSubtitleManifestURL(item embedItem) string {
	if item.ServerFM != nil {
		if manifest := subtitleInfoFromURL(strings.TrimSpace(*item.ServerFM)); manifest != "" {
			return manifest
		}
	}
	base := strings.TrimSpace(item.MaskedBaseURL)
	if base == "" {
		return ""
	}
	return strings.TrimRight(base, "/") + "/sub_filemoon.json"
}

func subtitleInfoFromURL(rawURL string) string {
	if rawURL == "" {
		return ""
	}
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(parsed.Query().Get("sub.info"))
}

func fetchSenshiSubtitle(manifestURL string) (string, error) {
	manifestURL = strings.TrimSpace(manifestURL)
	if manifestURL == "" {
		return "", nil
	}

	var tracks []senshiSubtitleTrack
	if err := fetchJSON(http.MethodGet, manifestURL, nil, &tracks); err != nil {
		return "", err
	}
	return pickSenshiSubtitleTrack(tracks), nil
}

func pickSenshiSubtitleTrack(tracks []senshiSubtitleTrack) string {
	var fallback string
	for _, track := range tracks {
		file := strings.TrimSpace(track.Src)
		if file == "" {
			continue
		}
		label := strings.ToLower(strings.TrimSpace(track.Label))
		if track.Default || strings.Contains(label, "eng") {
			return file
		}
		if fallback == "" {
			fallback = file
		}
	}
	return fallback
}

func resolveSenshiSubtitle(item embedItem) string {
	manifestURL := senshiSubtitleManifestURL(item)
	if manifestURL == "" {
		return ""
	}
	subtitle, err := fetchSenshiSubtitle(manifestURL)
	if err != nil {
		return ""
	}
	return subtitle
}

func isSubEmbedStatus(status string) bool {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case "hardsub", "softsub", "sub":
		return true
	default:
		return false
	}
}
