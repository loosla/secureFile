# secureFile

Desktop Application to Encrypt and Decrypt Text in a file with the Password

![Screenshot from 2024-06-09 17-07-02](https://github.com/loosla/secureFile/assets/12526985/ac92fb18-39a1-4d38-828d-4c96551c3070)

## How to use

1. Start the app
1. Type a password
1. Add your text to the text area: passwords or other secrets you would like to store safely
1. Hit "Save". Your text will be encrypted and saved in files/file.txt encrypted with the password set
1. Close the app

Now you can reopen and read your file using the password you set.

Note: Always remember the password you set before saving. If you change the password, the new password will be required to access the updated version of the file.

## Requirements

To run this application, you need to have the following installed on your system:

1. **Node.js and npm**

   - [Download and install Node.js](https://nodejs.org/)

2. **Go**
   - go version 1.21+
   - [Download and install Go](https://golang.org/dl/)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/loosla/secureFile.git
cd secureFile
```

### 2. Install Dependencies for React App

```bash
cd frontend
npm install
```

### 3. Build and Start the React App

```bash
npm run build
npm start-frontend
```

### 4. Start the Go Backend

Navigate to the backend directory and start the Go server:

```bash
cd ../backend
go run .
```

### 5. Install Dependencies for Electron

Navigate to electron directory and install Electron:

```bash
cd ../electron
npm install
```

### 6. Start the Electron App

Navigate to electron directory and start the Electron app:

```bash
cd ..
npm start-electron
```

## Quick start with everything pre-installed

```bash
npm start
```

# TODO

1. Add tests
1. Add config for host, port to share between BE and FE
1. Path to file or file name as an input?
