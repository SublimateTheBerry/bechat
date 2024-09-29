<div align="center">
  <h1>BeChat 💬</h1>
  <p>A simple chat application that integrates with Twitch and Goodgame platforms, perfect for OBS!</p>
  
  ![Twitch](https://img.shields.io/badge/Twitch-%239146FF.svg?style=for-the-badge&logo=Twitch&logoColor=white)
  ![Goodgame](https://img.shields.io/badge/Goodgame-%23404d59.svg?style=for-the-badge&logo=Goodgame&logoColor=white)
  ![GO](https://img.shields.io/badge/Go-1.22.7-blue.svg?style=for-the-badge)
  ![WebSocket](https://img.shields.io/badge/WebSocket-%23404d59.svg?style=for-the-badge&logo=WebSocket&logoColor=%23E0234E)
  ![Cross-platform](https://img.shields.io/badge/Platform-Cross--platform-important?style=for-the-badge)
  ![Release date](https://img.shields.io/github/release-date/SublimateTheBerry/bechat?style=for-the-badge)

  ![Example](https://github.com/user-attachments/assets/e2ec01ce-3dcc-4f1a-b152-bd2965dbd174)

</div>

## 📖 About
BeChat is a simple chat application that allows users to view and interact with chat messages from Twitch and Goodgame platforms. The application uses WebSocket technology to provide real-time updates to the chat interface, making it perfect for use with OBS (Open Broadcaster Software)!

## 🚀 Features
- Integrates with Twitch and Goodgame chat platforms
- Real-time chat message updates
- Displays username, color, and message content
- Identifies the platform the message originated from

## 🔧 Installation and Usage
1. Download the pre-built, cross-platform binary from the [latest release](https://github.com/SublimateTheBerry/bechat/releases/latest).
2. Extract the downloaded archive and run the appropriate executable for your operating system.
3. Run the application with the desired platform flags:

Windows:
```
bechat --tu=your-twitch-username --ggu=your-goodgame-username
```
Linux, MacOS:
```
chmod +x bechat
./bechat --tu=your-twitch-username --ggu=your-goodgame-username
```

Replace `your-twitch-username` and `your-goodgame-username` with the appropriate usernames.

4. In OBS, add a new Browser Source and point it to `http://localhost:12597`. This will display the chat interface in your OBS scene.

## 🧑‍💻 Manual Build
If you prefer to build the application from source, follow these steps:

1. Clone the repository:
```
git clone https://github.com/SublimateTheBerry/bechat.git
```
2. Navigate to the project directory:
```
cd bechat
```
3. Build the application:
```
go mod init bechat
go get github.com/gempir/go-twitch-irc/v3
go get github.com/gorilla/websocket
go build main.go bechat
```
4. Run the application with the desired platform flags:

Windows:
```
bechat --tu=your-twitch-username --ggu=your-goodgame-username
```
Linux, MacOS:
```
chmod +x bechat
./bechat --tu=your-twitch-username --ggu=your-goodgame-username
```

Replace `your-twitch-username` and `your-goodgame-username` with the appropriate usernames.

5. In OBS, add a new Browser Source and point it to `http://localhost:12597`. This will display the chat interface in your OBS scene.

## 🤝 Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## 📄 License
This project is licensed under the [MIT License](LICENSE).
