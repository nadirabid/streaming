# Gen HLS files

```bash
VID_NAME=IJS9E1 && FORMAT=mkv; eval 'mkdir $VID_NAME; ffmpeg -i $VID_NAME.$FORMAT  -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list $VID_NAME/outputlist.m3u8 -segment_format mpegts $VID_NAME/output%03d.ts'
```

```bash
VID_NAME=IJS9E1 && FORMAT=mkv; eval 'mkdir $VID_NAME; ffmpeg -i $VID_NAME.$FORMAT -vf scale=w=1280:h=720:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -b:a 128k -c:v h264 -profile:v main -crf 20 -g 48 -keyint_min 48 -sc_threshold 0 -b:v 2500k -maxrate 2675k -bufsize 3750k -hls_time 4 -hls_playlist_type vod -hls_segment_filename $VID_NAME/720p_%03d.ts $VID_NAME/720p.m3u8'
```
