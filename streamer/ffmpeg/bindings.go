package ffmpeg

import "strings"

// ffmpeg
//
//	-i ${NAME}
//	-c:v libx264
//	-preset veryfast
//	-g 48
//	-keyint_min 48
//	-sc_threshold 0
//	-b:v 2500k
//	-maxrate 2500k
//	-bufsize 5000k
//	-c:a aac
//	-b:a 128k
//	-hls_time 10
//	-hls_playlist_type vod
//	-hls_segment_filename "hls/${NAME}-%03d.ts" hls/${NAME}-playlist.m3u8
type ffmpegQueryBuilder struct {
	input  string
	Output string
	query  []string
}

func New() *ffmpegQueryBuilder {
	return &ffmpegQueryBuilder{
		input: "ffmpeg/bin/ffmpeg-6.0-amd64-static/ffmpeg",
	}
}

func (f *ffmpegQueryBuilder) SetInput(input string) *ffmpegQueryBuilder {
	f.input = input
	return f
}

func (f *ffmpegQueryBuilder) AddParam(output string) *ffmpegQueryBuilder {
	f.query = append(f.query, output)
	return f
}

func (f *ffmpegQueryBuilder) AddKeyValue(param, value string) *ffmpegQueryBuilder {
	f.query = append(f.query, param, value)
	return f
}

func (f *ffmpegQueryBuilder) Build() string {
	return f.input + " " + strings.Join(f.query, " ") + " " + f.Output
}

func (f *ffmpegQueryBuilder) Program() string {
	return f.input
}

func (f *ffmpegQueryBuilder) Args() []string {
	return f.query
}

type QueryBuilder interface {
	Program() string
	Args() []string
}
