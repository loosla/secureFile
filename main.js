const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path');
const axios = require('axios');
const { spawn } = require('child_process');

function createWindow () {
  const mainWindow = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js')
    }
  });

  mainWindow.loadURL('http://localhost:3000');
}

app.whenReady().then(() => {
  // Start the Go server
  const goServer = spawn('go', ['run', 'main.go']);

  goServer.stdout.on('data', (data) => {
    console.log(`Go server stdout: ${data}`);
  });

  goServer.stderr.on('data', (data) => {
    console.error(`Go server stderr: ${data}`);
  });

  // IPC handlers
  ipcMain.handle('fetch-text', async () => {
    const response = await axios.get('http://localhost:8080/api/text');
    return response.data;
  });

  ipcMain.handle('save-text', async (event, data) => {
    await axios.post('http://localhost:8080/api/save', data);
  });

  ipcMain.handle('fetch-file', async (event, password) => {
    try {
      const response = await axios.post('http://localhost:8080/get-file', { password });
      return response.data.content;
    } catch (error) {
      console.error('Error fetching file:', error);
      return 'Error fetching file';
    }
  });

  createWindow();

  app.on('activate', function () {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

app.on('window-all-closed', function () {
  if (process.platform !== 'darwin') app.quit();
});
