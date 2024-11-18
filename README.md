<p align="center">
  <img src="web/src/arlin.svg" style="width: 150px;"/>
  <h1 align="center">Arlin</h1>
</p>


**Arlin** is an open-source app that allows you to send remote input from your Android device to a Linux machine connected to the same network. Using **Zeroconf** for device discovery and **WebSocket** for seamless communication, Arlin provides an intuitive and user-friendly experience. Inspired by **KDE Connect**, Arlin focuses on simplicity while offering reliable remote control functionalities.

---

## Features

- **Seamless Remote Input**: Control your Linux system from your Android device with ease.
- **Zeroconf (mDNS)**: Automatically discover devices on the same local network without requiring any manual configuration.
- **WebSocket Communication**: Fast and secure real-time input transfer between Android and Linux.
- **Material You UI**: A modern, clean, and intuitive user interface designed for ease of use.
- **Android to Linux**: Send keyboard and mouse events from your Android device to your Linux system.

---

## Requirements

- **Android device** (running Android 5.0 or higher).
- **Linux machine** (Ubuntu, Fedora, or any Linux-based system).
- **Local network**: Both Android and Linux devices must be connected to the same network.

---

## Installation

### On Android:
1. Download and install **Arlin** from the [releases page](https://github.com/srshazin/arlin-client/releases) (for APK).
2. Open the app on your Android device.
3. Ensure your Android device is connected to the same Wi-Fi network as your Linux machine.

### On Linux:
- Make sure you have `bash`, `curl` and `zenity` installed. Than use this command to install
- Run the installer script
     ```bash
       curl -f https://raw.githubusercontent.com/srshazin/arlin/refs/heads/main/install.sh | bash
     ```

---

## Usage

1. Open the Arlin app on your Android device.
2. The app will automatically detect your Linux machine via Zeroconf (mDNS) on the same local network.
3. Once the devices are paired, you can start controlling the Linux machine with your Android device by sending keyboard and mouse events.
4. The user interface offers easy-to-use controls for mouse movement, clicks, and keyboard input.

---

## Features Overview

- **Keyboard Input**: Send keypresses from your Android device to the Linux machine.
- **Mouse Input**: Move the cursor and simulate left/right clicks from your Android device.
- **Connection Status**: The app provides feedback on the connection status to ensure you're always informed about the status of the remote input.

---

## Troubleshooting

- **Cannot detect Linux machine**: Make sure both devices are on the same local network and that mDNS (Zeroconf) is not being blocked by your router or firewall.
- **Connection issues**: Ensure that the **Arlin server** is running on your Linux machine and there are no firewall rules blocking WebSocket connections.

---

## Contributing

Contributions are welcome! Feel free to open issues, fork the repository, and submit pull requests. Here are a few ways you can help:

- Bug reports or feature requests.
- Code improvements or bug fixes.
- Help with documentation.

---

## License

Arlin is licensed under the [MIT License](LICENSE).

---

## Acknowledgments

- **KDE Connect**: For inspiring the concept of the app.
- **Zeroconf/mDNS**: For local network device discovery.
- **WebSocket**: For real-time, bidirectional communication.
- **Material You**: For the modern and responsive UI design.

---

## Contact

For support or questions, you can reach us at:  
- Email: [contact@shazin.me](mailto:support@arlinapp.com)  
- GitHub: [https://github.com/yourusername/arlin](https://github.com/yourusername/arlin)

---

