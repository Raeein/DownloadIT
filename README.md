# DownloadIT

## About The Project

Download free audio books from https://goldenaudiobook.com and https://dailyaudiobooks.com, this project is only for educational purposes and you should pay for
the books your read (or listen). 


### Built With

- Best language ever created by man kind to this day: Go
- ffmpeg if you want to merge the audio file pieces into one


## Getting Started

### Prerequisites

* ffmpeg
* go
* A book you want to do educational research on

### Installation and Usage

Below is the good old clone and run instruction. You can also install the binary with go install

1. Clone the repo
```bash
git clone https://github.com/Raeein/DownloadIT
```
2. Sample run
```bash
go run main.go download -u https://goldenaudiobook.com/james-clear-atomic-habits-audiobook/
Finding audio files...
Downloading  https://ipaudio.club/wp-content/uploads/GOLN/Atomic%20Habits/06.mp3
Downloading  https://ipaudio.club/wp-content/uploads/GOLN/Atomic%20Habits/03.mp3
Downloading  https://ipaudio.club/wp-content/uploads/GOLN/Atomic%20Habits/01.mp3
Downloading  https://ipaudio.club/wp-content/uploads/GOLN/Atomic%20Habits/04.mp3
Downloading  https://ipaudio.club/wp-content/uploads/GOLN/Atomic%20Habits/02.mp3
Downloading  https://ipaudio.club/wp-content/uploads/GOLN/Atomic%20Habits/05.mp3
Downloaded 06.mp3
Downloaded 02.mp3
Downloaded 03.mp3
Downloaded 05.mp3
Downloaded 04.mp3
Downloaded 01.mp3

Finished downloading files
```

## Roadmap

- [ ] Support other websites to download from...?


## License

Distributed under the MIT License. See `LICENSE.txt` for more information.
