const { contextBridge, ipcRenderer } = require('electron/renderer');

contextBridge.exposeInMainWorld('api', {
  saveText: (data) => ipcRenderer.invoke('save-text', data),
  fetchFile: (password) => ipcRenderer.invoke('fetch-file', password),
});
