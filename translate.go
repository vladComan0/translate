package translate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const googleTranslateEndpoint = "https://translate.googleapis.com/translate_a/single?client=gtx"

func Translate(text, srcLang, destLang string) (string, error) {
	var translatedSentences []string
	
	endpoint := fmt.Sprintf("%s&sl=%s&tl=%s&dt=t&q=%s", googleTranslateEndpoint, srcLang, destLang, url.QueryEscape(text))
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", fmt.Errorf("could not fetch translation: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	translationBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response body: %w", err)
	}

	var translation googleTranslateResponse
	if err := json.Unmarshal(translationBody, &translation); err != nil {
		return "", fmt.Errorf("could not unmarshal translation body: %w", err)
	}

	for _, sentence := range translation.Sentences {
		translatedSentences = append(translatedSentences, sentence.TranslatedText)
	}

	return strings.Join(translatedSentences, ""), nil
}

type googleTranslateResponse struct {
	Sentences      []sentence
	DetectedSource string
}

type sentence struct {
	TranslatedText string
	OriginalText   string
}

func (r *googleTranslateResponse) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) > 0 {
		_ = json.Unmarshal(raw[0], &r.Sentences)
	}

	if len(raw) > 2 {
		_ = json.Unmarshal(raw[2], &r.DetectedSource)
	}

	return nil
}

func (s *sentence) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) > 0 {
		_ = json.Unmarshal(raw[0], &s.TranslatedText)
	}

	if len(raw) > 1 {
		_ = json.Unmarshal(raw[1], &s.OriginalText)
	}

	return nil
}
