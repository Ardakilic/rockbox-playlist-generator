Rockbox Playlist Generator
--------

Rockbox Playlist Generator is a tool that automatically creates **Rockbox-compatible M3U playlists** per artist.  
It fetches the **top songs from Last.fm**, cross-checks them with **your locally stored music files**, and generates **playlists** based on matches.

## 🚀 Features
- 📡 **Fetches top songs** for each artist from **Last.fm**.
- 🎵 **Cross-checks with local music** files (MP3, FLAC, AAC, ALAC, OGG, WAV, M4A).
- 🎶 **Uses Album Artist** if available, otherwise falls back to Artist.
- 🎧 **Creates Rockbox-compatible M3U playlists**.
- 🛠️ **Configurable number of top songs** to fetch (default: **50**).
- 🖥️ **Fast and lightweight** (written in Golang 1.24).
- 🔐 **Environment-based configuration** with `.env` support.
- 🐳 **Docker support** for easy deployment.
- 🔌 **Uses [shkh/lastfm-go](https://github.com/shkh/lastfm-go)** library for Last.fm API integration.

## 📦 Installation

### Option 1: Download Pre-built Binary
You can download pre-built binaries for various platforms from the [releases page](https://github.com/Ardakilic/rockbox-playlist-generator/releases).

### Option 2: Native Installation

#### 1️⃣ Clone the repository
```sh
git clone https://github.com/Ardakilic/rockbox-playlist-generator.git
cd rockbox-playlist-generator
```

#### 2️⃣ Set up environment variables
Create a `.env` file by copying the example:
```sh
cp .env.example .env
```

Then edit the `.env` file with your Last.fm API credentials:
```
LASTFM_API_KEY=your_lastfm_api_key_here
LASTFM_API_SECRET=your_lastfm_api_secret_here
DEFAULT_TRACK_LIMIT=50
```

You can get a **free Last.fm API key** at [Last.fm API](https://www.last.fm/api/account/create).

#### 3️⃣ Build the application
```sh
go mod tidy
go build
```

### Option 3: Docker Installation

#### 1️⃣ Clone the repository
```sh
git clone https://github.com/Ardakilic/rockbox-playlist-generator.git
cd rockbox-playlist-generator
```

#### 2️⃣ Set up environment variables
```sh
cp .env.example .env
```

Edit the `.env` file with your Last.fm API credentials and paths:
```
# API credentials
LASTFM_API_KEY=your_lastfm_api_key_here
LASTFM_API_SECRET=your_lastfm_api_secret_here

# Optional settings
DEFAULT_TRACK_LIMIT=50
MUSIC_PATH=/path/to/your/music
OUTPUT_PATH=/path/to/your/playlists
```

#### 3️⃣ Run with Docker Compose
```sh
# Build and run with docker-compose
docker-compose -f deployments/compose.yaml up --build

# Run in detached mode
docker-compose -f deployments/compose.yaml up -d

# View logs
docker-compose -f deployments/compose.yaml logs -f

# Stop the application
docker-compose -f deployments/compose.yaml down
```

## ⚙️ Usage

### Native Usage
```sh
./rockbox-playlist-generator -path /your/music/folder
```

### Options
- `-path` → Path to your local music library (required)
- `-limit` → Number of top tracks to fetch per artist (default: from config or 50)

### Docker Usage
```sh
# Using docker-compose with values from .env file
docker-compose -f deployments/compose.yaml up

# Direct Docker run with environment variables
docker run -v /your/music/folder:/music:ro \
           -v /your/playlists/folder:/music/playlists:rw \
           -e LASTFM_API_KEY=your_key \
           -e LASTFM_API_SECRET=your_secret \
           rockbox-playlist-generator
```

#### Environment Variables for Docker
- `LASTFM_API_KEY` - Your Last.fm API key (required)
- `LASTFM_API_SECRET` - Your Last.fm API secret (required)
- `DEFAULT_TRACK_LIMIT` - Default number of tracks to fetch per artist (default: 50)
- `MUSIC_PATH` - Path to your music library on the host (default: /music)
- `OUTPUT_PATH` - Path to store playlists on the host (default: ./playlists)
- `TRACK_LIMIT` - Override the default track limit for this run (default: 0, which uses DEFAULT_TRACK_LIMIT)

## 🎵 Example Output
For an artist like **Tool**, the script generates:
```
rockbox_tool.m3u
```
With contents like:
```
#EXTM3U
#EXTINF:445,Tool - Schism (2001 - Lateralus)
/your/music/folder/Tool/Lateralus/Schism.mp3
#EXTINF:563,Tool - Lateralus (2001 - Lateralus)
/your/music/folder/Tool/Lateralus/Lateralus.flac
#EXTINF:423,Tool - The Pot (2006 - 10,000 Days)
/your/music/folder/Tool/10,000 Days/The Pot.mp3
...
```
💡 **Fully compatible with Rockbox!**  
📌 **Includes album names and release years** for better organization.

## 🏗️ Project Structure

```
rockbox-playlist-generator/
├── config/         # Application configuration
├── deployments/    # Deployment configurations
├── pkg/            # Reusable packages
│   ├── lastfm/     # Last.fm API client (using shkh/lastfm-go)
│   ├── music/      # Music file handling
│   └── playlist/   # Playlist generation
├── .env.example    # Example environment variables
├── Dockerfile      # Docker build definition
├── go.mod          # Go module definition
├── main.go         # Application entry point
└── README.md       # This file
```

## 📝 License
This project is licensed under the **MIT License**.

### 🔥 **Improvements**
✅ **Environment-based configuration** with `.env` support  
✅ **Clear separation of concerns** with modular package design  
✅ **Improved error handling** and configuration validation  
✅ **Last.fm API integration** using [shkh/lastfm-go](https://github.com/shkh/lastfm-go)  
✅ **Docker support** with simplified configuration  
✅ **Updated to Go 1.24** for better performance and features