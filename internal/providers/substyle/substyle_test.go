package substyle_test

import (
	"testing"

	"github.com/wraient/curd/internal/curdhost"
	"github.com/wraient/curd/internal/providers/substyle"
)

func TestChooseHardPromptsForSoftFallback(t *testing.T) {
	substyle.ResetForTest()
	prompts := 0
	previousPrompt := curdhost.PromptSelect
	curdhost.Out = func(string) {}
	curdhost.PromptSelect = func(options []curdhost.PromptOption) (curdhost.PromptOption, error) {
		prompts++
		return curdhost.PromptOption{Key: "soft", Label: options[0].Label}, nil
	}
	t.Cleanup(func() {
		curdhost.PromptSelect = previousPrompt
		substyle.ResetForTest()
	})

	style, err := substyle.Choose(true, false, "hard")
	if err != nil || style != "soft" {
		t.Fatalf("expected soft after prompt, got style=%q err=%v", style, err)
	}
	if prompts != 1 {
		t.Fatalf("expected 1 prompt, got %d", prompts)
	}
}

func TestChooseHardDeclinedSoftFallback(t *testing.T) {
	substyle.ResetForTest()
	previousPrompt := curdhost.PromptSelect
	curdhost.PromptSelect = func([]curdhost.PromptOption) (curdhost.PromptOption, error) {
		return curdhost.PromptOption{Key: "cancel", Label: "Cancel"}, nil
	}
	t.Cleanup(func() {
		curdhost.PromptSelect = previousPrompt
		substyle.ResetForTest()
	})

	_, err := substyle.Choose(true, false, "hard")
	if err == nil {
		t.Fatal("expected error when soft fallback is declined")
	}
}

func TestChooseSoftFallsBackToHard(t *testing.T) {
	style, err := substyle.Choose(false, true, "soft")
	if err != nil || style != "hard" {
		t.Fatalf("expected hard fallback, got style=%q err=%v", style, err)
	}
}

func TestChooseAskAutoUsesSoftOnly(t *testing.T) {
	substyle.ResetForTest()
	messages := 0
	previousOut := curdhost.Out
	curdhost.Out = func(string) { messages++ }
	t.Cleanup(func() {
		curdhost.Out = previousOut
		substyle.ResetForTest()
	})

	style, err := substyle.Choose(true, false, "ask")
	if err != nil || style != "soft" {
		t.Fatalf("expected soft, got style=%q err=%v", style, err)
	}
	if messages != 1 {
		t.Fatalf("expected status message, got %d", messages)
	}
}

func TestChooseAskDeclinedSoftOnly(t *testing.T) {
	t.Skip("ask mode auto-accepts soft-only streams")
}
