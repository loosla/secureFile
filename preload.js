const { contextBridge, ipcRenderer } = require('electron/renderer');

contextBridge.exposeInMainWorld('api', {
  fetchText: () => ipcRenderer.invoke('fetch-text'),
  saveText: (data) => ipcRenderer.invoke('save-text', data),
  fetchFile: (password) => ipcRenderer.invoke('fetch-file', password),
});
