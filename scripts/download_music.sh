#!/bin/bash


PLAYLISTS=(
    "https://www.youtube.com/watch?v=h3WFI_NIZ2o&list=PLmB6td997u3nZxSalWwUfsmiOs_VtwltW"
    "https://www.youtube.com/watch?v=A8VqP1vSGpI&list=PLZCsQDrzneRDdymbhZawfrsa4WLlMd-58"
    "https://www.youtube.com/watch?v=VceA2FfMkFk&list=PLozxAo05ybUDzeGVFGI3P5cqktV2a_0Y2"
    "https://www.youtube.com/watch?v=NUsoVlDFqZg&list=PLH6S8OjNLi-IGP4W4r1pUoUP9wvUEsFik"
    "https://www.youtube.com/watch?v=b4i7tbqKWp4&list=PL78F9D32B2E4C1135"
    "https://www.youtube.com/watch?v=xRrKgH-HmpU&list=PLFAO0bzDKkCOCkfPlSYppb4f2fg2vSQtW"
    "https://www.youtube.com/watch?v=AXN-_asIaYs&list=PL4F3AF1CE5C1A08DD"
    "https://www.youtube.com/watch?v=jXQPKYKhwxk&list=PL8BvQM-ayPDrtRskhqs3ijZ36XIQohd3v"
    "https://www.youtube.com/watch?v=YjwjZTPEvC8&list=PL75862C390AE9B856"
    "https://www.youtube.com/watch?v=gxlB1B9emDc&list=PL8QxoDgfjMvJMP-DBgqUFWQh5xYaMJAU7"
    "https://www.youtube.com/watch?v=PIh2xe4jnpk&list=PLgaFNC_I_ZkngtWHjGxSasNShPp19C9m6"
    "https://www.youtube.com/watch?v=Kr4EQDVETuA&list=PL6kFT_LlrYKW32II058F7c_NcJ3uCMNh6"
    "https://www.youtube.com/watch?v=fJ9rUzIMcZQ&list=PLmB6td997u3n26ovU3ZZmonWyVO90Nmq4"
    "https://www.youtube.com/watch?v=o1tj2zJ2Wvg&list=PLA8ACC4996D23A2D1"
    "https://www.youtube.com/watch?v=vj_rvLVpqg8&list=PLqLF6xXygDUGR4nPIg2KOP4TTmbeXLWtR"
    "https://www.youtube.com/watch?v=oEauWw9ZGrA&list=PL3T9ZdKtd6BRc0ZradCa9dd7gts2kBYw7"
    "https://www.youtube.com/watch?v=NO7EtdR3Dyw&list=PLexunLlotTJ8pCJ-gX3UG1jc3SUGssMfm"
    "https://www.youtube.com/watch?v=DkFJE8ZdeG8&list=PLe_WTKC1u0ncKYGkb5MS09VVJvxQWSNSa"
    "https://www.youtube.com/watch?v=8OalWvMJlkc&list=PL7tBPYQzCjeJTz5xu8e2L7dE87CNEDMir"
    "https://www.youtube.com/watch?v=iFnwmTeSlAQ&list=PL8EF488611688C0D9"
    "https://www.youtube.com/watch?v=whBcmlaSLJM&list=PLE7B074B45B91749F"
    "https://www.youtube.com/watch?v=hyoS7js863E&list=PL6fAo_rDvM5jANriWYO0_mODzcpHXFcA7"
    "https://www.youtube.com/watch?v=MnjAeFNCyUQ&list=PLMWmCw09FICBBLUXowsPJxD-uc7mIgC3p"
    "https://www.youtube.com/watch?v=OC1nFl0xcg4&list=PLu6UQfy2_dxHpdY5iJOXSH3L1pCgiIA-m"
    "https://www.youtube.com/watch?v=Xk0wdDTTPA0&list=PL5qQ2aCLUdNzr21ySiTb37YGZeKxCZIk4"
    "https://www.youtube.com/watch?v=okc7Vw2_p7c&list=PLf0HxqT8vJNPiv0owkBHz5aPRR3J2OMHe"
    "https://www.youtube.com/watch?v=YlUKcNNmywk&list=PLxA687tYuMWgT1rkPWHTiijDTYuhF4xj_"
    "https://www.youtube.com/watch?v=dU_in_BNJlg&list=PLF8rV_MCWF291UADeW437yzQIyBb6TNwD"
    "https://www.youtube.com/watch?v=fLexgOxsZu0&list=PL2gNzJCL3m_9QZh_MFe4wWtnO3tl-bgg5"
    "https://www.youtube.com/watch?v=B3gbisdtJnA&list=PL4iSbgi3WlCrubl-onZiOa2f1TSFvItwW"
    "https://www.youtube.com/watch?v=DksSPZTZES0&list=PLarPDo5YMm5QHzwunrAQEzSSnN-Z8ALtk"
    "https://www.youtube.com/watch?v=WpYeekQkAdc&list=PLQlc99hV-nkHY3pljwYrS5DSIKafG9Kbl"
    "https://www.youtube.com/watch?v=rFjJs6ZjPe8&list=PLmndwJP2qZ1n530HY2g2OHTcqjRGB0xr9"
    "https://www.youtube.com/watch?v=hTWKbfoikeg&list=PL58492F7012617224"
    "https://www.youtube.com/watch?v=lsZG7n7ries&list=PLp3o6lwxnqMkm9dUECwqVBQY0IN6n1wH_"
    "https://www.youtube.com/watch?v=sZfZ8uWaOFI&list=PLEGyvPj66TeqE8zI1JGV6rxtDFxf0qWYq"
    "https://www.youtube.com/watch?v=IhP3J0j9JmY&list=PLMEZyDHJojxOivUPWX1aasnKcpau8WZfP"
    "https://www.youtube.com/watch?v=_FrOQC-zEog&list=PL1tiBqitg38_Rsqb2qiTvm3hKX2Y2qUgg"
    "https://www.youtube.com/watch?v=86URGgqONvA&list=PLgd0mymAKt3QBJJ288YKvvWaeJolNTYIt"
    "https://www.youtube.com/watch?v=kXYiU_JCYtU&list=PL9LkJszkF_Z6bJ82689htd2wch-HVbzCO"
)


CMD="yt-dlp"
SONGS_DIR="${SONGS_DIR:-.}"
FLAGS="-f bestaudio --extract-audio --audio-format opus --audio-quality 160K --embed-metadata -P"
LOG_FILE="download_log.txt"
MAX_CONCURRENT=8

log() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') - $1" | tee -a "$LOG_FILE"
}

download_playlist() {
    local playlist="$1"
    local thread_id="$2"
    log "[Thread-$thread_id] Iniciando descarga: $playlist"
    if $CMD $FLAGS "$SONGS_DIR" "$playlist"; then
        log "[Thread-$thread_id] Descarga completada: $playlist"
    else
        log "[Thread-$thread_id] Error en la descarga: $playlist"
    fi
}

main() {
    mkdir -p "$SONGS_DIR"
    > "$LOG_FILE"
    log "Inicio del proceso de descargas"

    for playlist in "${PLAYLISTS[@]}"; do
        while [ $(jobs -r | wc -l) -ge $MAX_CONCURRENT ]; do sleep 1; done
        download_playlist "$playlist" "$((${#PLAYLISTS[@]} - ${#PLAYLISTS[@]} + 1))" &
    done

    wait
    log "Todas las descargas han finalizado."
}

main
