package mediainfo

import "context"

type MediaInfo struct {
	Path              string
	Duration          float64
	Bitrate           int64
	Size              int64
	VideoCodec        string
	Resolution        string
	Width             int
	Height            int
	AspectRatio       string
	FrameRate         float64
	BitDepth          int
	HDR               string
	ColorSpace        string
	AudioCodecs       []string
	AudioChannels     []string
	AudioLanguages    []string
	Atmos             bool
	SubtitleLanguages []string
	SubtitleFormats   []string
	ForcedSubtitles   []string
	Container         string
	Chapters          int
}

type MediaInfoProvider interface {
	Analyze(ctx context.Context, path string) (MediaInfo, error)
	Probe(ctx context.Context, path string) (MediaInfo, error)
}

const EventMediaAnalyzed = "media.analyzed"

type MediaAnalyzedPayload struct {
	Path       string  `json:"path"`
	Codec      string  `json:"codec"`
	Resolution string  `json:"resolution"`
	Duration   float64 `json:"duration"`
}
