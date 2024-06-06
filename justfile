root := justfile_directory() 

# Directories
scripts_dir := root / "scripts"
songs_dir := root / "songs"

# Commands
deno_exec := "deno run --allow-read --allow-write --allow-net --allow-env"

default:
    just --list

download-songs:
    SONGS_DIR={{ songs_dir }} \
    WEB_ENDPOINT={{ "https://mp3teca.co" }} \
    SERVER_ENDPOINT={{ "https://severmp3teca.xyz/-/mp3" }} \
    {{ deno_exec }} {{ scripts_dir}}/download-music.ts