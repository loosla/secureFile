const { contextBridge, ipcRenderer } = require('electron/renderer');

contextBridge.exposeInMainWorld('api', {
  filesContent: (password) => ipcRenderer.invoke('files-content', password),
  filesSave: (data) => ipcRenderer.invoke('files-save', data),
});
