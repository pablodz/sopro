package audioconfig

import (
	"fmt"
	"log"
	"strings"
)

// WavConfig is a struct that contains the configuration of a WAV audio file.
type WavConfig struct {
	BitDepth   int   // the bit depth (e.g. 8, 16, 24)
	Channels   int   // the number of channels (e.g. 1, 2)
	Encoding   int   // the encoding format (e.g. "PCM", "IEEE_FLOAT")
	SampleRate int   // the sample rate in Hertz
	WaveFormat uint8 // the wave format (e.g. "PCM", "IEEE_FLOAT", "ALAW", "MULAW")
}

// Audio formats Wav RIFF (4 bytes)
// Documentation: https://www-mmsp.ece.mcgill.ca/Documents/AudioFormats/WAVE/WAVE.html
const (
	// Audio Format WAV RIFF: PCM (4 bytes)
	WAVE_FORMAT_PCM = 0x0001
	// Audio Format WAV RIFF: IEEE Float (4 bytes)
	WAVE_FORMAT_IEEE_FLOAT = 0x0003
	// Audio Format WAV RIFF: ALAW (4 bytes)
	WAVE_FORMAT_ALAW = 0x0006
	// Audio Format WAV RIFF: MULAW (4 bytes)
	WAVE_FORMAT_MULAW = 0x0007
)

// PrintWavHeaders prints the headers of a WAV file
// first 44 bytes of a WAV file
func PrintWavHeaders(headersWav []byte) {
	if len(headersWav) != 44 {
		log.Println("[ERROR] Headers are not 44 bytes long")
		return
	}
	fmt.Println("Headers (WAV):")
	comments := []string{
		"(4) Chunk ID [RIFF]",
		"(4) Chunk size",
		"(4) Format [WAVE]",
		"(4) Sub-chunk 1 ID [fmt ]",
		"(4) Sub-chunk 1 size",
		"(2) Audio format (PCM) & (2) Number of channels",
		"(4) Sample rate",
		"(4) Byte rate",
		"(2) Block align & (2) Bits per sample",
		"(4) Sub-chunk 2 ID [data]",
		"(4) Sub-chunk 2 size",
	}
	for i := 0; i < 44; i += 4 {
		fmt.Println(
			fmt.Sprintf("[%2d,%2d]", i, i+4),
			fmt.Sprintf("% 2x", headersWav[i:i+4]),
			"\t<"+strings.ToUpper(string(headersWav[i:i+4]))+">\t",
			comments[i/4],
		)
	}
}
