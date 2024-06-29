const { contextBridge, ipcRenderer } = require('electron/renderer');

contextBridge.exposeInMainWorld('api', {
  update: (data) => ipcRenderer.invoke('update', data),
  get: (password) => ipcRenderer.invoke('get', password),
});
