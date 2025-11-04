# Goch 🚀

A lightweight, terminal-first LAN chat app written in Go. Goch discovers peers on your local network and lets you chat in real time through a clean TUI powered by Bubble Tea.

---

## Features ✨

- LAN presence discovery via UDP broadcast (default port 8787)
- Direct user-to-user chat via UDP unicast (default port 8989)
- Modern terminal UI using Bubble Tea and Lipgloss
- Keyboard-first UX: select a user, type, Enter to send; quit with q or Ctrl+C
- Auto-detects subnet broadcast address for discovery
- Simple, readable Go code with internal packages

---

## Project structure 🗂️

```text
goch/
├─ cmd/                 # CLI entrypoints (Cobra)
├─ internal/
│  ├─ chat/             # Chat send/receive (UDP), message model
│  ├─ config/           # App configuration (ports, intervals)
│  ├─ discovery/        # Presence broadcast + discovery
│  ├─ tui/              # Bubble Tea TUI (model/view/update)
│  │  └─ shared/        # Styles
│  └─ utils/
│     └─ network/       # Network helpers (broadcast detection)
├─ Makefile             # build & test tasks
├─ main.go              # app bootstrap
└─ README.md            # this file
```

---

## Install, build, run 🛠️

Prereqs: Go 1.20+ recommended (works with Go modules), a local LAN.

```sh
git clone https://github.com/huseynovvusal/goch.git
cd goch
make build    # builds ./bin/goch
./bin/goch    # run
```

Run all tests:

```sh
make test
```

---

## How it works 🧠

- Discovery: each client periodically broadcasts its name on the LAN using UDP broadcast on port 8787. Peers listening on that port add/update the online user list.
- Chat: when you select a user, messages are sent directly to that peer’s IP using UDP unicast on port 8989. Messages are JSON-encoded and include sender info (name, IP) and content.

Key files:

- `internal/discovery/`: `BroadcastPresence`, `ListenForPresence`, in-memory online user tracking
- `internal/chat/`: `SendChatMessage`, `ListenForChatMessages`, `NetworkMessage` model
- `internal/tui/`: Bubble Tea model that renders online users, handles selection, and shows a chat pane + input
- `internal/utils/network/`: auto-detect subnet broadcast address via interface + mask math

---

## Usage ☕

1. Start `goch` on at least two machines on the same LAN (Wi‑Fi/Ethernet).
2. Enter your name.
3. Wait a moment for discovery; pick a user with ↑/↓, then press Enter to start chatting.
4. Type a message and press Enter to send.
5. Quit with `q` or `Ctrl+C`.

Notes:

- Your own identity (name + IP) is shown under the greeting and is not included in the online list.
- The first discovered user is auto-selected so you see a `>` cursor by default when the list is non-empty.

---

## Configuration ⚙️

Defaults live in `internal/config/config.go`.

- Presence (broadcast) port: `8787`
- Chat (unicast) port: `8989`
- Online user refresh interval: see `ONLINE_USERS_REFRESH_INTERVAL`

Network helper:

- The app computes the subnet broadcast address from the active IPv4 interface (network = ip & mask; broadcast = ip | ~mask) and uses it for presence broadcasting.

---

## Troubleshooting 🧩

If you don’t see other users:

- Ensure both machines are on the same subnet (e.g., 192.168.100.x). Check with `ip a` (Linux) or `ifconfig`/`ipconfig getifaddr en0` (macOS).
- Firewalls can block UDP broadcast/unicast. Temporarily disable to test: Linux `sudo ufw disable`, macOS System Settings → Network → Firewall → Off.
- Some routers block global broadcast `255.255.255.255`. The app auto-detects the subnet broadcast (e.g., `192.168.100.255`).
- AP/Client isolation: disable in your router if clients cannot see each other.
- VPNs/multiple interfaces can route packets incorrectly. Disconnect VPNs and use the same physical network.
- Use packet capture to confirm traffic:
  - On Linux: `sudo tcpdump -i any udp port 8787` (presence) or `sudo tcpdump -i any udp port 8989` (chat)
  - On macOS: `sudo tcpdump -i any udp port 8787` or `sudo tcpdump -i any udp port 8989`

If macOS sees HP but HP doesn’t see macOS:

- Verify you can ping each other both ways. If one direction fails, resolve network/firewall first.
- Confirm the app sends chat to the peer’s actual LAN IP (not broadcast) and listens on `0.0.0.0:8989`.

---

## Keyboard shortcuts ⌨️

- `Enter` — submit name / send message (context-sensitive)
- `↑/↓` — navigate users
- `q` / `Ctrl+C` — quit

---

## Contributing 🤝

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for how to set up your environment, run tests, and submit PRs.

Quick start:

```sh
git checkout -b feat/awesome
make test
make build
```

Open a PR against `main` with a clear description and screenshots (if UI changes).

---

## License 📄

This project is licensed under the Apache License 2.0 — see [LICENSE](LICENSE).

---

## Acknowledgments 🙏

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) for terminal styling

# Goch

Goch is a CLI chat app that lets you find and chat with people on your local area network.
