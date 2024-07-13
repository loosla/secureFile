const { contextBridge, ipcRenderer } = require('electron/renderer');

contextBridge.exposeInMainWorld('api', {
  filesContent: (data) => ipcRenderer.invoke('files-content', data),
  filesSave: (data) => ipcRenderer.invoke('files-save', data),
});
