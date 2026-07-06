package senshi

import "testing"

func TestSubtitleInfoFromURL(t *testing.T) {
	rawURL := "https://embed.example/e/abc/?sub.info=https%3A%2F%2Fninstream.com%2Fmanifest.json"
	if got := subtitleInfoFromURL(rawURL); got != "https://ninstream.com/manifest.json" {
		t.Fatalf("unexpected manifest url %q", got)
	}
}

func TestPickSenshiSubtitleTrackPrefersEnglishDefault(t *testing.T) {
	tracks := []senshiSubtitleTrack{
		{Src: "https://cdn.example/jpn.vtt", Label: "JPN"},
		{Src: "https://cdn.example/eng.vtt", Label: "ENG", Default: true},
	}
	if got := pickSenshiSubtitleTrack(tracks); got != "https://cdn.example/eng.vtt" {
		t.Fatalf("unexpected subtitle %q", got)
	}
}

func TestSenshiSubtitleManifestURLUsesServerFM(t *testing.T) {
	serverFM := "https://embed.example/e/abc/?sub.info=https://ninstream.com/manifest.json"
	item := embedItem{ServerFM: &serverFM}
	if got := senshiSubtitleManifestURL(item); got != "https://ninstream.com/manifest.json" {
		t.Fatalf("unexpected manifest url %q", got)
	}
}

func TestSenshiSubtitleManifestURLFallsBackToMaskedBase(t *testing.T) {
	item := embedItem{MaskedBaseURL: "https://ninstream.com/example/base"}
	if got := senshiSubtitleManifestURL(item); got != "https://ninstream.com/example/base/sub_filemoon.json" {
		t.Fatalf("unexpected manifest url %q", got)
	}
}
