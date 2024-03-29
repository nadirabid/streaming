#!/bin/sh

set -e # exit when any command fails

vlod() {
    # Usage create-vod-hls.sh SOURCE_FILE [OUTPUT_NAME]
    [[ ! "${1}" ]] && echo "Usage: create-vod-hls.sh SOURCE_FILE [OUTPUT_NAME]" && exit 1

    # comment/add lines here to control which renditions would be created
    renditions=(
        # resolution  bitrate  audio-rate
        #  "426x240    400k    64k"
        #  "640x360    800k     96k"
        "842x480    1400k    128k"
        "1280x720   2800k    128k"
        "1920x1080  5000k    192k"
    )

    segment_target_duration=4       # try to create a new segment every X seconds
    max_bitrate_ratio=1.07          # maximum accepted bitrate fluctuations
    rate_monitor_buffer_ratio=1.5   # maximum buffer size between bitrate conformance checks

    #########################################################################

    source="${1}"
    target="${2}"
    if [[ ! "${target}" ]]; then
    target="${source##*/}" # leave only last component of path
    target="${target%.*}"  # strip extension
    fi
    mkdir -p ${target}


    key_frames_interval="$(echo `ffprobe ${source} 2>&1 | grep -oE '[[:digit:]]+(.[[:digit:]]+)? fps' | grep -oE '[[:digit:]]+(.[[:digit:]]+)?'`*2 | bc || echo '')"
    key_frames_interval=${key_frames_interval:-50}
    key_frames_interval=$(echo `printf "%.1f\n" $(bc -l <<<"$key_frames_interval/10")`*10 | bc) # round
    key_frames_interval=${key_frames_interval%.*} # truncate to integer

    # static parameters that are similar for all renditions
    static_params="-c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0"
    static_params+=" -g ${key_frames_interval} -keyint_min ${key_frames_interval} -hls_time ${segment_target_duration}"
    static_params+=" -hls_playlist_type vod"

    # misc params
    misc_params="-hide_banner -y"

    master_playlist="#EXTM3U
    #EXT-X-VERSION:3
    "
    cmd=""
    for rendition in "${renditions[@]}"; do
    # drop extraneous spaces
    rendition="${rendition/[[:space:]]+/ }"

    # rendition fields
    resolution="$(echo ${rendition} | cut -d ' ' -f 1)"
    bitrate="$(echo ${rendition} | cut -d ' ' -f 2)"
    audiorate="$(echo ${rendition} | cut -d ' ' -f 3)"

    # calculated fields
    width="$(echo ${resolution} | grep -oE '^[[:digit:]]+')"
    height="$(echo ${resolution} | grep -oE '[[:digit:]]+$')"
    maxrate="$(echo "`echo ${bitrate} | grep -oE '[[:digit:]]+'`*${max_bitrate_ratio}" | bc)"
    bufsize="$(echo "`echo ${bitrate} | grep -oE '[[:digit:]]+'`*${rate_monitor_buffer_ratio}" | bc)"
    bandwidth="$(echo ${bitrate} | grep -oE '[[:digit:]]+')000"
    name="${height}p"
    
    cmd+=" ${static_params} -vf scale=w=${width}:h=${height}:force_original_aspect_ratio=decrease"
    cmd+=" -b:v ${bitrate} -maxrate ${maxrate%.*}k -bufsize ${bufsize%.*}k -b:a ${audiorate}"
    cmd+=" -hls_segment_filename ${target}/${name}_%03d.ts ${target}/${name}.m3u8"
    
    # add rendition entry in the master playlist
    master_playlist+="#EXT-X-STREAM-INF:BANDWIDTH=${bandwidth},RESOLUTION=${resolution}\n${name}.m3u8\n"
    done

    # start conversion
    echo -e "Executing command:\nffmpeg ${misc_params} -i ${source} ${cmd}"
    ffmpeg ${misc_params} -i ${source} ${cmd}

    # create master playlist file
    echo -e "${master_playlist}" > ${target}/playlist.m3u8

    echo "Done - encoded HLS is at ${target}/"
}

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

while getopts ":n:u:f:" arg; do
  case ${arg} in
    n ) 
        NAME=$OPTARG 
        ;;
    u ) 
        URL=$OPTARG 
        ;;
    f )
        FOLDER=$OPTARG 
        ;;
  esac
done

if [ -z "${BASE_DIR}" ]; then
    echo "BASE_DIR IS EMPTY"
    exit 1
fi

if [ -z "${NAME}" ]; then
    echo "NAME IS EMPTY"
    exit 1
fi

if [ -z "${URL}" ]; then
    echo "URL IS EMPTY"
    exit 1
fi

if [ -z "${FOLDER}" ]; then
    echo "FOLDER IS EMPTY"
    exit 1
fi

rm -rf $FOLDER/$NAME* # clean old video files

brew list youtube-dl || brew install youtube-dl

youtube-dl -f 'bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best' -o $FOLDER/$NAME.mp4 $URL

mkdir -p $FOLDER/$NAME/thumbnails

# cut to 30 seconds or hls will take time
ffmpeg -ss 00:00:00 -i $FOLDER/$NAME.mp4 -ss 00:00:00 -t 00:00:30 $FOLDER/$NAME/vid.mp4

# create hls
vlod $FOLDER/$NAME/vid.mp4 $FOLDER/$NAME/hls

# create thumbnails
for size in 842x480 1280x720 1920x1080
do
    ffmpeg -i $FOLDER/$NAME/vid.mp4 -vf  "thumbnail,scale=$size" -frames:v 1 $FOLDER/$NAME/thumbnails/$size.jpg
done

rm -rf $FOLDER/$NAME*.mp4
rm -rf $FOLDER/$NAME*/*.mp4
