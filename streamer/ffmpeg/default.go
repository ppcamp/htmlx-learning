package ffmpeg

import (
	"fmt"

	"github.com/ppcamp/movies-to-watch/streamer/config"
)

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
func DefaultQuery(input string) *ffmpegQueryBuilder {
	q := NewQuery().
		SetInput(input).
		SetVideoCodec("libx264").
		SetAudioCodec("aac").
		SetPreset(VeryFast).
		SetVideoBitRate("2500k").
		SetAudioBitRate("128k").
		SetHlsPlaylistType(Vod).
		SetHlsTsDuration(10).
		SetHlsTsNamer(fmt.Sprintf("hls/%s-%%05d.ts", input)).
		SetOutput(fmt.Sprintf("hls/%s", DefaultFileName(input)))
		// AddKeyValue("g", "48").
		// AddKeyValue("keyint_min", "48").
		// AddKeyValue("sc_threshold", "0").
		// AddKeyValue("maxrate", "2500k").
		// AddKeyValue("bufsize", "5000k").
	return q
}

func DefaultFileName(input string) string {
	return fmt.Sprintf("%s-playlist.%s", input, config.PlaylistExt())
}
