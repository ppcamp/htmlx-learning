package ffmpeg

import (
	"fmt"
	"strings"
)

type TPreset int
type THlsPlaylist int

const (
	VeryFast TPreset = iota
	Faster
	Fast
	Medium
	Slow
	Slower
	VerySlow
)

const (
	Event THlsPlaylist = iota
	Vod
)

type ffmpegQueryBuilder struct {
	Input, Output string
	kw            map[string]string
}

func NewQuery() *ffmpegQueryBuilder {
	return &ffmpegQueryBuilder{
		Input: "ffmpeg/bin/ffmpeg-6.0-amd64-static/ffmpeg",
	}
}

func (f *ffmpegQueryBuilder) SetProgram(program string) *ffmpegQueryBuilder {
	f.Input = program
	return f
}

func (f *ffmpegQueryBuilder) SetOutput(output string) *ffmpegQueryBuilder {
	f.Output = output
	return f
}

func (f *ffmpegQueryBuilder) SetInput(i string) *ffmpegQueryBuilder {
	return f.AddKeyValue("i", i)
}

func (f *ffmpegQueryBuilder) SetVideoCodec(codec string) *ffmpegQueryBuilder {
	return f.AddKeyValue("c:v", codec)
}

func (f *ffmpegQueryBuilder) SetVideoBitRate(br string) *ffmpegQueryBuilder {
	return f.AddKeyValue("b:v", br)
}

func (f *ffmpegQueryBuilder) SetAudioBitRate(br string) *ffmpegQueryBuilder {
	return f.AddKeyValue("b:a", br)
}

func (f *ffmpegQueryBuilder) SetHlsTsDuration(s int) *ffmpegQueryBuilder {
	return f.AddKeyValue("hls_time", fmt.Sprint(s))
}

func (f *ffmpegQueryBuilder) SetHlsTsNamer(format string) *ffmpegQueryBuilder {
	return f.AddKeyValue("hls_segment_filename", format)
}

func (f *ffmpegQueryBuilder) SetHlsPlaylistType(s THlsPlaylist) *ffmpegQueryBuilder {
	switch s {
	case Event:
		return f.AddKeyValue("hls_playlist_type", "event")
	case Vod:
		return f.AddKeyValue("hls_playlist_type", "vod")
	default:
		return f
	}
}

func (f *ffmpegQueryBuilder) SetPreset(preset TPreset) *ffmpegQueryBuilder {
	switch preset {
	case VeryFast:
		return f.AddKeyValue("preset", "veryfast")
	case Faster:
		return f.AddKeyValue("preset", "faster")
	case Fast:
		return f.AddKeyValue("preset", "fast")
	case Medium:
		return f.AddKeyValue("preset", "medium")
	case Slow:
		return f.AddKeyValue("preset", "slow")
	case Slower:
		return f.AddKeyValue("preset", "slower")
	case VerySlow:
		return f.AddKeyValue("preset", "veryslow")
	default:
		return f
	}
}

func (f *ffmpegQueryBuilder) SetAudioCodec(codec string) *ffmpegQueryBuilder {
	return f.AddKeyValue("c:a", codec)
}

func (f *ffmpegQueryBuilder) AddParam(param string) *ffmpegQueryBuilder {
	f.kw[f.prepend(param)] = ""
	return f
}

func (f *ffmpegQueryBuilder) AddKeyValue(param, value string) *ffmpegQueryBuilder {
	f.kw[f.prepend(param)] = value
	return f
}

func (f *ffmpegQueryBuilder) prepend(param string) string {
	return fmt.Sprintf("-%s", param)
}

func (f *ffmpegQueryBuilder) Build() string {
	return f.Input + strings.Join(f.Args(), " ") + " " + f.Output
}

func (f *ffmpegQueryBuilder) Program() string {
	return f.Input
}

func (f *ffmpegQueryBuilder) Args() []string {
	s := []string{}
	for k, v := range f.kw {
		if v == "" {
			s = append(s, k)
		} else {
			s = append(s, k, v)
		}
	}
	return s
}

type QueryBuilder interface {
	Program() string
	Args() []string
}
